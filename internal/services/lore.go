package services

import "github.com/AdityaTaggar05/annora-lore/internal/repository"

type LoreService struct {
	Repo *repository.LoreRepository
}

func NewLoreService(repo *repository.LoreRepository) *LoreService {
	return &LoreService{
		Repo: repo,
	}
}
