package storage

import (
	"log"
	"time"

	r "github.com/dancannon/gorethink"
)

type Session *r.Session

type Storage struct {
	Session     Session
	ConnectOpts r.ConnectOpts
}

type StorageConfig interface {
	Configure(*Storage)
}

func (s *Storage) GetSession() Session {
	return s.Session
}

func NewStorage(settings ...StorageConfig) *Storage {
	var s Storage

	s.ConnectOpts = r.ConnectOpts{
		Address:     "localhost:28015",
		Database:    "foos",
		MaxIdle:     10,
		IdleTimeout: time.Second * 10,
	}

	for _, setting := range settings {
		setting.Configure(&s)
	}

	session, err := r.Connect(s.ConnectOpts)

	if err != nil {
		log.Fatalln(err.Error())
	}

	s.Session = session

	return &s
}
