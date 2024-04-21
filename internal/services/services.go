package services

import "github.com/Mirzoev-Parviz/movie-recommendation/internal/repository"

type Service struct {
}

func NewServices(repo *repository.Repository) *Service {
	return &Service{}
}
