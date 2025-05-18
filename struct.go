//What is a Struct?
//A struct is a collection of fields. It lets you define and group related data together.

//Think of it like a blueprint for a real-world entity.

// basic syntax

/*type person struct{
	name string
	age int
}*/
/*
package main
import "fmt"

type person struct{
	name string
	age int
}

func main() {
    person1 := person{name: "John", age: 25}
	person2 := person{name: "Bob", age: 25}			
    fmt.Println("Name:", person1.name)
    fmt.Println("Age:", person1.age)
	fmt.Println("Name:", person2.name)
    fmt.Println("Age:", person2.age)
}
	*/

/*	package main

import "fmt"

type User struct {
    Username string
    Email    string
    IsActive bool
}

func main() {
    u := User{
        Username: "coder123",
        Email:    "coder@example.com",
        IsActive: true,
    }

    fmt.Printf("User: %s (%s)\n", u.Username, u.Email)
}
*/

//You can define a struct inline without naming it:
/*employee := struct {
    ID   int
    Name string
}{ID: 1, Name: "Alice"}
*/

//Structs with Functions (Methods)
/*
type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
*/
/*
package main

import "fmt"

type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    rect := Rectangle{Width: 10.0, Height: 5.0}
    fmt.Println("Area:", rect.Area())
}
*/

//Structs with Pointers
/*
package main

import "fmt"

type User struct {
    Username string
    IsActive bool
}

func (u *User) Deactivate() {
    u.IsActive = false
}

func main() {
    u := User{Username: "alice", IsActive: true}
    fmt.Println("Before:", u.IsActive)
    u.Deactivate()
    fmt.Println("After:", u.IsActive)
}
*/