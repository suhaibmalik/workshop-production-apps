package main

import (
	"net/http"
	"workshop-production-apps/internal/app/store"
	"workshop-production-apps/pkg/config"
	"workshop-production-apps/pkg/logs"
	"workshop-production-apps/pkg/middlewares"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	cfg := config.Config{}

	r.Mount("/store", store.StoreRouter(cfg))

	r.With(middlewares.DumpHeaders).Get("/dump_headers", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("check the logs"))
	})

	logs.Info(nil, "server is about to be up")

	http.ListenAndServe(":3000", r)
}
