package service

import (
	"github.com/straph/pkg/domain/tech"
	"github.com/straph/pkg/domain/tech/repository"
)

type Service struct {
	repository repository.Repository
}

func new(repository repository.Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) RegisterTech(tech *tech.Tech) error {
	return s.repository.Create(tech)
}
