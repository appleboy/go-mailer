# go-mailer

[![Go Report Card](https://goreportcard.com/badge/github.com/appleboy/go-mailer)](https://goreportcard.com/report/github.com/appleboy/go-mailer)
[![GoDoc](https://godoc.org/github.com/appleboy/go-mailer?status.svg)](https://godoc.org/github.com/appleboy/go-mailer)

[English](README.md) | [繁體中文](README.zh-TW.md) | [简体中文](README.zh-CN.md)

A unified email sending package for Go that supports multiple email service providers with a simple, consistent API.

## Table of Contents

- [go-mailer](#go-mailer)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Supported Providers](#supported-providers)
  - [Installation](#installation)
  - [Quick Start](#quick-start)
    - [SMTP Configuration](#smtp-configuration)
    - [Amazon SES Configuration](#amazon-ses-configuration)
  - [Configuration](#configuration)
    - [Config Structure](#config-structure)
    - [SMTP Configuration Options](#smtp-configuration-options)
      - [Port Settings](#port-settings)
      - [Common SMTP Providers](#common-smtp-providers)
        - [Gmail](#gmail)
        - [Outlook/Hotmail](#outlookhotmail)
        - [Yahoo](#yahoo)
    - [Amazon SES Setup](#amazon-ses-setup)
  - [API Reference](#api-reference)
    - [Mail Interface](#mail-interface)
    - [Methods](#methods)
      - [From(name, address string) Mail](#fromname-address-string-mail)
      - [To(addresses ...string) Mail](#toaddresses-string-mail)
      - [Cc(addresses ...string) Mail](#ccaddresses-string-mail)
      - [Subject(subject string) Mail](#subjectsubject-string-mail)
      - [Body(body string) Mail](#bodybody-string-mail)
      - [Send() (interface{}, error)](#send-interface-error)
  - [Advanced Usage](#advanced-usage)
    - [Using Global Client](#using-global-client)
    - [Error Handling](#error-handling)
    - [Multiple Recipients](#multiple-recipients)
  - [Dependencies](#dependencies)
  - [Requirements](#requirements)
  - [License](#license)
  - [Contributing](#contributing)
  - [Support](#support)

## Features

- 🚀 **Multiple Providers**: Support for SMTP and Amazon SES
- 🔧 **Unified Interface**: Single API for all email providers
- 📧 **Rich Email Features**: HTML/Text content, CC/BCC, multiple recipients
- ⚙️ **Easy Configuration**: Simple configuration structure
- 🔐 **Secure**: Built-in SSL/TLS support for SMTP
- 📊 **Logging**: Integrated with zerolog for structured logging

## Supported Providers

- **SMTP**: Standard SMTP servers (Gmail, Outlook, custom servers)
- **Amazon SES**: AWS Simple Email Service

## Installation

```bash
go get github.com/appleboy/go-mailer
```

## Quick Start

### SMTP Configuration

```go
package main

import (
    "log"
    "github.com/appleboy/go-mailer"
)

func main() {
    // Configure SMTP settings
    config := mailer.Config{
        Driver:   "smtp",
        Host:     "smtp.gmail.com",
        Port:     "587",
        Username: "your-email@gmail.com",
        Password: "your-app-password",
    }

    // Create email engine
    engine, err := mailer.NewEngine(config)
    if err != nil {
        log.Fatal(err)
    }

    // Send email
    _, err = engine.
        From("John Doe", "john@example.com").
        To("recipient@example.com", "another@example.com").
        Cc("cc@example.com").
        Subject("Hello from go-mailer!").
        Body("<h1>Hello World!</h1><p>This is a test email.</p>").
        Send()

    if err != nil {
        log.Fatal(err)
    }

    log.Println("Email sent successfully!")
}
```

### Amazon SES Configuration

```go
package main

import (
    "log"
    "github.com/appleboy/go-mailer"
)

func main() {
    // Configure SES settings
    config := mailer.Config{
        Driver: "ses",
        Region: "us-west-2", // Your AWS region
    }

    // Create email engine
    engine, err := mailer.NewEngine(config)
    if err != nil {
        log.Fatal(err)
    }

    // Send email
    _, err = engine.
        From("Sender Name", "verified-sender@example.com").
        To("recipient@example.com").
        Subject("Hello from SES!").
        Body("<h1>Hello from Amazon SES!</h1>").
        Send()

    if err != nil {
        log.Fatal(err)
    }

    log.Println("Email sent via SES successfully!")
}
```

## Configuration

### Config Structure

```go
type Config struct {
    Host     string // SMTP host (required for SMTP driver)
    Port     string // SMTP port (required for SMTP driver)
    Username string // SMTP username
    Password string // SMTP password
    Driver   string // Email driver: "smtp" or "ses"
    Region   string // AWS region (required for SES driver)
}
```

### SMTP Configuration Options

#### Port Settings

- **Port 25**: Plain SMTP (no encryption)
- **Port 465**: SMTP with SSL encryption
- **Port 587**: SMTP with TLS encryption (recommended)

#### Common SMTP Providers

##### Gmail

```go
config := mailer.Config{
    Driver:   "smtp",
    Host:     "smtp.gmail.com",
    Port:     "587",
    Username: "your-email@gmail.com",
    Password: "your-app-password", // Use App Password, not regular password
}
```

##### Outlook/Hotmail

```go
config := mailer.Config{
    Driver:   "smtp",
    Host:     "smtp-mail.outlook.com",
    Port:     "587",
    Username: "your-email@outlook.com",
    Password: "your-password",
}
```

##### Yahoo

```go
config := mailer.Config{
    Driver:   "smtp",
    Host:     "smtp.mail.yahoo.com",
    Port:     "587",
    Username: "your-email@yahoo.com",
    Password: "your-app-password",
}
```

### Amazon SES Setup

For SES, ensure your AWS credentials are configured via:

- AWS credentials file (`~/.aws/credentials`)
- Environment variables (`AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`)
- IAM roles (for EC2/Lambda)

Required SES permissions:

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "ses:SendEmail",
                "ses:SendRawEmail"
            ],
            "Resource": "*"
        }
    ]
}
```

## API Reference

### Mail Interface

All email operations use the `Mail` interface:

```go
type Mail interface {
    From(name, address string) Mail
    To(addresses ...string) Mail
    Cc(addresses ...string) Mail
    Subject(subject string) Mail
    Body(body string) Mail
    Send() (interface{}, error)
}
```

### Methods

#### From(name, address string) Mail

Set the sender information.

- `name`: Sender's display name (optional, can be empty string)
- `address`: Sender's email address

#### To(addresses ...string) Mail

Add recipient email addresses. Can be called multiple times or with multiple addresses.

#### Cc(addresses ...string) Mail

Add CC (Carbon Copy) recipients. Can be called multiple times or with multiple addresses.

#### Subject(subject string) Mail

Set the email subject line.

#### Body(body string) Mail

Set the email body content. Supports HTML formatting.

#### Send() (interface{}, error)

Send the email and return the result or error.

## Advanced Usage

### Using Global Client

You can set a global client for convenience:

```go
// Initialize global client
_, err := mailer.NewEngine(config)
if err != nil {
    log.Fatal(err)
}

// Use global client
_, err = mailer.Client.
    From("Sender", "sender@example.com").
    To("recipient@example.com").
    Subject("Test").
    Body("Hello!").
    Send()
```

### Error Handling

```go
engine, err := mailer.NewEngine(config)
if err != nil {
    switch {
    case strings.Contains(err.Error(), "driver is required"):
        log.Fatal("Email driver not specified")
    case strings.Contains(err.Error(), "host and port are required"):
        log.Fatal("SMTP configuration incomplete")
    case strings.Contains(err.Error(), "region is required"):
        log.Fatal("SES region not specified")
    default:
        log.Fatal("Configuration error:", err)
    }
}

result, err := engine.From("test", "test@example.com").
    To("recipient@example.com").
    Subject("Test").
    Body("Hello").
    Send()

if err != nil {
    log.Printf("Failed to send email: %v", err)
    return
}

log.Printf("Email sent successfully: %+v", result)
```

### Multiple Recipients

```go
// Method 1: Multiple calls
engine.To("user1@example.com").
    To("user2@example.com").
    Cc("cc1@example.com").
    Cc("cc2@example.com")

// Method 2: Multiple arguments
engine.To("user1@example.com", "user2@example.com").
    Cc("cc1@example.com", "cc2@example.com")

// Method 3: Mixed approach
engine.To("user1@example.com", "user2@example.com").
    To("user3@example.com").
    Cc("manager@example.com")
```

## Dependencies

- [go-simple-mail/v2](https://github.com/xhit/go-simple-mail) - SMTP client
- [aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) - AWS SES client
- [zerolog](https://github.com/rs/zerolog) - Structured logging

## Requirements

- Go 1.24.0 or higher

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Support

If you have any questions or need help, please:

- Open an issue on [GitHub](https://github.com/appleboy/go-mailer/issues)
- Check the [documentation](https://godoc.org/github.com/appleboy/go-mailer)
