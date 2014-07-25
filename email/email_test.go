package email

import (
	"io"
	"strings"
	"testing"

	"gopkg.in/check.v1"
)

func TestEmail(t *testing.T) {
	check.Suite(&EmailSuite{
		Sender:    "info@satisfeet.me",
		Recipient: "i@bodokaiser.io",
		Subject:   "Order Received!",
		Reader:    strings.NewReader("Hello"),
	})
	check.TestingT(t)
}

type EmailSuite struct {
	Sender    string
	Recipient string
	Subject   string
	Reader    io.Reader
}

func (s *EmailSuite) TestValid(c *check.C) {
	c.Check((&Email{
		Sender:    s.Sender,
		Recipient: s.Recipient,
		Subject:   s.Subject,
		Reader:    s.Reader,
	}).Valid(), check.IsNil)

	c.Check((&Email{
		Recipient: s.Recipient,
		Subject:   s.Subject,
		Reader:    s.Reader,
	}).Valid(), check.Equals, ErrInvalidSender)
	c.Check((&Email{
		Sender:  s.Sender,
		Subject: s.Subject,
		Reader:  s.Reader,
	}).Valid(), check.Equals, ErrInvalidRecipient)
	c.Check((&Email{
		Sender:    s.Sender,
		Recipient: s.Recipient,
		Reader:    s.Reader,
	}).Valid(), check.Equals, ErrInvalidSubject)
	c.Check((&Email{
		Sender:    s.Sender,
		Recipient: s.Recipient,
		Subject:   s.Subject,
	}).Valid(), check.Equals, ErrInvalidReader)
}
