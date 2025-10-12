package mailer

import (
	"testing"
)

func TestNewEngine(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid SMTP config",
			config: Config{
				Driver:   "smtp",
				Host:     "smtp.gmail.com",
				Port:     "587",
				Username: "test@gmail.com",
				Password: "password",
			},
			wantErr: false,
		},
		{
			name: "valid SES config",
			config: Config{
				Driver: "ses",
				Region: "us-east-1",
			},
			wantErr: false,
		},
		{
			name: "missing driver",
			config: Config{
				Host: "smtp.gmail.com",
				Port: "587",
			},
			wantErr: true,
			errMsg:  "driver is required",
		},
		{
			name: "SMTP missing host",
			config: Config{
				Driver: "smtp",
				Port:   "587",
			},
			wantErr: true,
			errMsg:  "SMTP host and port are required",
		},
		{
			name: "SMTP missing port",
			config: Config{
				Driver: "smtp",
				Host:   "smtp.gmail.com",
			},
			wantErr: true,
			errMsg:  "SMTP host and port are required",
		},
		{
			name: "SES missing region",
			config: Config{
				Driver: "ses",
			},
			wantErr: true,
			errMsg:  "SES region is required",
		},
		{
			name: "unsupported driver",
			config: Config{
				Driver: "sendgrid",
			},
			wantErr: true,
			errMsg:  "unsupported email driver: sendgrid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mail, err := NewEngine(tt.config)

			if tt.wantErr {
				if err == nil {
					t.Errorf("NewEngine() expected error but got none")
					return
				}
				if err.Error() != tt.errMsg {
					t.Errorf("NewEngine() error = %v, wantErr %v", err.Error(), tt.errMsg)
				}
				return
			}

			if err != nil {
				t.Errorf("NewEngine() unexpected error = %v", err)
				return
			}

			if mail == nil {
				t.Errorf("NewEngine() returned nil mail instance")
			}
		})
	}
}

func TestConfig(t *testing.T) {
	config := Config{
		Host:     "smtp.gmail.com",
		Port:     "587",
		Username: "test@example.com",
		Password: "password123",
		Driver:   "smtp",
		Region:   "us-east-1",
	}

	if config.Host != "smtp.gmail.com" {
		t.Errorf("Config.Host = %v, want %v", config.Host, "smtp.gmail.com")
	}

	if config.Port != "587" {
		t.Errorf("Config.Port = %v, want %v", config.Port, "587")
	}

	if config.Username != "test@example.com" {
		t.Errorf("Config.Username = %v, want %v", config.Username, "test@example.com")
	}

	if config.Password != "password123" {
		t.Errorf("Config.Password = %v, want %v", config.Password, "password123")
	}

	if config.Driver != "smtp" {
		t.Errorf("Config.Driver = %v, want %v", config.Driver, "smtp")
	}

	if config.Region != "us-east-1" {
		t.Errorf("Config.Region = %v, want %v", config.Region, "us-east-1")
	}
}