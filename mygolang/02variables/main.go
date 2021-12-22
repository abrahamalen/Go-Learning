package main

import "fmt"

const LoginToken = "abcdefg" // first letter capital in golang implies it is a public variable, similar to public cont loginToken

func main() {
	var username string = "bond"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallFloat float64 = 255.4545457866
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type

	var website = "udemy.com"
	fmt.Println(website)

	// no var style

	numberOfUser := 300000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
