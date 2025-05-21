// mini project: banking transaction processor

//Problem Statement
//A bank server processes money transfers from users. Some transactions are quick, some are delayed. You want to:

// Process transactions concurrently (using goroutines)

// Get updates through channels

// Handle timeouts (cancel if taking too long)

// Combine everything with select

//Features
//User inputs multiple transactions (amounts)

//Each transaction is processed in a goroutine

//If a transaction takes longer than 3 seconds, it's canceled

//	Final result shows successful and timed-out transactions


/*
package main

import (
    "fmt"
    "math/rand"
    "time"
)

// Simulates a transaction processor with random delay
func processTransaction(id int, amount int, ch chan string) {
    delay := time.Duration(rand.Intn(5)+1) * time.Second
    time.Sleep(delay) // Simulate processing time
    ch <- fmt.Sprintf("✅ Transaction #%d of ₹%d processed in %v", id, amount, delay)
}

func main() {
    rand.Seed(time.Now().UnixNano())

    var n int
    fmt.Print("Enter number of transactions: ")
    fmt.Scanln(&n)

    amounts := make([]int, n)

    for i := 0; i < n; i++ {
        fmt.Printf("Enter amount for transaction #%d: ₹", i+1)
        fmt.Scanln(&amounts[i])
    }

    fmt.Println("\n🔄 Processing transactions...\n")

    for i := 0; i < n; i++ {
        ch := make(chan string)
        go processTransaction(i+1, amounts[i], ch)

        select {
        case msg := <-ch:
            fmt.Println(msg)
        case <-time.After(3 * time.Second):
            fmt.Printf("❌ Transaction #%d of ₹%d failed: Timeout after 3s\n", i+1, amounts[i])
        }
    }

    fmt.Println("\n🏁 All transactions attempted.")
}
