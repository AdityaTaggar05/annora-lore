package https

import (
	"net/http"

	"github.com/go-chi/chi"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Service is up and running!"))
		w.WriteHeader(http.StatusOK)
	})

	return r
}
