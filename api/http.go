package api

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// NewHandler initializes a new router
func NewHandler(s service) *chi.Mux {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	// Public routes
	r.Group(func(r chi.Router) {
		r.Get("/lastfm/{username}", promhttp.InstrumentHandlerCounter(
			promauto.NewCounterVec(
				prometheus.CounterOpts{
					Namespace: "pmt",
					Subsystem: "api_lastfm",
					Name:      "username_reqs_total",
					Help:      "Total number of requests by HTTP code",
				},
				[]string{"code", "method"},
			),
			getLFMTaste(s),
		))
	})

	return r
}

// getLFMTaste returns the generated taste from Last.FM for a user
func getLFMTaste(s service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var username, period string
		var limit int
		username = chi.URLParam(r, "username")
		if username == "" {
			http.Error(w, "username not allowed to be empty", http.StatusBadRequest)
			return
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

		al, err := s.GenerateTaste(username, period, limit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, al)
		return
	}
}
