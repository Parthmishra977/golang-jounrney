/*
package main

import (
    "fmt"
    "time"
)

// Function to simulate sending email
func sendEmail(email, message string, done chan string) {
    time.Sleep(2 * time.Second) // Simulate delay
    fmt.Printf("📧 Email sent to %s: %s\n", email, message)
    done <- "email"
}

// Function to simulate sending SMS
func sendSMS(phone, message string, done chan string) {
    time.Sleep(1 * time.Second) // Simulate delay
    fmt.Printf("📱 SMS sent to %s: %s\n", phone, message)
    done <- "sms"
}

func main() {
    var email, phone, message string

    fmt.Print("Enter email: ")
    fmt.Scanln(&email)

    fmt.Print("Enter phone number: ")
    fmt.Scanln(&phone)

    fmt.Print("Enter message: ")
    fmt.Scanln(&message)

    // Create channel
    done := make(chan string)

    // Start goroutines
    go sendEmail(email, message, done)
    go sendSMS(phone, message, done)

    // Wait for both to finish
    for i := 0; i < 2; i++ {
        status := <-done
        fmt.Println("✅ Finished sending:", status)
    }

    fmt.Println("📬 All notifications sent.")
}
/*
Real-World Uses of Goroutines
Sending notifications in background

Running APIs and DB calls in parallel

Real-time chat apps

Web scrapers (parallel downloads)

Games: Physics & AI running in parallel
*/