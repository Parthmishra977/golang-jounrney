//A pointer is a variable that stores the memory address of another variable.

//In Go, you use * to denote a pointer type.

//Use & to get the address of a variable.

//Basic Syntax

//var x int = 10
//var p *int = &x


/*&x gives the address of x

*p gives the value at that address
*/

/*package main

import "fmt"

func main() {
    var x int = 42
    var p *int = &x

    fmt.Println("x:", x)     // Value
    fmt.Println("p:", p)     // Address
    fmt.Println("*p:", *p)   // Value via pointer
}
*/


//Update Price by Reference

/*package main

import "fmt"

func updatePrice(price *float64) {
    *price = *price * 0.9 // Apply 10% discount
}

func main() {
    price := 100.0
    updatePrice(&price)
    fmt.Println("Discounted Price:", price)
}*/


//Why Use Pointers?
//Improve performance (avoid copying large structs)

//Allow functions to modify arguments

//Implement shared state


//Pointers can be nil (i.e., not pointing to anything yet).

//var p *int
//fmt.Println(p) Output: <nil>
/*
package main

import "fmt"

type Config struct {
    Port int
}

func loadConfig(cfg *Config) {
    if cfg == nil {
        fmt.Println("No config provided!")
        return
    }
    fmt.Println("Running on port:", cfg.Port)
}

func main() {
    var c *Config = nil
    loadConfig(c)

    cfg := &Config{Port: 8080}
    loadConfig(cfg)
}
*/
/*
package main
import "fmt"

func changeVal(x int) {
    x = 100
}

func changeRef(x *int) {
    *x = 100
}

func main() {
    val := 10
    changeVal(val)       // Pass by value
    fmt.Println(val)     // Output: 10

    changeRef(&val)      // Pass by reference
    fmt.Println(val)     // Output: 100
}
*/

