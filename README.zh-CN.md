# go-mailer

[![Trivy](https://github.com/appleboy/go-mailer/actions/workflows/trivy.yml/badge.svg)](https://github.com/appleboy/go-mailer/actions/workflows/trivy.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/appleboy/go-mailer)](https://goreportcard.com/report/github.com/appleboy/go-mailer)
[![GoDoc](https://godoc.org/github.com/appleboy/go-mailer?status.svg)](https://godoc.org/github.com/appleboy/go-mailer)

[English](README.md) | [繁體中文](README.zh-TW.md) | [简体中文](README.zh-CN.md)

一个统一的 Go 语言邮件发送包，支持多个邮件服务提供商，提供简单且一致的 API。

## 目录

- [go-mailer](#go-mailer)
  - [目录](#目录)
  - [功能特性](#功能特性)
  - [支持的服务提供商](#支持的服务提供商)
  - [安装](#安装)
  - [快速开始](#快速开始)
    - [SMTP 配置](#smtp-配置)
    - [Amazon SES 配置](#amazon-ses-配置)
  - [配置](#配置)
    - [配置结构](#配置结构)
    - [SMTP 配置选项](#smtp-配置选项)
      - [端口设置](#端口设置)
      - [常见 SMTP 服务提供商](#常见-smtp-服务提供商)
        - [Gmail](#gmail)
        - [Outlook/Hotmail](#outlookhotmail)
        - [Yahoo](#yahoo)
    - [Amazon SES 设置](#amazon-ses-设置)
  - [API 参考](#api-参考)
    - [Mail 接口](#mail-接口)
    - [方法](#方法)
      - [From(name, address string) Mail](#fromname-address-string-mail)
      - [To(addresses ...string) Mail](#toaddresses-string-mail)
      - [Cc(addresses ...string) Mail](#ccaddresses-string-mail)
      - [Subject(subject string) Mail](#subjectsubject-string-mail)
      - [Body(body string) Mail](#bodybody-string-mail)
      - [Send() (interface{}, error)](#send-interface-error)
  - [高级用法](#高级用法)
    - [使用全局客户端](#使用全局客户端)
    - [错误处理](#错误处理)
    - [多收件人](#多收件人)
  - [依赖](#依赖)
  - [要求](#要求)
  - [许可证](#许可证)
  - [贡献](#贡献)
  - [支持](#支持)

## 功能特性

- 🚀 **多服务提供商**：支持 SMTP 和 Amazon SES
- 🔧 **统一接口**：所有邮件提供商使用单一 API
- 📧 **丰富的邮件功能**：HTML/文本内容、抄送/密送、多收件人
- ⚙️ **简单配置**：简单的配置结构
- 🔐 **安全**：内置 SSL/TLS 支持 SMTP
- 📊 **日志记录**：集成 zerolog 结构化日志记录

## 支持的服务提供商

- **SMTP**：标准 SMTP 服务器（Gmail、Outlook、自定义服务器）
- **Amazon SES**：AWS 简单邮件服务

## 安装

```bash
go get github.com/appleboy/go-mailer
```

## 快速开始

### SMTP 配置

```go
package main

import (
    "log"
    "github.com/appleboy/go-mailer"
)

func main() {
    // 配置 SMTP 设置
    config := mailer.Config{
        Driver:   "smtp",
        Host:     "smtp.gmail.com",
        Port:     "587",
        Username: "your-email@gmail.com",
        Password: "your-app-password",
    }

    // 创建邮件引擎
    engine, err := mailer.NewEngine(config)
    if err != nil {
        log.Fatal(err)
    }

    // 发送邮件
    _, err = engine.
        From("John Doe", "john@example.com").
        To("recipient@example.com", "another@example.com").
        Cc("cc@example.com").
        Subject("来自 go-mailer 的问候！").
        Body("<h1>Hello World!</h1><p>这是一封测试邮件。</p>").
        Send()

    if err != nil {
        log.Fatal(err)
    }

    log.Println("邮件发送成功！")
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
    // 配置 SES 设置
    config := mailer.Config{
        Driver: "ses",
        Region: "us-west-2", // 您的 AWS 区域
    }

    // 创建邮件引擎
    engine, err := mailer.NewEngine(config)
    if err != nil {
        log.Fatal(err)
    }

    // 发送邮件
    _, err = engine.
        From("发送者姓名", "verified-sender@example.com").
        To("recipient@example.com").
        Subject("来自 SES 的问候！").
        Body("<h1>来自 Amazon SES 的问候！</h1>").
        Send()

    if err != nil {
        log.Fatal(err)
    }

    log.Println("通过 SES 成功发送邮件！")
}
```

## 配置

### 配置结构

```go
type Config struct {
    Host     string // SMTP 主机（SMTP 驱动程序必需）
    Port     string // SMTP 端口（SMTP 驱动程序必需）
    Username string // SMTP 用户名
    Password string // SMTP 密码
    Driver   string // 邮件驱动程序："smtp" 或 "ses"
    Region   string // AWS 区域（SES 驱动程序必需）
}
```

### SMTP 配置选项

#### 端口设置

- **端口 25**：纯 SMTP（无加密）
- **端口 465**：SMTP with SSL 加密
- **端口 587**：SMTP with TLS 加密（推荐）

#### 常见 SMTP 服务提供商

##### Gmail

```go
config := mailer.Config{
    Driver:   "smtp",
    Host:     "smtp.gmail.com",
    Port:     "587",
    Username: "your-email@gmail.com",
    Password: "your-app-password", // 使用应用程序密码，而非普通密码
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

### Amazon SES 设置

对于 SES，请确保您的 AWS 凭证已通过以下方式配置：

- AWS 凭证文件 (`~/.aws/credentials`)
- 环境变量 (`AWS_ACCESS_KEY_ID`、`AWS_SECRET_ACCESS_KEY`)
- IAM 角色（适用于 EC2/Lambda）

所需的 SES 权限：

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

## API 参考

### Mail 接口

所有邮件操作都使用 `Mail` 接口：

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

设置发送者信息。

- `name`：发送者的显示姓名（可选，可以是空字符串）
- `address`：发送者的邮件地址

#### To(addresses ...string) Mail

添加收件人邮件地址。可以多次调用或使用多个地址。

#### Cc(addresses ...string) Mail

添加抄送（Carbon Copy）收件人。可以多次调用或使用多个地址。

#### Subject(subject string) Mail

设置邮件主题。

#### Body(body string) Mail

设置邮件内容。支持 HTML 格式。

#### Send() (interface{}, error)

发送邮件并返回结果或错误。

## 高级用法

### 使用全局客户端

您可以设置全局客户端以便于使用：

```go
// 初始化全局客户端
_, err := mailer.NewEngine(config)
if err != nil {
    log.Fatal(err)
}

// 使用全局客户端
_, err = mailer.Client.
    From("发送者", "sender@example.com").
    To("recipient@example.com").
    Subject("测试").
    Body("您好！").
    Send()
```

### 错误处理

```go
engine, err := mailer.NewEngine(config)
if err != nil {
    switch {
    case strings.Contains(err.Error(), "driver is required"):
        log.Fatal("未指定邮件驱动程序")
    case strings.Contains(err.Error(), "host and port are required"):
        log.Fatal("SMTP 配置不完整")
    case strings.Contains(err.Error(), "region is required"):
        log.Fatal("未指定 SES 区域")
    default:
        log.Fatal("配置错误：", err)
    }
}

result, err := engine.From("test", "test@example.com").
    To("recipient@example.com").
    Subject("测试").
    Body("您好").
    Send()

if err != nil {
    log.Printf("发送邮件失败：%v", err)
    return
}

log.Printf("邮件发送成功：%+v", result)
```

### 多收件人

```go
// 方法 1：多次调用
engine.To("user1@example.com").
    To("user2@example.com").
    Cc("cc1@example.com").
    Cc("cc2@example.com")

// 方法 2：多个参数
engine.To("user1@example.com", "user2@example.com").
    Cc("cc1@example.com", "cc2@example.com")

// 方法 3：混合方式
engine.To("user1@example.com", "user2@example.com").
    To("user3@example.com").
    Cc("manager@example.com")
```

## 依赖

- [go-simple-mail/v2](https://github.com/xhit/go-simple-mail) - SMTP 客户端
- [aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) - AWS SES 客户端
- [zerolog](https://github.com/rs/zerolog) - 结构化日志记录

## 要求

- Go 1.24.0 或更高版本

## 许可证

本项目采用 MIT 许可证 - 详情请见 [LICENSE](LICENSE) 文件。

## 贡献

1. Fork 此仓库
2. 创建您的功能分支 (`git checkout -b feature/amazing-feature`)
3. 提交您的更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 打开 Pull Request

## 支持

如果您有任何问题或需要帮助，请：

- 在 [GitHub](https://github.com/appleboy/go-mailer/issues) 上开启 issue
- 查看[文档](https://godoc.org/github.com/appleboy/go-mailer)
