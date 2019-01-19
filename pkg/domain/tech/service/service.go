package service

import (
	"github.com/straph/pkg/domain/tech/data"
	"github.com/straph/pkg/domain/tech/repository"
)

type TechService struct {
	repository repository.Repository
}

func New(repository repository.Repository) *TechService {
	return &TechService{repository: repository}
}

func (s *TechService) RegisterTech(tech *data.TechRequest) error {
	return s.repository.Create(tech.ToDomain())
}
