package user

import (
	"github.com/straph/pkg/domain/knowledge"
	"time"
)

type User struct {
	Name string
	Email string
	Knowledge []knowledge.Knowledge
	RegisteredAt time.Time
}