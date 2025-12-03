# go-mailer

[![Trivy](https://github.com/appleboy/go-mailer/actions/workflows/trivy.yml/badge.svg)](https://github.com/appleboy/go-mailer/actions/workflows/trivy.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/appleboy/go-mailer)](https://goreportcard.com/report/github.com/appleboy/go-mailer)
[![GoDoc](https://godoc.org/github.com/appleboy/go-mailer?status.svg)](https://godoc.org/github.com/appleboy/go-mailer)

[English](README.md) | [繁體中文](README.zh-TW.md) | [简体中文](README.zh-CN.md)

一個統一的 Go 語言電子郵件發送套件，支援多個電子郵件服務提供者，並提供簡單且一致的 API。

## 目錄

- [go-mailer](#go-mailer)
  - [目錄](#目錄)
  - [功能特色](#功能特色)
  - [支援的服務提供者](#支援的服務提供者)
  - [安裝](#安裝)
  - [快速開始](#快速開始)
    - [SMTP 配置](#smtp-配置)
    - [Amazon SES 配置](#amazon-ses-配置)
  - [配置](#配置)
    - [配置結構](#配置結構)
    - [SMTP 配置選項](#smtp-配置選項)
      - [埠設定](#埠設定)
      - [常見 SMTP 服務提供者](#常見-smtp-服務提供者)
        - [Gmail](#gmail)
        - [Outlook/Hotmail](#outlookhotmail)
        - [Yahoo](#yahoo)
    - [Amazon SES 設定](#amazon-ses-設定)
  - [API 參考](#api-參考)
    - [Mail 介面](#mail-介面)
    - [方法](#方法)
      - [From(name, address string) Mail](#fromname-address-string-mail)
      - [To(addresses ...string) Mail](#toaddresses-string-mail)
      - [Cc(addresses ...string) Mail](#ccaddresses-string-mail)
      - [Subject(subject string) Mail](#subjectsubject-string-mail)
      - [Body(body string) Mail](#bodybody-string-mail)
      - [Send() (interface{}, error)](#send-interface-error)
  - [進階用法](#進階用法)
    - [使用全域客戶端](#使用全域客戶端)
    - [錯誤處理](#錯誤處理)
    - [多收件人](#多收件人)
  - [相依性](#相依性)
  - [需求](#需求)
  - [授權](#授權)
  - [貢獻](#貢獻)
  - [支援](#支援)

## 功能特色

- 🚀 **多服務提供者**：支援 SMTP 和 Amazon SES
- 🔧 **統一介面**：所有電子郵件提供者使用單一 API
- 📧 **豐富的電子郵件功能**：HTML/文字內容、副本/密件副本、多收件人
- ⚙️ **簡單配置**：簡單的配置結構
- 🔐 **安全**：內建 SSL/TLS 支援 SMTP
- 📊 **日誌記錄**：整合 zerolog 結構化日誌記錄

## 支援的服務提供者

- **SMTP**：標準 SMTP 伺服器（Gmail、Outlook、自定義伺服器）
- **Amazon SES**：AWS 簡易電子郵件服務

## 安裝

```bash
go get github.com/appleboy/go-mailer
```

## 快速開始

### SMTP 配置

```go
package main

import (
    "log"
    "github.com/appleboy/go-mailer"
)

func main() {
    // 配置 SMTP 設定
    config := mailer.Config{
        Driver:   "smtp",
        Host:     "smtp.gmail.com",
        Port:     "587",
        Username: "your-email@gmail.com",
        Password: "your-app-password",
    }

    // 建立電子郵件引擎
    engine, err := mailer.NewEngine(config)
    if err != nil {
        log.Fatal(err)
    }

    // 發送電子郵件
    _, err = engine.
        From("John Doe", "john@example.com").
        To("recipient@example.com", "another@example.com").
        Cc("cc@example.com").
        Subject("來自 go-mailer 的問候！").
        Body("<h1>Hello World!</h1><p>這是一封測試電子郵件。</p>").
        Send()

    if err != nil {
        log.Fatal(err)
    }

    log.Println("電子郵件發送成功！")
}
```

### Amazon SES 配置

```go
package main

import (
    "log"
    "github.com/appleboy/go-mailer"
)

func main() {
    // 配置 SES 設定
    config := mailer.Config{
        Driver: "ses",
        Region: "us-west-2", // 您的 AWS 區域
    }

    // 建立電子郵件引擎
    engine, err := mailer.NewEngine(config)
    if err != nil {
        log.Fatal(err)
    }

    // 發送電子郵件
    _, err = engine.
        From("發送者姓名", "verified-sender@example.com").
        To("recipient@example.com").
        Subject("來自 SES 的問候！").
        Body("<h1>來自 Amazon SES 的問候！</h1>").
        Send()

    if err != nil {
        log.Fatal(err)
    }

    log.Println("透過 SES 成功發送電子郵件！")
}
```

## 配置

### 配置結構

```go
type Config struct {
    Host     string // SMTP 主機（SMTP 驅動程式必需）
    Port     string // SMTP 埠（SMTP 驅動程式必需）
    Username string // SMTP 使用者名稱
    Password string // SMTP 密碼
    Driver   string // 電子郵件驅動程式："smtp" 或 "ses"
    Region   string // AWS 區域（SES 驅動程式必需）
}
```

### SMTP 配置選項

#### 埠設定

- **埠 25**：純 SMTP（無加密）
- **埠 465**：SMTP with SSL 加密
- **埠 587**：SMTP with TLS 加密（建議）

#### 常見 SMTP 服務提供者

##### Gmail

```go
config := mailer.Config{
    Driver:   "smtp",
    Host:     "smtp.gmail.com",
    Port:     "587",
    Username: "your-email@gmail.com",
    Password: "your-app-password", // 使用應用程式密碼，而非一般密碼
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

### Amazon SES 設定

對於 SES，請確保您的 AWS 憑證已透過以下方式配置：

- AWS 憑證檔案 (`~/.aws/credentials`)
- 環境變數 (`AWS_ACCESS_KEY_ID`、`AWS_SECRET_ACCESS_KEY`)
- IAM 角色（適用於 EC2/Lambda）

所需的 SES 權限：

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

## API 參考

### Mail 介面

所有電子郵件操作都使用 `Mail` 介面：

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

### 方法

#### From(name, address string) Mail

設定發送者資訊。

- `name`：發送者的顯示姓名（選擇性，可以是空字串）
- `address`：發送者的電子郵件地址

#### To(addresses ...string) Mail

新增收件人電子郵件地址。可以多次呼叫或使用多個地址。

#### Cc(addresses ...string) Mail

新增副本（Carbon Copy）收件人。可以多次呼叫或使用多個地址。

#### Subject(subject string) Mail

設定電子郵件主旨。

#### Body(body string) Mail

設定電子郵件內容。支援 HTML 格式。

#### Send() (interface{}, error)

發送電子郵件並返回結果或錯誤。

## 進階用法

### 使用全域客戶端

您可以設定全域客戶端以便於使用：

```go
// 初始化全域客戶端
_, err := mailer.NewEngine(config)
if err != nil {
    log.Fatal(err)
}

// 使用全域客戶端
_, err = mailer.Client.
    From("發送者", "sender@example.com").
    To("recipient@example.com").
    Subject("測試").
    Body("您好！").
    Send()
```

### 錯誤處理

```go
engine, err := mailer.NewEngine(config)
if err != nil {
    switch {
    case strings.Contains(err.Error(), "driver is required"):
        log.Fatal("未指定電子郵件驅動程式")
    case strings.Contains(err.Error(), "host and port are required"):
        log.Fatal("SMTP 配置不完整")
    case strings.Contains(err.Error(), "region is required"):
        log.Fatal("未指定 SES 區域")
    default:
        log.Fatal("配置錯誤：", err)
    }
}

result, err := engine.From("test", "test@example.com").
    To("recipient@example.com").
    Subject("測試").
    Body("您好").
    Send()

if err != nil {
    log.Printf("發送電子郵件失敗：%v", err)
    return
}

log.Printf("電子郵件發送成功：%+v", result)
```

### 多收件人

```go
// 方法 1：多次呼叫
engine.To("user1@example.com").
    To("user2@example.com").
    Cc("cc1@example.com").
    Cc("cc2@example.com")

// 方法 2：多個參數
engine.To("user1@example.com", "user2@example.com").
    Cc("cc1@example.com", "cc2@example.com")

// 方法 3：混合方式
engine.To("user1@example.com", "user2@example.com").
    To("user3@example.com").
    Cc("manager@example.com")
```

## 相依性

- [go-simple-mail/v2](https://github.com/xhit/go-simple-mail) - SMTP 客戶端
- [aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) - AWS SES 客戶端
- [zerolog](https://github.com/rs/zerolog) - 結構化日誌記錄

## 需求

- Go 1.24.0 或更高版本

## 授權

本專案採用 MIT 授權 - 詳情請見 [LICENSE](LICENSE) 檔案。

## 貢獻

1. Fork 此儲存庫
2. 建立您的功能分支 (`git checkout -b feature/amazing-feature`)
3. 提交您的變更 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 開啟 Pull Request

## 支援

如果您有任何問題或需要協助，請：

- 在 [GitHub](https://github.com/appleboy/go-mailer/issues) 上開啟 issue
- 查看[文件](https://godoc.org/github.com/appleboy/go-mailer)
