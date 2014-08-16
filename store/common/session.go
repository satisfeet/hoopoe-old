package common

import "github.com/satisfeet/hoopoe/store/mongo"

// The Session type manages multiple database sessions which can be controlled
// through an unified interface.
type Session struct {
	mongo *mongo.Session
}

// Returns an initialized Session.
func NewSession() *Session {
	return &Session{
		mongo: &mongo.Session{},
	}
}

// Connects to all defined databases.
func (s *Session) Dial(url string) error {
	return s.mongo.Dial(url)
}

// Disconnects from all defined databases.
func (s *Session) Close() error {
	return s.mongo.Close()
}
