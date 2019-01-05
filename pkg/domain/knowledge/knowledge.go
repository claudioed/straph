package knowledge

import (
	"github.com/straph/pkg/domain/tech"
	"github.com/straph/pkg/domain/user"
	"time"
)

type Knowledge struct {
	Id string
	Tech tech.Tech
	User user.User
	References []Reference
	RegisteredAt time.Time
}