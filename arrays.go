//What is an Array?
//An array is a fixed-size collection of elements of the same type.
// syntax = var arr [5]int//array of 5 integers 
//sample code
package main

import "fmt"

/*func main() {
    var numbers [3]int
    numbers[0] = 10
    numbers[1] = 20
    numbers[2] = 30

    fmt.Println("Array values:", numbers)
}*/

//Store Weekly Sales
/*func main() {
    var weeklySales [7]float64 = [7]float64{123.45, 140.90, 99.99, 120.00, 135.00, 0.0, 0.0}

    fmt.Println("Monday to Sunday Sales:")
    for i, sale := range weeklySales {
        fmt.Printf("Day %d: $%.2f\n", i+1, sale)
    }
}*/

//What is a Slice?
//A slice is like an array, but resizable.

//It’s the most used collection type in Go.
//syntax
/*func main(){
nums:=[]int{10,20,30}
nums = append(nums,40)
}*/

/*func main(){
	scores := []int{90,85,88}
	scores = append(scores,92)
	fmt.Println("Scores:", scores)
}*/



/*func main() {
    cart := []string{"Laptop", "Mouse"}
    cart = append(cart, "Keyboard", "Monitor")

    fmt.Println("Items in Cart:")
    for _, item := range cart {
        fmt.Println("-", item)
    }
}*/

//What is a Map?
//A map is a collection of key-value pairs. It’s like a dictionary or hash table.

//m := make(map[string]int)
//m["apple"] = 10

/*func main() {
    ages := map[string]int{
        "Alice": 30,
        "Bob":   25,
    }

    ages["Carol"] = 28

    fmt.Println("Ages:", ages)
}*/



/*func main() {
    grades := map[string]string{
        "Alice": "A",
        "Bob":   "B+",
        "John":  "A-",
    }

    grades["Eve"] = "B"

    fmt.Println("Student Grades:")
    for student, grade := range grades {
        fmt.Printf("%s: %s\n", student, grade)
    }
}
*/


