package https

import (
	"net/http"

	"github.com/AdityaTaggar05/annora-lore/internal/api/https/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter(loreHandler handlers.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Service is up and running!"))
		w.WriteHeader(http.StatusOK)
	})

	r.Post("/nodes/create", loreHandler.HandleCreateNode)
	r.Post("/nodes/relate", loreHandler.HandleCreateRelation)

	r.Get("/world/{world_id}/nodes/{node_id}/relations", loreHandler.FetchRelationsOfNode)
	r.Get("/world/{world_id}/nodes/{node_id}", loreHandler.FetchNodeByID)
	r.Get("/world/{world_id}/nodes", loreHandler.FetchNodesByType)

	return r
}
