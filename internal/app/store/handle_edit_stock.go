package store

import (
	"encoding/json"
	"net/http"
	"workshop-production-apps/pkg/jsend"
)

type EditStockRequest struct {
	Apples int `json:"apples"`
	Donuts int `json:"donuts"`
	Pears  int `json:"pears"`
}

func HandleEditStock(storage *StoreMemoryStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := EditStockRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			jsend.Malformed().Respond(w)
			return
		}

		storage.lock.Lock()
		defer storage.lock.Unlock()

		apples := req.Apples + int(storage.Apples)
		donuts := req.Donuts + int(storage.Donuts)
		pears := req.Pears + int(storage.Pears)

		if apples < 0 {
			apples = 0
		}

		if donuts < 0 {
			donuts = 0
		}

		if pears < 0 {
			pears = 0
		}

		storage.Apples = uint(apples)
		storage.Donuts = uint(donuts)
		storage.Pears = uint(pears)

		jsend.Ok(storage).Respond(w)
	}
}
