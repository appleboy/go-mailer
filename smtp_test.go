package mailer

import (
	"testing"
)

func TestSMTPEngine(t *testing.T) {
	smtp, err := SMTPEngine("smtp.gmail.com", "587", "test@gmail.com", "password")
	if err != nil {
		t.Errorf("SMTPEngine() unexpected error = %v", err)
	}

	if smtp == nil {
		t.Errorf("SMTPEngine() returned nil")
	}

	if smtp.host != "smtp.gmail.com" {
		t.Errorf("SMTPEngine() host = %v, want %v", smtp.host, "smtp.gmail.com")
	}

	if smtp.port != "587" {
		t.Errorf("SMTPEngine() port = %v, want %v", smtp.port, "587")
	}

	if smtp.username != "test@gmail.com" {
		t.Errorf("SMTPEngine() username = %v, want %v", smtp.username, "test@gmail.com")
	}

	if smtp.password != "password" {
		t.Errorf("SMTPEngine() password = %v, want %v", smtp.password, "password")
	}
}

func TestSMTPFrom(t *testing.T) {
	smtp := &SMTP{}

	result := smtp.From("John Doe", "john@example.com")

	// Type assertion to access internal fields
	smtpResult, ok := result.(SMTP)
	if !ok {
		t.Errorf("From() did not return SMTP type")
	}

	if smtpResult.from.Name != "John Doe" {
		t.Errorf("From() name = %v, want %v", smtpResult.from.Name, "John Doe")
	}

	if smtpResult.from.Address != "john@example.com" {
		t.Errorf("From() address = %v, want %v", smtpResult.from.Address, "john@example.com")
	}
}

func TestSMTPTo(t *testing.T) {
	smtp := &SMTP{}

	result := smtp.To("recipient1@example.com", "recipient2@example.com")

	smtpResult, ok := result.(SMTP)
	if !ok {
		t.Errorf("To() did not return SMTP type")
	}

	expectedTo := []string{"recipient1@example.com", "recipient2@example.com"}
	if len(smtpResult.to) != len(expectedTo) {
		t.Errorf("To() length = %v, want %v", len(smtpResult.to), len(expectedTo))
	}

	for i, addr := range expectedTo {
		if smtpResult.to[i] != addr {
			t.Errorf("To() address[%d] = %v, want %v", i, smtpResult.to[i], addr)
		}
	}
}

func TestSMTPCc(t *testing.T) {
	smtp := &SMTP{}

	result := smtp.Cc("cc1@example.com", "cc2@example.com")

	smtpResult, ok := result.(SMTP)
	if !ok {
		t.Errorf("Cc() did not return SMTP type")
	}

	expectedCc := []string{"cc1@example.com", "cc2@example.com"}
	if len(smtpResult.cc) != len(expectedCc) {
		t.Errorf("Cc() length = %v, want %v", len(smtpResult.cc), len(expectedCc))
	}

	for i, addr := range expectedCc {
		if smtpResult.cc[i] != addr {
			t.Errorf("Cc() address[%d] = %v, want %v", i, smtpResult.cc[i], addr)
		}
	}
}

func TestSMTPSubject(t *testing.T) {
	smtp := &SMTP{}

	result := smtp.Subject("Test Subject")

	smtpResult, ok := result.(SMTP)
	if !ok {
		t.Errorf("Subject() did not return SMTP type")
	}

	if smtpResult.subject != "Test Subject" {
		t.Errorf("Subject() = %v, want %v", smtpResult.subject, "Test Subject")
	}
}

func TestSMTPBody(t *testing.T) {
	smtp := &SMTP{}

	result := smtp.Body("<h1>Test Body</h1>")

	smtpResult, ok := result.(SMTP)
	if !ok {
		t.Errorf("Body() did not return SMTP type")
	}

	if smtpResult.body != "<h1>Test Body</h1>" {
		t.Errorf("Body() = %v, want %v", smtpResult.body, "<h1>Test Body</h1>")
	}
}

func TestSMTPChaining(t *testing.T) {
	smtp := &SMTP{}

	result := smtp.
		From("John Doe", "john@example.com").
		To("recipient@example.com").
		Cc("cc@example.com").
		Subject("Test Subject").
		Body("Test Body")

	smtpResult, ok := result.(SMTP)
	if !ok {
		t.Errorf("Method chaining did not return SMTP type")
	}

	// Verify all fields are set correctly
	if smtpResult.from.Name != "John Doe" {
		t.Errorf("Chained from name = %v, want %v", smtpResult.from.Name, "John Doe")
	}

	if smtpResult.from.Address != "john@example.com" {
		t.Errorf("Chained from address = %v, want %v", smtpResult.from.Address, "john@example.com")
	}

	if len(smtpResult.to) != 1 || smtpResult.to[0] != "recipient@example.com" {
		t.Errorf("Chained to = %v, want %v", smtpResult.to, []string{"recipient@example.com"})
	}

	if len(smtpResult.cc) != 1 || smtpResult.cc[0] != "cc@example.com" {
		t.Errorf("Chained cc = %v, want %v", smtpResult.cc, []string{"cc@example.com"})
	}

	if smtpResult.subject != "Test Subject" {
		t.Errorf("Chained subject = %v, want %v", smtpResult.subject, "Test Subject")
	}

	if smtpResult.body != "Test Body" {
		t.Errorf("Chained body = %v, want %v", smtpResult.body, "Test Body")
	}
}