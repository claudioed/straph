package data

import (
	"github.com/straph/pkg/domain/tech"
	"time"
)

type TechRequest struct {
	Name     string
	Maturity string
}

func (techRequest *TechRequest) ToDomain() *tech.Tech {

	return &tech.Tech{
		Id:        "",
		Name:      techRequest.Name,
		Maturity:  techRequest.Maturity,
		CreatedAt: time.Now(),
	}

}
