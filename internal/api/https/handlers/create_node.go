package handlers

import (
	"log"
	"net/http"

	"github.com/AdityaTaggar05/annora-lore/pkg/response"
)

func (h *Handler) HandleCreateNode(w http.ResponseWriter, r *http.Request) {
	err := h.Service.CreateNode(r.Context())

	if err != nil {
		log.Println("err: ", err)
		response.InternalServerError(w, err.Error())
		return
	}

	response.Created(w, nil, "Node created successfully")
}
