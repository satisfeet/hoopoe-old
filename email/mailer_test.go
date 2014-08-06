package email

import (
	"strings"
	"testing"

	"gopkg.in/check.v1"
)

func TestMailer(t *testing.T) {
	check.Suite(&MailerSuite{
		Email: &Email{
			Subject:   "Welcome to Satisfeet!",
			Sender:    "info@satisfeet.me",
			Recipient: "i@bodokaiser.io",
			Reader:    strings.NewReader("Welcome!"),
		},
	})
	check.TestingT(t)
}

type MailerSuite struct {
	Email *Email
}

func (s *MailerSuite) SetUpSuite(c *check.C) {
	c.Skip("too much traffic")
}

func (s *MailerSuite) TestOpen(c *check.C) {
	m := &Mailer{
		Service:  GMAIL,
		Username: "kyogron@googlemail.com",
		Password: "nambu007",
	}
	c.Check(m.Open(), check.IsNil)
}

func (s *MailerSuite) TestClose(c *check.C) {
	m := &Mailer{
		Service:  GMAIL,
		Username: "kyogron@googlemail.com",
		Password: "nambu007",
	}
	c.Check(m.Open(), check.IsNil)
	c.Check(m.Close(), check.IsNil)
}

func (s *MailerSuite) TestEmail(c *check.C) {
	m := &Mailer{
		Service:  GMAIL,
		Username: "kyogron@googlemail.com",
		Password: "nambu007",
	}
	c.Check(m.Open(), check.IsNil)
	c.Check(m.Email(s.Email), check.IsNil)
	c.Check(m.Close(), check.IsNil)
}
