package knowledge

import (
	"github.com/straph/pkg/domain/user"
	"time"
)

type Star struct {
	Evaluation int8
	User user.User
	CreatedAt time.Time
}