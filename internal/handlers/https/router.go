package https

import (
	"net/http"

	"github.com/go-chi/chi"
)

func NewRouter(loreHandler Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Service is up and running!"))
		w.WriteHeader(http.StatusOK)
	})

	r.Post("/nodes/create", loreHandler.HandleCreateNode)

	return r
}
