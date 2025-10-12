package mailer

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/rs/zerolog/log"
)

const (
	// CharSet The character encoding for the email.
	CharSet = "UTF-8"
)

// SES for aws ses
type SES struct {
	cfg     aws.Config
	source  *string
	to      []string
	cc      []string
	subject *string
	body    *string
}

// From for sender information
func (c SES) From(name, address string) Mail {
	c.source = aws.String(fmt.Sprintf("%s <%s>", name, address))

	return c
}

// To for mailto list
func (c SES) To(address ...string) Mail {
	for _, v := range address {
		c.to = append(c.to, v)
	}

	return c
}

// Cc for cc list
func (c SES) Cc(address ...string) Mail {
	for _, v := range address {
		c.cc = append(c.cc, v)
	}

	return c
}

// Subject for email title
func (c SES) Subject(subject string) Mail {
	c.subject = aws.String(subject)

	return c
}

// Body for email body
func (c SES) Body(body string) Mail {
	c.body = aws.String(body)

	return c
}

// Send email
func (c SES) Send() (interface{}, error) {
	// Create an SES client.
	svc := ses.NewFromConfig(c.cfg)

	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &types.Destination{
			CcAddresses: c.cc,
			ToAddresses: c.to,
		},
		Message: &types.Message{
			Body: &types.Body{
				Html: &types.Content{
					Charset: aws.String(CharSet),
					Data:    c.body,
				},
				Text: &types.Content{
					Charset: aws.String(CharSet),
					Data:    c.body,
				},
			},
			Subject: &types.Content{
				Charset: aws.String(CharSet),
				Data:    c.subject,
			},
		},
		Source: c.source,
		// Uncomment to use a configuration set
		// ConfigurationSetName: aws.String(ConfigurationSet),
	}

	// Attempt to send the email.
	resp, err := svc.SendEmail(context.TODO(), input)
	// Display error messages if they occur.
	if err != nil {
		log.Error().Err(err).Msg("AWS SES Error")
		return nil, err
	}

	return resp, nil
}

// SESEngine initial ses
func SESEngine(region string) (*SES, error) {
	// Create a new config in the specified region.
	// Replace with the AWS Region you're using for Amazon SES.
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return nil, err
	}

	return &SES{
		cfg: cfg,
	}, nil
}
