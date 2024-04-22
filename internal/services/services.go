package services

import "github.com/Mirzoev-Parviz/movie-recommendation/internal/repository"

type Service struct {
	CSV_Parser
	RCM
}

func NewServices(repo *repository.Repository) *Service {
	return &Service{
		NewCSV_Parser(),
		NewRCM_Service(),
	}
}
