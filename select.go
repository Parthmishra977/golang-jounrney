//select Statement in Go

//Why Use select?
//When you have multiple channels and want to wait for any one of them to respond, select is the perfect tool.
/*
select {
case msg1 := <-ch1:
    // handle message from channel 1
case msg2 := <-ch2:
    // handle message from channel 2
default:
    // optional: handle if no channel is ready
}
*/
/*
 Real-World Example: Racing Download Servers
You are downloading a file from multiple mirror servers, and you want to pick the fastest response.
*/

package main
import(
	"fmt"
	"time"
	"math/rand"
)
func downloadfrom