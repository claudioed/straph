package repository

import (
	"github.com/straph/pkg/domain/user"
)

type Repository interface {
	Create(user *user.User) (error error)
	Delete(id string) (error error)
	ById(id string) (user *user.User, err error)
}
