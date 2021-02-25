package messages

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Email struct {
	To string
	From string
	Date time.Time
	Title string
	Content string
}

// CmdlineEmail asks user to input an email from the command line
func CmdlineEmail() Email{
	var email Email

	// Prompt user and take user inputs
	fmt.Println("Enter sender's name: ")
	email.From = takeUserInput()
	fmt.Println("Enter receiver's name")
	email.To = takeUserInput()
	fmt.Println("Enter email title:")
	email.Title = takeUserInput()
	fmt.Println("Enter email content:")
	email.Content = takeUserInput()

	email.Date = time.Now()

	return email
}

// helper function for CmdlineEmail
func takeUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// PrintEmail prints fields of the Email struct
func PrintEmail(email Email ) {
	fmt.Println("Printing email...")
	fmt.Println("From: " + email.From)
	fmt.Println("Date: " + email.Date.Format("January 2, 2006"))
	fmt.Println("Title: " + email.Title)
	fmt.Println("Content: " + email.Content + "\n")
}