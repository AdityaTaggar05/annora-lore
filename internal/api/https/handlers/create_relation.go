package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AdityaTaggar05/annora-lore/internal/api/https/dto/request"
	dtoresponse "github.com/AdityaTaggar05/annora-lore/internal/api/https/dto/response"
	"github.com/AdityaTaggar05/annora-lore/pkg/response"
)

func (h *Handler) HandleCreateRelation(w http.ResponseWriter, r *http.Request) {
	var dto request.CreateRelationRequestDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		log.Println("err: ", err)
		response.BadRequest(w, "invalid request JSON", nil)
		return
	}

	if err := h.Validator.Struct(dto); err != nil {
		log.Println("err: ", err)
		response.Error(w, http.StatusUnprocessableEntity, "Unprocessable Entity", err.Error(), nil)
		return
	}

	from, rel, to, err := h.Service.CreateRelation(r.Context(), dto)
	if err != nil {
		log.Println("err: ", err)
		response.InternalServerError(w, err.Error())
		return
	}

	response.Created(w, dtoresponse.ToLoreRelationResponse(*from, *rel, *to), "Relation created successfully")
}
