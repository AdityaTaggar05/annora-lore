package https

import (
	"github.com/AdityaTaggar05/annora-lore/internal/services"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	Service *services.LoreService
	Validator *validator.Validate
}

func NewHandler(service *services.LoreService) *Handler {
	return &Handler{
		Service: service,
		Validator: validator.New(),
	}
}
