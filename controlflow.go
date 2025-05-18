//Control flow helps your program make decisions and repeat tasks.

//if elseif else
//syntax

/*if condition {
    // do something
} else if anotherCondition {
    // do something else
} else {
    // default case
}*/

package main
import "fmt"

/*func main() {
  var  age int
  fmt.Print("Enter some text: ")
  fmt.Scanln(&age)

    if age >= 18 {
        fmt.Println("You are an adult.")
    } else {
        fmt.Println("You are a minor.")
    }
}*/
/*func main(){
	var loginAttempts int
	fmt.Println("enter no of login attempts")
	fmt.Scan(&loginAttempts )
	if loginAttempts == 0 {
        fmt.Println("First attempt.")
    } else if loginAttempts < 3 {
        fmt.Println("Careful! You have tried", loginAttempts, "times.")
    } else {
        fmt.Println("Account locked. Too many attempts.")
    }
}*/

//switch
//Used for checking multiple conditions more cleanly than many if-else blocks.
//syntax
/*switch variable {
case value1:
    // action
case value2:
    // action
default:
    // default action
}*/


/*func main(){
	var day string
	fmt.Println("enter the day")
	fmt.Scanln(&day)
	switch day {
	case "Monday":
		fmt.Println("Start of the week.")
	case "Friday":
        fmt.Println("Weekend is coming.")
	default:
        fmt.Println("Just another day.")
	}

}*/




/*func main() {
    status := "shipped"

    switch status {
    case "pending":
        fmt.Println("Your order is being processed.")
    case "shipped":
        fmt.Println("Your order is on the way.")
    case "delivered":
        fmt.Println("Your order has been delivered.")
    default:
        fmt.Println("Unknown status.")
    }
}*/

//for loop
//Go’s only loop keyword is for, but it supports:

//Traditional loops

//Range loops

//Infinite loops


//Traditional for

/*func main() {
    for i := 1; i <= 5; i++ {
        fmt.Println("Step:", i)
    }
}*/


//Loop over slice (using range)

/*func main(){
	items := []string{"book","pen","notebook"}

	for index,item := range items{
		fmt.Printf("Item %d: %s\n", index, item)
	}
}*/


/*func main() {
    users := []string{"Alice", "Bob", "Charlie"}

    for _, user := range users {
        fmt.Println("Welcome,", user)
    }
}
*/

/*func main() {
    count := 0

    for {
        count++
        fmt.Println("Count:", count)

        if count == 3 {
            break // exit the loop
        }
    }
}*/


