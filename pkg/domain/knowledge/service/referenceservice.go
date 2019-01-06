package service

import (
	"github.com/straph/pkg/domain/knowledge/data"
	"github.com/straph/pkg/domain/user/repository"
)

type Service struct {
	userRepository repository.Repository
}

func new(repository repository.Repository) *Service {
	return &Service{userRepository: repository}
}

func (s *Service) AddReference(refRequest *data.ReferenceRequest, userReqReference string) error {
	s.userRepository.ById(userReqReference)
	return nil
}
