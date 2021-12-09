package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")

	/*
		    creating variable
			var name string = "Joe" is equal to name := "Joe" (:= this symbol recognizes the type of the variable)
			name := "Joe"
			fmt.Println(name)

			name = "Eden"
			age := 22
			fmt.Printf("Hello %v, you are %v!", name, age)
	*/

	fmt.Printf("Enter your name: ")
	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello %v, Welcome to my game!\n", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Yay you can play!")
	} else {
		fmt.Println("You cannot play!")
		return
	}

	score := 0
	num_questions := 2

	fmt.Printf("What is better, the RTX 3080 or RTX 3090? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2) // Scan don't read spaces so to get both rtx and 30380 we use 2 variables

	if answer+" "+answer2 == "RTX 3090" || answer+" "+answer2 == "rtx 3090" {
		fmt.Println("Correct!")
		score++
	} else {
		println("Incorrect!")
	}

	fmt.Printf("How many cores does the Ryzen 9 3090x have? ")
	var cores int
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("You scored %v out of %v.", score, num_questions)

	percent := (float64(score) / float64(num_questions)) * 100

	fmt.Printf("\nYou score: %v%% \n", percent) // use %% inside "" to get % in the output
}
