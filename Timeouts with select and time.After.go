//Why Use Timeouts?
//Sometimes you’re waiting for data on a channel, but what if it never comes?
//You don’t want your program to hang forever — use a timeout to escape!
//syntax
/*
select {
case msg := <-ch:
    fmt.Println("Received:", msg)
case <-time.After(2 * time.Second):
    fmt.Println("⏰ Timeout! No response in 2 seconds.")
}*/

//api call with timeout using select

package main

import (
    "fmt"
    "time"
)

// Simulate API call that takes variable time
func slowAPI(ch chan string) {
    time.Sleep(5 * time.Second) // Simulate slow response
    ch <- "✅ API response received"
}

func main() {
    ch := make(chan string)

    go slowAPI(ch)

    fmt.Println("Calling API... waiting for response (timeout = 3s)")

    select {
    case res := <-ch:
        fmt.Println(res)
    case <-time.After(3 * time.Second):
        fmt.Println("⏰ Timeout! API is taking too long.")
    }
}

/* Real-World Uses
API calls that shouldn't wait forever

IoT devices waiting for sensor values

Chat apps waiting for messages from peers

Game servers waiting for a player move

*/