package service

import (
	"github.com/straph/pkg/domain/knowledge/data"
	"github.com/straph/pkg/domain/user/repository"
)

type KnowledgeService struct {
	userRepository repository.Repository
}

func newKnowledgeService(repository repository.Repository) *KnowledgeService {
	return &KnowledgeService{userRepository: repository}
}

func (s *KnowledgeService) AddReference(refRequest *data.ReferenceRequest, userReqReference string) error {
	s.userRepository.ById(userReqReference)
	return nil
}
