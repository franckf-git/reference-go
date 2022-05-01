package main

import "github.com/go-gomail/gomail"

func main() {

	m := gomail.NewMessage()
	m.SetHeader("From", "from@example.com")
	m.SetHeader("To", "to@example.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/plain", "Hello!")

	d := gomail.Dialer{Host: "192.168.1.78", Port: 1025}
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
