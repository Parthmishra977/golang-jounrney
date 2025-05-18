//What is a Channel?
//A channel is a typed pipe through which goroutines communicate with each other by sending and receiving values.

//You can think of it like a WhatsApp chat between two goroutines — one sends a message, the other receives it.
/*
ch := make(chan int)      // create channel
ch <- 5                   // send 5 into channel
x := <-ch                 // receive from channel
*/


//<- is the channel operator
//data moves in the direction of the arrow

//Real-World Example: Order Status Notifier

//You place an order and want updates on its progress (e.g., Confirmed, Packed, Shipped).
//Each step is handled by a different goroutine, which sends a message through a channel.
/*
package main

import (
    "fmt"
    "time"
)

// Simulate order steps using goroutines
func orderConfirmed(ch chan string) {
    time.Sleep(1 * time.Second)
    ch <- "✅ Order Confirmed"
}

func orderPacked(ch chan string) {
    time.Sleep(2 * time.Second)
    ch <- "📦 Order Packed"
}

func orderShipped(ch chan string) {
    time.Sleep(3 * time.Second)
    ch <- "🚚 Order Shipped"
}

func main() {
    var userName string
    fmt.Print("Enter your name: ")
    fmt.Scanln(&userName)

    fmt.Printf("\nHello %s, tracking your order...\n\n", userName)

    // Create a channel
    statusChan := make(chan string)

    // Start goroutines
    go orderConfirmed(statusChan)
    go orderPacked(statusChan)
    go orderShipped(statusChan)

    // Receive 3 messages
    for i := 0; i < 3; i++ {
        update := <-statusChan
        fmt.Println(update)
    }

    fmt.Println("\n🎉 Your order is on the way!")
}
*/

// Code with Comments: Parallel Sum Calculator Using Channels and Goroutines
/*
package main

import (
    "fmt"
)

// sum calculates the total of a slice of integers
// and sends the result into the provided channel.
func sum(nums []int, ch chan int) {
    total := 0
    for _, val := range nums {
        total += val // Add each number to total
    }
    ch <- total // Send the total to the channel
}

func main() {
    var n int
    fmt.Print("How many numbers do you want to add? ")
    fmt.Scanln(&n) // Get the number of inputs from the user

    // Create a slice to store the numbers
    numbers := make([]int, n)
    fmt.Println("Enter", n, "numbers:")

    // Take input numbers from user
    for i := 0; i < n; i++ {
        fmt.Printf("Enter number %d: ", i+1)
        fmt.Scanln(&numbers[i])
    }

    // Divide the numbers into two halves
    mid := n / 2

    // Create a channel for communication between goroutines
    ch := make(chan int)

    // Start goroutine to sum the first half of the slice
    go sum(numbers[:mid], ch)

    // Start goroutine to sum the second half of the slice
    go sum(numbers[mid:], ch)

    // Receive results from both goroutines
    sum1 := <-ch // Receive first result
    sum2 := <-ch // Receive second result

    // Output individual and final totals
    fmt.Printf("\nFirst half sum: %d\n", sum1)
    fmt.Printf("Second half sum: %d\n", sum2)
    fmt.Printf("Total sum: %d\n", sum1+sum2)
}
*/
