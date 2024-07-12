package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/caarlos0/env"
	"github.com/dewey/paste-my-taste/api"
	"github.com/dewey/paste-my-taste/client/lastfm"
	"github.com/dewey/paste-my-taste/config"
	"github.com/dewey/paste-my-taste/store"
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

	storageRepo, err := store.NewStoreBackend(cfg)
	if err != nil {
		return
	}
	if err := storageRepo.Save("test", "test"); err != nil {
		l.Log("msg", fmt.Sprintf("failed to save test key: %s,", err), "backend", cfg.StorageBackend)
	}

	as := api.NewService(l, cfg, lfm, storageRepo)

	// TODO(dewey): Figure out if there's a better way
	var assetPath string
	switch cfg.Environment {
	case "prod":
		assetPath = "/web/dist"
	default:
		assetPath = "./web/dist"
	}

	r := chi.NewRouter()
	r.Handle("/*", http.FileServer(http.Dir(assetPath)))
	r.Handle("/metrics", promhttp.Handler())
	r.Mount("/api", api.NewHandler(*as))

	l.Log("msg", fmt.Sprintf("paste-my-taste listening on http://localhost:%d", cfg.Port))
	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), r)
	if err != nil {
		panic(err)
	}
}
