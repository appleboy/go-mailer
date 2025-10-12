package mailer

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
)

func TestSESEngine(t *testing.T) {
	// Note: This test requires AWS credentials to be configured properly
	// We'll test the basic structure without actually connecting to AWS
	ses, err := SESEngine("us-east-1")
	if err != nil {
		// Skip this test if AWS configuration is not available
		t.Skipf("SESEngine() error (expected in CI/test environment): %v", err)
	}

	if ses != nil {
		if ses.cfg.Region != "us-east-1" {
			t.Errorf("SESEngine() region = %v, want %v", ses.cfg.Region, "us-east-1")
		}
	}
}

func TestSESFrom(t *testing.T) {
	ses := &SES{}

	result := ses.From("John Doe", "john@example.com")

	sesResult, ok := result.(SES)
	if !ok {
		t.Errorf("From() did not return SES type")
	}

	expectedSource := "John Doe <john@example.com>"
	if sesResult.source == nil || *sesResult.source != expectedSource {
		t.Errorf("From() source = %v, want %v", aws.ToString(sesResult.source), expectedSource)
	}
}

func TestSESTo(t *testing.T) {
	ses := &SES{}

	result := ses.To("recipient1@example.com", "recipient2@example.com")

	sesResult, ok := result.(SES)
	if !ok {
		t.Errorf("To() did not return SES type")
	}

	expectedTo := []string{"recipient1@example.com", "recipient2@example.com"}
	if len(sesResult.to) != len(expectedTo) {
		t.Errorf("To() length = %v, want %v", len(sesResult.to), len(expectedTo))
	}

	for i, addr := range expectedTo {
		if sesResult.to[i] != addr {
			t.Errorf("To() address[%d] = %v, want %v", i, sesResult.to[i], addr)
		}
	}
}

func TestSESCc(t *testing.T) {
	ses := &SES{}

	result := ses.Cc("cc1@example.com", "cc2@example.com")

	sesResult, ok := result.(SES)
	if !ok {
		t.Errorf("Cc() did not return SES type")
	}

	expectedCc := []string{"cc1@example.com", "cc2@example.com"}
	if len(sesResult.cc) != len(expectedCc) {
		t.Errorf("Cc() length = %v, want %v", len(sesResult.cc), len(expectedCc))
	}

	for i, addr := range expectedCc {
		if sesResult.cc[i] != addr {
			t.Errorf("Cc() address[%d] = %v, want %v", i, sesResult.cc[i], addr)
		}
	}
}

func TestSESSubject(t *testing.T) {
	ses := &SES{}

	result := ses.Subject("Test Subject")

	sesResult, ok := result.(SES)
	if !ok {
		t.Errorf("Subject() did not return SES type")
	}

	if sesResult.subject == nil || *sesResult.subject != "Test Subject" {
		t.Errorf("Subject() = %v, want %v", aws.ToString(sesResult.subject), "Test Subject")
	}
}

func TestSESBody(t *testing.T) {
	ses := &SES{}

	result := ses.Body("<h1>Test Body</h1>")

	sesResult, ok := result.(SES)
	if !ok {
		t.Errorf("Body() did not return SES type")
	}

	if sesResult.body == nil || *sesResult.body != "<h1>Test Body</h1>" {
		t.Errorf("Body() = %v, want %v", aws.ToString(sesResult.body), "<h1>Test Body</h1>")
	}
}

func TestSESChaining(t *testing.T) {
	ses := &SES{}

	result := ses.
		From("John Doe", "john@example.com").
		To("recipient@example.com").
		Cc("cc@example.com").
		Subject("Test Subject").
		Body("Test Body")

	sesResult, ok := result.(SES)
	if !ok {
		t.Errorf("Method chaining did not return SES type")
	}

	// Verify all fields are set correctly
	expectedSource := "John Doe <john@example.com>"
	if sesResult.source == nil || *sesResult.source != expectedSource {
		t.Errorf("Chained from source = %v, want %v", aws.ToString(sesResult.source), expectedSource)
	}

	if len(sesResult.to) != 1 || sesResult.to[0] != "recipient@example.com" {
		t.Errorf("Chained to = %v, want %v", sesResult.to, []string{"recipient@example.com"})
	}

	if len(sesResult.cc) != 1 || sesResult.cc[0] != "cc@example.com" {
		t.Errorf("Chained cc = %v, want %v", sesResult.cc, []string{"cc@example.com"})
	}

	if sesResult.subject == nil || *sesResult.subject != "Test Subject" {
		t.Errorf("Chained subject = %v, want %v", aws.ToString(sesResult.subject), "Test Subject")
	}

	if sesResult.body == nil || *sesResult.body != "Test Body" {
		t.Errorf("Chained body = %v, want %v", aws.ToString(sesResult.body), "Test Body")
	}
}

func TestSESCharSet(t *testing.T) {
	if CharSet != "UTF-8" {
		t.Errorf("CharSet = %v, want %v", CharSet, "UTF-8")
	}
}