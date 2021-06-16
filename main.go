package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
)

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "template.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		email := r.FormValue("email")
		title := r.FormValue("title")
		content := r.FormValue("content")
		sendMail(email, title, content)
		fmt.Fprintf(w, "Name = %s\n", email)
		fmt.Fprintf(w, "Address = %s\n", title)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	http.HandleFunc("/", hello)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}

func sendMail(email string, title string, content string) {
	// Sender data.
	from := "noreply@daugia.io"
	password := "abcD123$"

	// Receiver email address.
	to := []string{
		email,
	}

	// smtp server configuration.
	smtpHost := "smtp.yandex.com"
	smtpPort := "587"

	// Message.
	message := []byte("To: " + email + "\r\n" +
		"Subject:" + title + "\r\n" +
		"\r\n" +
		content + "\r\n")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
