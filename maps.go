//Maps in Go – With Real-World Examples

//A map in Go is an unordered collection of key-value pairs. It's like a dictionary in Python or an object in JavaScript.

//long form

//var m map[string]int = map[string]int{"alice":23,"bob":30}

//shorthand
/*
m:= map[string]int{
	"sam":23,
	"james":30,
}*/

package main

import "fmt"
/*
func main(){
	scores := map[string]int{
		"alice":90,
		"bob":85,
	}
	scores["charlie"]=88
	fmt.Println("Scores:", scores)
}*/

//Check if Key Exists (Comma-ok idiom)

/*
value, exists := m["Dan"]
if exists {
    fmt.Println("Found:", value)
} else {
    fmt.Println("Key not found")
}*/

/*
func main() {
    var users = make(map[string]string)
    var count int

    fmt.Print("Enter number of users: ")
    fmt.Scanln(&count)

    for i := 0; i < count; i++ {
        var username, password string
        fmt.Printf("Enter username for user %d: ", i+1)
        fmt.Scanln(&username)
        fmt.Printf("Enter password for %s: ", username)
        fmt.Scanln(&password)
        users[username] = password
    }

    var loginUsername, loginPassword string
    fmt.Print("Login - Enter username: ")
    fmt.Scanln(&loginUsername)
    fmt.Print("Login - Enter password: ")
    fmt.Scanln(&loginPassword)

    if pass, ok := users[loginUsername]; ok && pass == loginPassword {
        fmt.Println("Login successful")
    } else {
        fmt.Println("Invalid credentials")
    }
}
*/
