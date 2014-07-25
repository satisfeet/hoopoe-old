package email

import (
	"errors"
	"io"
)

type Email struct {
	Subject   string
	Sender    string
	Recipient string
	Reader    io.Reader
}

var (
	ErrInvalidSender    = errors.New("invalid name")
	ErrInvalidRecipient = errors.New("invalid email")
	ErrInvalidReader    = errors.New("missing reader")
	ErrInvalidSubject   = errors.New("missing subject")
)

// Returns error if email field is missing.
func (e *Email) Valid() error {
	if len(e.Subject) == 0 {
		return ErrInvalidSubject
	}
	if len(e.Sender) == 0 {
		return ErrInvalidSender
	}
	if len(e.Recipient) == 0 {
		return ErrInvalidRecipient
	}
	if e.Reader == nil {
		return ErrInvalidReader
	}
	return nil
}
