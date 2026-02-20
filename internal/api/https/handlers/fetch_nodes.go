package handlers

import (
	"fmt"
	"log"
	"net/http"

	dtoresponse "github.com/AdityaTaggar05/annora-lore/internal/api/https/dto/response"
	"github.com/AdityaTaggar05/annora-lore/internal/domain/validation"
	"github.com/AdityaTaggar05/annora-lore/pkg/response"
	"github.com/go-chi/chi"
)

func (h *Handler) FetchNodeByID(w http.ResponseWriter, r *http.Request) {
	worldID := chi.URLParam(r, "world_id")
	nodeID := chi.URLParam(r, "node_id")

	if err := validation.ValidateID(worldID); err != nil {
		log.Println("err:", err)
		response.BadRequest(w, "invalid world id", map[string]any{ "world_id": worldID })
		return
	}

	if err := validation.ValidateID(nodeID); err != nil {
		log.Println("err:", err)
		response.BadRequest(w, "invalid node id", map[string]any{ "node_id": nodeID })
		return
	}

	node, err := h.Service.FetchNodeByID(r.Context(), worldID, nodeID)

	if err != nil {
		fmt.Println("err:", err)
		response.InternalServerError(w, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, dtoresponse.ToLoreNodeResponse(*node))
}

func (h *Handler) FetchNodesByType(w http.ResponseWriter, r *http.Request) {
	worldID := chi.URLParam(r, "world_id")
	nodeType := r.URL.Query()["type"]

	if err := validation.ValidateID(worldID); err != nil {
		log.Println("err:", err)
		response.BadRequest(w, "invalid world id", map[string]any{ "world_id": worldID })
		return
	}

	nodes, err := h.Service.FetchNodesByType(r.Context(), worldID, nodeType)

	if err != nil {
		fmt.Println("err:", err)
		response.InternalServerError(w, err.Error())
		return
	}

	resNodes := make([]dtoresponse.LoreNodeResponseDTO, len(nodes))

	for i, node := range nodes {
		resNodes[i] = dtoresponse.ToLoreNodeResponse(*node)
	}

	response.JSON(w, http.StatusOK, resNodes)
}
