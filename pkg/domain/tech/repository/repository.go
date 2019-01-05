package repository

import "github.com/straph/pkg/domain/tech"

type Repository interface {
	Create(tech *tech.Tech) (error error)
	Delete(id string) (error error)
	ById(id string) (tech *tech.Tech, err error)
	All() (techs []*tech.Tech, err error)
}
