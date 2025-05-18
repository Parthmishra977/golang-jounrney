/*package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
package main: Defines the package name. main is the entry point.

import "fmt": Imports a package for formatted I/O.

func main(): The main function where the program starts.*/

package main

import "fmt"

func main() {
    var name string = "Alice"
    var age int = 25
    var height float64 = 5.6
    isStudent := true // shorthand

    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Height:", height)
    fmt.Println("Student?", isStudent)
}
