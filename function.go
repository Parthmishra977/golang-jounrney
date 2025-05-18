//A function is a reusable block of code that performs a specific task.

//In Go, functions are first-class citizens, meaning:

//You can assign them to variables.

//Pass them as arguments.

//Return them from other functions.

/*func functionName(parameter type) returnType {
    // function body
    return value
}*/

/*package main

import "fmt"

func greet(name string) string {
    return "Hello, " + name
}

func main() {
    message := greet("Alice")
    fmt.Println(message)
}*/

//Multiple Parameters and Return Values

//func add(a int, b int) int {
  //  return a + b
//}

/*func divide(dividend, divisor float64) (float64, string) {
    if divisor == 0 {
        return 0, "Cannot divide by zero"
    }
    return dividend / divisor, "Success"
}*/
/*
package main

import "fmt"

func add(x, y int) int {
    return x + y
}
func substract(x,y int)int{
	return x-y
}

func main() {
    fmt.Println("Add:", add(10, 5))
    fmt.Println("Subtract:", substract(10, 5))
}
*/
/*
package main

import "fmt"

func room(length, width float64) (area, perimeter float64) {
    area = length * width
    perimeter = 2 * (length + width)
    return
}

func main() {
    a, p := room(10.5, 5.0)
    fmt.Println("Area:", a)
    fmt.Println("Perimeter:", p)
}

*/

//Variadic Functions
//A function can accept a variable number of arguments using ....
/*
package main

import "fmt"

func sum(numbers ...int) int {
    total := 0
    for _, n := range numbers {
        total += n
    }
    return total
}


func totalCart(prices ...float64) float64 {
    total := 0.0
    for _, price := range prices {
        total += price
    }
    return total
}

func main() {
    total := totalCart(10.5, 20.0, 5.25)
    fmt.Println("Total cart amount:", total)
}
*/
//Function as a Value (Higher Order Functions)
//Functions can be assigned to variables and passed around.
//func operation(a int, b int, f func(int, int) int) int {
//  return f(a, b)
//}

/*
package main

import "fmt"

func discount(a, b int) int {
    return a - b
}

func bill(a, b int, calc func(int, int) int) int {
    return calc(a, b)
}

func main() {
    result := bill(100, 20, discount)
    fmt.Println("Final bill after discount:", result)
}*/

