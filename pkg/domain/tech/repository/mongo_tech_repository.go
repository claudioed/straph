package repository

import (
	"github.com/straph/pkg/domain/tech"
	"github.com/straph/pkg/infra/database/session"
	"gopkg.in/mgo.v2"
	"os"
)

const CollectionName = "technology"

type MongoTechRepository struct {
	collection *mgo.Collection
}

func newMongoTechRepository(session *session.Session, dbName string) *MongoTechRepository {
	database := os.Getenv("DATABASE")
	collection := session.GetCollection(database, CollectionName)
	return &MongoTechRepository{collection}
}

func (m *MongoTechRepository) Create(tech *tech.Tech) (error error) {
	return m.collection.Insert(tech)
}

func (m *MongoTechRepository) ById(id string) (tech *tech.Tech, err error) {
	return nil, nil
}
