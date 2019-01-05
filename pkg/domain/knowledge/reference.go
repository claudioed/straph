package knowledge

import (
	"github.com/straph/pkg/domain/user"
	"time"
)

type Reference struct {
	Level string
	Comment string
	User user.User
	CreatedAt time.Time
	Stars []Star
}