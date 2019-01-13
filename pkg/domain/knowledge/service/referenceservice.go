package service

import (
	"github.com/straph/pkg/domain/knowledge/data"
	"github.com/straph/pkg/domain/user/repository"
)

type ReferenceService struct {
	userRepository repository.Repository
}

func newReferenceService(repository repository.Repository) *ReferenceService {
	return &ReferenceService{userRepository: repository}
}

func (s *ReferenceService) AddReference(refRequest *data.ReferenceRequest, userReqReference string) error {
	s.userRepository.ById(userReqReference)
	return nil
}
