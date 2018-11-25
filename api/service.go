package api

import (
	"errors"
	"time"

	"github.com/go-kit/kit/log"

	"github.com/dewey/paste-my-taste/client/lastfm"
	"github.com/prometheus/client_golang/prometheus"

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
	[]string{"period"},
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

// ServeFeed returns the generated feed from the store backend
func (s *service) GenerateTaste(username string, period string, limit int) ([]lastfm.TopArtist, error) {
	start := time.Now()
	var al []lastfm.TopArtist
	tal, err := s.client.GetTopArtists(username, period, limit)
	if err != nil {
		return nil, err
	}
	if len(tal) < 3 {
		return nil, errors.New("not enough listening data available")
	}
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
				tt, err := s.client.GetTopTags(ta.Mbid)
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
		if len(tt) > 1 {
			ta.Genre = tt
			if _, ok := m[ta.Genre]; !ok {
				al = append(al, ta)
				m[ta.Genre] = struct{}{}
			}
		} else {
			al = append(al, ta)
		}
	}
	duration := time.Since(start)
	fetchDurationHistogram.WithLabelValues(period).Observe(duration.Seconds())
	return tal, nil
}
