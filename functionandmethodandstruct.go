/*func greet(name string) {
    fmt.Println("Hello", name)
}

type Person struct {
    Name string
}

func (p Person) greet() {
    fmt.Println("Hello", p.Name)
}
*/

//Student Report Card System

package main

import (
    "fmt"
)

type Student struct {
    Name  string
    Marks []float64
}


// Method: calculates average marks

/*
func (s Student) Average() float64 {
    total := 0.0
    for _, mark := range s.Marks {
        total += mark
    }
    return total / float64(len(s.Marks))
}

func main() {
    var students []Student
    var numStudents int

    fmt.Print("Enter number of students: ")
    fmt.Scanln(&numStudents)

    for i := 0; i < numStudents; i++ {
        var name string
        var numSubjects int

        fmt.Printf("Enter name of student %d: ", i+1)
        fmt.Scanln(&name)

        fmt.Printf("Enter number of subjects for %s: ", name)
        fmt.Scanln(&numSubjects)

        marks := make([]float64, numSubjects)
        for j := 0; j < numSubjects; j++ {
            fmt.Printf("Enter mark for subject %d: ", j+1)
            fmt.Scanln(&marks[j])
        }

        students = append(students, Student{Name: name, Marks: marks})
    }

    fmt.Println("\nStudent Report Card:")
    for _, student := range students {
        fmt.Printf("Name: %s | Average Score: %.2f\n", student.Name, student.Average())
    }
}
*/
