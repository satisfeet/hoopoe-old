package email

import (
	"crypto/tls"
	"errors"
	"io"
	"net/smtp"
)

const (
	// Available email services to use.
	GMAIL = "smtp.gmail.com:465"
)

// Mailer handles an smtp connection to a defined service.
type Mailer struct {
	conn *smtp.Client

	Service  string
	Username string
	Password string
}

var ErrNotConnected = errors.New("not connected")

// Returns credentials formatted as smtp plain auth struct.
func (m *Mailer) auth() smtp.Auth {
	return smtp.PlainAuth("", m.Username, m.Password, m.Service)
}

// Opens encrypted connection to email server.
func (m *Mailer) Open() error {
	t, err := tls.Dial("tcp", m.Service, &tls.Config{
		ServerName:         m.Service,
		InsecureSkipVerify: true,
	})
	if err != nil {
		return err
	}

	c, err := smtp.NewClient(t, m.Service)
	if err != nil {
		return err
	}

	if err := c.Auth(m.auth()); err != nil {
		return err
	}
	m.conn = c

	return nil
}

// Closes connection to email server.
func (m *Mailer) Close() error {
	return m.conn.Close()
}

// Sends an initialized email through connection.
func (m *Mailer) Email(e *Email) error {
	if m.conn == nil {
		return ErrNotConnected
	}

	if err := e.Valid(); err != nil {
		return err
	}

	m.conn.Mail(e.Sender)
	m.conn.Rcpt(e.Recipient)

	w, err := m.conn.Data()
	if err != nil {
		return err
	}
	defer w.Close()

	_, err = io.Copy(w, e.Reader)
	if err != nil {
		return err
	}
	return nil
}
