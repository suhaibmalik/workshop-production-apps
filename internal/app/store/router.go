package store

import (
	"net/http"
	"sync"
	"workshop-production-apps/pkg/config"

	"github.com/go-chi/chi"
)

type StoreMemoryStorage struct {
	lock   sync.Mutex
	Apples uint `json:"apples"`
	Donuts uint `json:"donuts"`
	Pears  uint `json:"pears"`
}

func StoreRouter(cfg config.Config) http.Handler {
	storage := StoreMemoryStorage{}

	r := chi.NewRouter()

	r.Get("/stock", HandleGetStock(&storage))
	r.Post("/stock", HandleEditStock(&storage))

	return r
}
