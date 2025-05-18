//What is a Struct?
//A struct is a custom data type that groups related data fields. Think of it like a record or a class without methods.
/*
type Person struct {
    Name string
    Age  int
}
*/
//people := map[string]Person{}
/*
package main

import (
    "fmt"
)

type Employee struct {
    Age  int
    Role string
}

func main() {
    employees := make(map[string]Employee)
    var count int

    fmt.Print("Enter number of employees: ")
    fmt.Scanln(&count)

    for i := 0; i < count; i++ {
        var name, role string
        var age int

        fmt.Printf("Enter name of employee %d: ", i+1)
        fmt.Scanln(&name)
        fmt.Printf("Enter age of %s: ", name)
        fmt.Scanln(&age)
        fmt.Printf("Enter role of %s: ", name)
        fmt.Scanln(&role)

        employees[name] = Employee{Age: age, Role: role}
    }

    // Search for an employee
    var searchName string
    fmt.Print("Enter the name of the employee to search: ")
    fmt.Scanln(&searchName)

    if emp, ok := employees[searchName]; ok {
        fmt.Printf("Employee Found: %s (Age: %d, Role: %s)\n", searchName, emp.Age, emp.Role)
    } else {
        fmt.Println("Employee not found.")
    }
}
*/