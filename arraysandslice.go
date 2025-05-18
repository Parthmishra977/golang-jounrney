//Arrays and Slices in Go – With Real-World Examples

//What are Arrays?
//An array is a fixed-size collection of elements of the same type.

//var nums [3]int = [3]int{1, 2, 3}
/*
package main
import "fmt"

func main(){
	var days [3]string = [3]string{"mom","tue","wed"}
	fmt.Println("day 1:",days[0])
	fmt.Println("All Days:", days)
}*/

//What are Slices?
//A slice is like an array, but it can grow or shrink. It's much more commonly used than arrays in Go.

package main
import "fmt"
/*
func main(){
	fruits := []string{"Apple","bannana","mango"}
	   fruits = append(fruits, "Orange") // Add new item
	      fmt.Println("Fruits:", fruits)
}*/


/*
func main() {
    tasks := []string{"Buy milk", "Read Go book"}

    // Add more tasks
    tasks = append(tasks, "Exercise", "Call Mom")

	 fmt.Println("To-Do List:")
    for i, task := range tasks {
        fmt.Printf("%d. %s\n", i+1, task)
	}
}*/

//Slicing

//list := []int{10, 20, 30, 40, 50}
//part := list[1:4] // 20, 30, 40

//Length and Capacity

//fmt.Println(len(list)) // number of elements
//fmt.Println(cap(list)) // capacity of underlying array


/*
func main() {
    results := []string{"Alice", "Bob", "Charlie", "Daisy", "Eve"}

    page1 := results[0:2]
    page2 := results[2:4]

    fmt.Println("Page 1:", page1)
    fmt.Println("Page 2:", page2)
}*/

//Multidimensional Arrays & Slices

//Example: 2D Array (like a matrix)


/*
grid := [2][2]int{
    {1, 2},
    {3, 4},
}
*/

//Tic-Tac-Toe Board

/*

func main() {
    board := [3][3]string{
        {"X", "O", "X"},
        {"O", "X", "O"},
        {"", "X", ""},
    }

    for _, row := range board {
        fmt.Println(row)
    }
}
*/