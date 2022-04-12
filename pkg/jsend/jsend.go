package jsend

import (
	"encoding/json"
	"net/http"
)

// Ref: https://github.com/omniti-labs/jsend

type JSend struct {
	StatusCode int         `json:"-"`
	Status     string      `json:"status"`
	Data       interface{} `json:"data,omitempty"`
	Message    interface{} `json:"message,omitempty"`
}

func (j JSend) Respond(w http.ResponseWriter) {
	js, err := json.Marshal(j)
	if err != nil {
		// todo: panic instead?
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(j.StatusCode)
	w.Write(js)
}
