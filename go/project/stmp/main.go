package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/smtp"
	"os"
)

// 邮箱服务器配置信息
type configInfo struct {
	smtpAddr string
	smtpPort string
	secret   string
}

// 邮件内容信息
type emailContent struct {
	fromAddr    string
	contentType string
	theme       string
	message     string
	toAddr      []string
}

func sendEmail(cfg *configInfo, e *emailContent) error {
	// 拼接smtp服务器地址
	smtpAddr := cfg.smtpAddr + ":" + cfg.smtpPort
	// 认证信息
	//auth := smtp.PlainAuth("", e.fromAddr, cfg.secret, cfg.smtpAddr)
	// 配置邮件内容类型
	if e.contentType == "html" {
		e.contentType = "Content-Type: text/html; charset=UTF-8"
	} else {
		e.contentType = "Content-Type: text/plain; charset=UTF-8"
	}
	// 当有多个收件人
	for _, to := range e.toAddr {
		msg := []byte("To: " + to + "\r\n" +
			"From: " + e.fromAddr + "\r\n" +
			"Subject: " + e.theme + "\r\n" +
			e.contentType + "\r\n\r\n" +
			"<html><h1>" + e.message + "</h1></html>")

		c, err := smtp.Dial(smtpAddr)
		if err != nil {
			return err
		}

		if err = c.Mail(e.fromAddr); err != nil {
			return err
		}
		for _, addr := range e.toAddr {
			if err = c.Rcpt(addr); err != nil {
				return err
			}
		}
		w, err := c.Data()
		if err != nil {
			return err
		}
		_, err = w.Write(msg)
		if err != nil {
			return err
		}
		err = w.Close()
		if err != nil {
			return err
		}
		return c.Quit()

	}
	return nil
}

var (
	serveEmail     string
	serveEmailPort string
	fromEmail      string
	toEmail        string
	fileName       string
	plain          string
)

func main() {
	flag.StringVar(&serveEmail, "s", "mail.snapmail.cc", "发送方")
	flag.StringVar(&serveEmailPort, "sp", "25", "发送方")
	flag.StringVar(&fromEmail, "l", "alaugbicv@snapmail.cc", "发送方")
	flag.StringVar(&toEmail, "r", "ablo804700@snapmail.cc", "接收方")
	flag.StringVar(&fileName, "f", "", "文件名")
	flag.StringVar(&plain, "t", "test message", "文件名")
	flag.Parse()

	if _, err := os.Stat(fileName); err != nil {
		log.Fatal(err)
	} else {
		f, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		b, err := io.ReadAll(f)
		if err != nil {
			log.Fatal(err)
		}
		plain = string(b)
	}

	// 收集配置信息
	config := configInfo{
		// smtp服务器地址
		smtpAddr: serveEmail,
		// smtp服务器密钥
		secret: "",
		// smtp服务器端口
		smtpPort: serveEmailPort,
	}
	// 收集邮件内容
	content := emailContent{
		// 发件人
		fromAddr: fromEmail,
		// 收件人(可有多个)
		toAddr: []string{toEmail},
		// 邮件格式
		contentType: "text",
		// 邮件主题
		theme: "test",
		// 邮件内容
		message: plain,
	}
	// 发送邮件
	err := sendEmail(&config, &content)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("发送成功")
	}
}
