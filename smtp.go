package mailer

import (
	"time"

	"github.com/xhit/go-simple-mail/v2"
)

type from struct {
	Name    string
	Address string
}

// SMTP setting
type SMTP struct {
	host     string
	port     string
	username string
	password string
	from     from
	to       []string
	cc       []string
	subject  string
	body     string
}

// From for sender information
func (c SMTP) From(name, address string) Mail {
	c.from = from{
		Name:    name,
		Address: address,
	}

	return c
}

// To for mailto list
func (c SMTP) To(address ...string) Mail {
	c.to = address

	return c
}

// Cc for cc list
func (c SMTP) Cc(address ...string) Mail {
	c.cc = address

	return c
}

// Subject for email title
func (c SMTP) Subject(subject string) Mail {
	c.subject = subject

	return c
}

// Body for email body
func (c SMTP) Body(body string) Mail {
	c.body = body

	return c
}

// Send email
func (c SMTP) Send() (interface{}, error) {
	server := mail.NewSMTPClient()
	server.Host = c.host
	server.Port = 587 // default port
	switch c.port {
	case "25":
		server.Port = 25
	case "465":
		server.Port = 465
		server.Encryption = mail.EncryptionSSL
	case "587":
		server.Port = 587
		server.Encryption = mail.EncryptionTLS
	}
	server.Username = c.username
	server.Password = c.password
	server.Authentication = mail.AuthPlain
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	smtpClient, err := server.Connect()
	if err != nil {
		return nil, err
	}
	defer smtpClient.Close()

	email := mail.NewMSG()
	if c.from.Name != "" {
		email.SetFrom(c.from.Name + " <" + c.from.Address + ">")
	} else {
		email.SetFrom(c.from.Address)
	}
	email.SetSubject(c.subject).
		SetBody(mail.TextHTML, c.body)

	for _, to := range c.to {
		email.AddTo(to)
	}

	for _, cc := range c.cc {
		email.AddCc(cc)
	}

	err = email.Send(smtpClient)
	return nil, err
}

// SMTPEngine initial smtp object
func SMTPEngine(host, port, username, password string) (*SMTP, error) {
	return &SMTP{
		host:     host,
		port:     port,
		username: username,
		password: password,
	}, nil
}
