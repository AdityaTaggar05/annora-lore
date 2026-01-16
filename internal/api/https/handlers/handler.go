package handlers

import (
	"github.com/AdityaTaggar05/annora-lore/internal/domain/service"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	Service   *service.LoreService
	Validator *validator.Validate
}

func NewHandler(service *service.LoreService) *Handler {
	return &Handler{
		Service:   service,
		Validator: validator.New(),
	}
}
