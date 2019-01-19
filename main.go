package straph

import (
	"github.com/straph/pkg/infra/database/session"
	"github.com/straph/pkg/infra/server"
	"log"
)

func main() {

	databaseSession, err := session.NewSession()
	if err != nil {
		log.Fatalln("Unable to connect database")
	}

	defer databaseSession.Close()

	techRepository := newMongoTechRepository(databaseSession.Copy())

	server.new

}
