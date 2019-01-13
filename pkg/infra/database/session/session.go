package session

import (
	"gopkg.in/mgo.v2"
	"log"
	"os"
)

type Session struct {
	session *mgo.Session
}

func NewSession(url string) (*Session, error) {
	host := os.Getenv("MONGO_HOST")
	log.Print("Connecting mongodb at host: " + host)
	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}
	return &Session{session}, err
}

func (s *Session) Copy() *Session {
	return &Session{s.session.Copy()}
}

func (s *Session) GetCollection(db string, col string) *mgo.Collection {
	log.Print("Getting collection name: " + col)
	return s.session.DB(db).C(col)
}

func (s *Session) Close() {
	if s.session != nil {
		log.Print("Closing mongoDB connection")
		s.session.Close()
	}
}
