package api

import (
	"errors"
	"time"

	"github.com/go-kit/kit/log"

	"github.com/dewey/paste-my-taste/client/lastfm"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/dewey/paste-my-taste/config"
	"github.com/dewey/paste-my-taste/store"
)

// Service provides access to the functions by the public API
type Service interface {
	// ServeFeed serves a feed based on the plugin and format
	GenerateTaste(username string, period string, limit int) (string, error)
}

type service struct {
	l                 log.Logger
	cfg               config.Config
	storageRepository store.StorageRepository
	client            *lastfm.Client
}

// NewService initializes a new generation service
func NewService(l log.Logger, cfg config.Config, c *lastfm.Client, sr store.StorageRepository) *service {
	return &service{
		l:                 l,
		cfg:               cfg,
		client:            c,
		storageRepository: sr,
	}
}

func init() {
	prometheus.MustRegister(fetchDurationHistogram, cacheHitTotalCounter, cacheRequestTotalCounter, cacheMissTotalCounter)
}

var fetchDurationHistogram = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Namespace: "pmt",
		Subsystem: "fetch",
		Name:      "fetch_duration_seconds",
		Help:      "Time to fetch data from Last.FM API.",
		Buckets:   []float64{5, 10, 20, 60, 120},
	},
	[]string{"period", "task"},
)

var cacheHitTotalCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: "pmt",
		Subsystem: "cache",
		Name:      "hit_total",
		Help:      "Cache hits on the k/v store",
	},
)

var cacheMissTotalCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: "pmt",
		Subsystem: "cache",
		Name:      "miss_total",
		Help:      "Cache misses on the k/v store",
	},
)

var cacheRequestTotalCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: "pmt",
		Subsystem: "cache",
		Name:      "reqs_total",
		Help:      "Number of requests on the k/v store",
	},
)

var tasteGenerateOps = promauto.NewCounter(prometheus.CounterOpts{
	Namespace: "pmt",
	Subsystem: "generate",
	Name:      "taste_generate_ops_total",
	Help:      "The total number of generated tastes",
})

var tasteGenerateDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Namespace: "pmt",
		Subsystem: "generate",
		Name:      "taste_generate_duration_seconds",
		Help:      "Time to generate the full taste.",
		Buckets:   []float64{5, 10, 20, 60, 120},
	},
	[]string{"period"},
)

// ServeFeed returns the generated feed from the store backend
func (s *service) GenerateTaste(username string, period string, limit int) ([]lastfm.TopArtist, error) {
	start := time.Now()
	var al []lastfm.TopArtist
	var tal []lastfm.TopArtist
	var err error

	// Last.FM doesn't have a two week parameter so we have to use another route that supports ranges
	switch period {
	case "2week":
		tal, err = s.client.GetWeeklyArtistChart(username, start.AddDate(0, 0, -14).Unix(), start.Unix())
		if err != nil {
			s.l.Log("err", err)
			return nil, errors.New("error getting weekly top artists from Last.FM")
		}
	default:
		tal, err = s.client.GetTopArtists(username, period, limit)
		if err != nil {
			s.l.Log("err", err)
			return nil, errors.New("error getting top artists from Last.FM")
		}
	}

	fetchDurationHistogram.WithLabelValues(period, "get_top_artists").Observe(time.Since(start).Seconds())
	if len(tal) < 3 {
		return nil, errors.New("not enough listening data available")
	}
	// TODO(dewey): This is a mess, needs to be cleaned up and split
	m := make(map[string]struct{})
	for i, ta := range tal {
		ta := ta

		// We only check for the top 4 artists with a MB ID
		var tt string
		if i < 5 && ta.Mbid != "" {
			// We check if we already have the artist in our cache
			tt, err = s.storageRepository.Get(ta.Mbid)
			cacheRequestTotalCounter.Add(1)
			if err != nil {
				// If we don't have it in our store we get the unique top tag for the first 4 artists
				start = time.Now()
				tt, err := s.client.GetTopTags(ta.Mbid)
				fetchDurationHistogram.WithLabelValues(period, "get_top_tags").Observe(time.Since(start).Seconds())
				cacheMissTotalCounter.Add(1)
				if err != nil {
					s.l.Log("err", err)
					continue
				}
				if len(tt) > 1 {
					ta.Genre = tt[0]
					if _, ok := m[ta.Genre]; !ok {
						al = append(al, ta)
						m[ta.Genre] = struct{}{}
					}
					// We store the newly fetched information in our storage backend so we don't have to fetch it again for the next user
					s.storageRepository.Save(ta.Mbid, tt[0])
				}
				continue
			}
			cacheHitTotalCounter.Add(1)
		} else {
			al = append(al, ta)
		}

		// If we have top tags in our storage already we add them to the artist, or else append the artist without tags
		if tt != "" {
			ta.Genre = tt
			if _, ok := m[ta.Genre]; !ok {
				al = append(al, ta)
				m[ta.Genre] = struct{}{}
			}
		}
	}
	tasteGenerateDuration.WithLabelValues(period).Observe(time.Since(start).Seconds())
	tasteGenerateOps.Inc()
	return al, nil
}
