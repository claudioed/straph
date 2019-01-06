package user

import (
	"github.com/straph/pkg/domain/knowledge"
	"time"
)

type User struct {
	Id           string
	Name         string
	Email        string
	Knowledge    []knowledge.Knowledge
	RegisteredAt time.Time
}
