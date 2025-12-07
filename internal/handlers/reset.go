package handlers

import (
	"net/http"
	"sync/atomic"
)

type APIConfig struct {
	FileserverHits atomic.Int32
}

func (a *APIConfig) Reset(w http.ResponseWriter, r *http.Request) {
	a.FileserverHits.Store(0)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits reset to 0"))
}
