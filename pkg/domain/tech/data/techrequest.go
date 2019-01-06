package data

import (
	"github.com/satori/go.uuid"
	"github.com/straph/pkg/domain/tech"
	"time"
)

type TechRequest struct {
	Name     string `json:"name"`
	Maturity string `json:"maturity"`
}

func (techRequest *TechRequest) ToDomain() *tech.Tech {

	return &tech.Tech{
		Id:        uuid.NewV4().String(),
		Name:      techRequest.Name,
		Maturity:  techRequest.Maturity,
		CreatedAt: time.Now(),
	}

}
