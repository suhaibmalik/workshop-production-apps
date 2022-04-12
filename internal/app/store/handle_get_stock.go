package store

import (
	"net/http"
	"workshop-production-apps/pkg/jsend"
)

func HandleGetStock(storage *StoreMemoryStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		storage.lock.Lock()
		jsend.Ok(storage).Respond(w)
		storage.lock.Unlock()
	}
}
