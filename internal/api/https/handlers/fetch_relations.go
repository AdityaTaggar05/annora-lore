package handlers

import (
	"log"
	"net/http"

	dtoresponse "github.com/AdityaTaggar05/annora-lore/internal/api/https/dto/response"
	"github.com/AdityaTaggar05/annora-lore/internal/domain/validation"
	"github.com/AdityaTaggar05/annora-lore/pkg/response"
	"github.com/go-chi/chi"
)

func (h *Handler) FetchRelationsOfNode(w http.ResponseWriter, r *http.Request) {
	worldID := chi.URLParam(r, "world_id")
	nodeID := chi.URLParam(r, "node_id")

	if err := validation.ValidateID(worldID); err != nil {
		log.Println("err:", err)
		response.BadRequest(w, "invalid world id", map[string]any{"world_id": worldID})
		return 
	}

	if err := validation.ValidateID(nodeID); err != nil {
		log.Println("err:", err)
		response.BadRequest(w, "invalid node id", map[string]any{"node_id": nodeID})
		return 
	}

	relations, err := h.Service.FetchAllRelationsOfNode(r.Context(), nodeID)

	if err != nil {
		log.Println("err:", err)
		response.InternalServerError(w, err.Error())
		return
	}
	
	data := make([]dtoresponse.LoreRelationResponseDTO, len(relations))
	for i, relation := range relations {
		data[i] = dtoresponse.ToLoreRelationResponse(*relation)
	}

	response.JSON(w, http.StatusOK, data)
}
