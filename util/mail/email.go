package mail

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/smtp"
	"os"
)

var (
	host     = "smtp.qcloudmail.com"
	port     = 465
	email    = os.Getenv("EMAIL_MAILBOX")
	password = os.Getenv("EMAIL_PASSWORD")
)

func SendEmail(mailbox string, content string) error {
	header := make(map[string]string)
	header["From"] = "Fuever " + "<" + email + ">"
	header["To"] = mailbox
	header["Subject"] = "verification"
	header["Content-Type"] = "text/plain; charset=UTF-8"
	body := content
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body
	auth := smtp.PlainAuth(
		"",
		email,
		password,
		host,
	)
	err := SendMailWithTLS(
		fmt.Sprintf("%s:%d", host, port),
		auth,
		email,
		[]string{mailbox},
		[]byte(message),
	)
	if err != nil {
		fmt.Println("error occur: ", err)
	} else {
		fmt.Println("send successful")
	}
	return err
}

func Dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		log.Println("tls.Dial Error:", err)
		return nil, err
	}

	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

func SendMailWithTLS(addr string, auth smtp.Auth, from string,
	to []string, msg []byte) (err error) {
	//create smtp client
	c, err := Dial(addr)
	if err != nil {
		log.Println("Create smtp client error:", err)
		return err
	}
	defer c.Close()
	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				log.Println("Error during AUTH", err)
				return err
			}
		}
	}
	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
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
