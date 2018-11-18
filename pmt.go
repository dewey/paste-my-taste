package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/cors"
	"github.com/go-chi/render"

	"github.com/caarlos0/env"
	"github.com/dewey/paste-my-taste/client/lastfm"
	"github.com/dewey/paste-my-taste/config"
	"github.com/go-chi/chi"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

func main() {
	var cfg config.Config
	err := env.Parse(&cfg)
	if err != nil {
		panic(err)
	}

	l := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	switch strings.ToLower(cfg.Environment) {
	case "develop":
		l = level.NewFilter(l, level.AllowInfo())
	case "prod":
		l = level.NewFilter(l, level.AllowError())
	}
	l = log.With(l, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	t := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 20 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
	}
	c := &http.Client{
		Timeout:   time.Second * 30,
		Transport: t,
	}

	lfm := lastfm.New(c, cfg.APIKey)

	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))
	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		// AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		// ExposedHeaders:   []string{"Link"},
		// AllowCredentials: true,
		// MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	})

	r.Get("/api/lastfm/{username}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("api hit")
		var username, period string
		var limit int
		username = chi.URLParam(r, "username")
		if username == "" {
			http.Error(w, "username not allowed to be empty", http.StatusBadRequest)
			return
		}
		b, err := httputil.DumpRequest(r, true)
		if err == nil {
			fmt.Println(string(b))
		}
		q := r.URL.Query()
		if q.Get("period") != "" {
			period = q.Get("period")
		} else {
			period = "1month"
		}

		if q.Get("limit") != "" {
			l, err := strconv.Atoi(q.Get("limit"))
			if err != nil {
				http.Error(w, "limit has to be an integer", http.StatusInternalServerError)
				return
			}
			limit = l
		} else {
			limit = 15
		}

		var al []lastfm.TopArtist
		tal, err := lfm.GetTopArtists(username, period, limit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if len(tal) <= 3 {
			http.Error(w, "not enough listening data available", http.StatusBadRequest)
			return
		}
		for i, ta := range tal {
			if i < len(tal)/3 {
				tt, err := lfm.GetTopTags(ta.Mbid)
				if err != nil {
					continue
				}
				if len(tt) > 1 {
					ta := ta
					ta.Genre = tt[0]
					al = append(al, ta)
				}
			}
		}
		render.JSON(w, r, al)
		return
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("nothing to see here"))
	})

	l.Log("msg", fmt.Sprintf("paste-my-taste listening on http://localhost:%d", cfg.Port))
	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), r)
	if err != nil {
		panic(err)
	}
}
