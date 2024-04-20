package services

import "recommendation/internal/repository"

type Service struct {
}

func NewServices(repo *repository.Repository) *Service {
	return &Service{}
}
