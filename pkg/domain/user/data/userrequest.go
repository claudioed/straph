package data

import (
	"github.com/satori/go.uuid"
	"github.com/straph/pkg/domain/user"
	"time"
)

type UserRequest struct {
	Name  string
	Email string
}

func (userRequest *UserRequest) ToDomain() *user.User {

	return &user.User{
		Id:           uuid.NewV4().String(),
		Name:         userRequest.Name,
		Email:        userRequest.Email,
		RegisteredAt: time.Now(),
	}

}
