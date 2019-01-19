package server

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/straph/pkg/domain/tech/service"
	"log"
	"net/http"
	"os"
)

const DefaultServerPort string = "8080"

type Server struct {
	router *mux.Route
}

func NewServer(service service.TechService) *Server {

	return nil
}

func (s *Server) Start() {
	var serverPort = os.Getenv("SERVER_PORT")
	if len(serverPort) > 0 {
		log.Println("Server Port is " + serverPort)
	} else {
		log.Println("Missing SERVER_PORT environment variable. The default port is" + DefaultServerPort)
		serverPort = DefaultServerPort
	}

	log.Println("Starting server ...")
	log.Println("Listening on port " + serverPort)

	if err := http.ListenAndServe(serverPort, handlers.LoggingHandler(os.Stdout, s.router.GetHandler())); err != nil {
		log.Fatal("error on start http server: ", err)
	}

}

func (s *Server) newSubRouter(path string) *mux.Router {
	return s.router.PathPrefix(path).Subrouter()
}
