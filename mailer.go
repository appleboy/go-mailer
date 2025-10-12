package mailer

import (
	"errors"
	"github.com/rs/zerolog/log"
)

// Mail defines the interface for email sending operations
type Mail interface {
	From(string, string) Mail
	To(...string) Mail
	Cc(...string) Mail
	Subject(string) Mail
	Body(string) Mail
	Send() (interface{}, error)
}

// Config holds the configuration for email drivers
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Driver   string
	Region   string
}

// Client for mail interface
var Client Mail

// NewEngine creates and returns a new Mail instance based on the provided configuration
func NewEngine(c Config) (Mail, error) {
	if c.Driver == "" {
		return nil, errors.New("driver is required")
	}

	var mail Mail
	var err error
	switch c.Driver {
	case "smtp":
		if c.Host == "" || c.Port == "" {
			return nil, errors.New("SMTP host and port are required")
		}
		mail, err = SMTPEngine(
			c.Host,
			c.Port,
			c.Username,
			c.Password,
		)
		if err != nil {
			return nil, err
		}
		Client = mail
	case "ses":
		if c.Region == "" {
			return nil, errors.New("SES region is required")
		}
		mail, err = SESEngine(c.Region)
		if err != nil {
			return nil, err
		}
		Client = mail
	default:
		log.Error().Str("driver", c.Driver).Msg("Unknown email driver")
		return nil, errors.New("unsupported email driver: " + c.Driver)
	}

	return mail, nil
}
