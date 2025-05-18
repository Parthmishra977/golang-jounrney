//Interface Composition and Embedding in Go

 //What is Interface Composition?
//Interface composition means creating larger interfaces by combining smaller interfaces.

//This allows you to:

//Build modular, flexible interfaces

//Add behavior in layers

//Keep code clean and organized

//syntax example

/*
type Reader interface{
	read()
}
type writer interface{
	writer()
}
//combined interface
type Readerwriter interface{
	Reader
	writer
}*/

//Any type that implements both Read() and Write() methods will satisfy ReadWriter.

//Real-World Example: Multi-Channel Notification System


/*
We'll build a notification system where:

Each Notifier can either send an Email or SMS

A MultiNotifier sends both — using interface composition

The user selects which notification type to send
*/
/*
package main

import (
    "fmt"
)

// Simple interfaces
type EmailSender interface {
    SendEmail(message string)
}

type SMSSender interface {
    SendSMS(message string)
}

// Composed interface
type Notifier interface {
    EmailSender
    SMSSender
}

// Concrete struct
type NotificationService struct {
    Email string
    Phone string
}

// Implement SendEmail
func (n NotificationService) SendEmail(message string) {
    fmt.Printf("📧 Email sent to %s: %s\n", n.Email, message)
}

// Implement SendSMS
func (n NotificationService) SendSMS(message string) {
    fmt.Printf("📱 SMS sent to %s: %s\n", n.Phone, message)
}

func main() {
    var email, phone, message, mode string

    fmt.Print("Enter email address: ")
    fmt.Scanln(&email)

    fmt.Print("Enter phone number: ")
    fmt.Scanln(&phone)

    fmt.Print("Enter message to send: ")
    fmt.Scanln(&message)

    fmt.Print("Send via (email/sms/both): ")
    fmt.Scanln(&mode)

    notifier := NotificationService{Email: email, Phone: phone}

    switch mode {
    case "email":
        notifier.SendEmail(message)
    case "sms":
        notifier.SendSMS(message)
    case "both":
        var multiNotifier Notifier = notifier // Implicitly satisfies composed interface
        multiNotifier.SendEmail(message)
        multiNotifier.SendSMS(message)
    default:
        fmt.Println("Invalid option. Please choose email/sms/both.")
    }
}


Concept	Description
Interface composition	Combine small interfaces into larger ones
Reusability	Allows decoupling of logic (e.g., EmailSender)
Flexibility	Add/remove features without touching other parts
Real use	Notification, Payment gateways, I/O interfaces*/