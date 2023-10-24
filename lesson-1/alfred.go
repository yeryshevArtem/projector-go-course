package main

import (
	"fmt"
)

func main() {
	var name string
	var age int
	var jsQuestion string
	var goQuestion string

	fmt.Printf("Hello! What is your name: ")

	fmt.Scanf("%s", &name)

	fmt.Printf("Nice to meet you %s\n", name)

	fmt.Printf("How old are you: ")

	fmt.Scan(&age)

	fmt.Printf("Wooow! You are so young. You are just %d years old\n", age)

	fmt.Print("Do you like JavaScript: ")

	fmt.Scan(&jsQuestion)

	fmt.Printf("So your answer is %s\n", jsQuestion)

	fmt.Printf(
		`
		I guess your answer was 'Yes'! Good choice :) \n
		Let's double check... Your answer was %s
		But it's nice to have one more programming language in your skillset? Do you want to learn Go: `, jsQuestion)

	fmt.Scan(&goQuestion)

	fmt.Printf("And again, I'm pretty sure that you answer was 'Yes', but let's double check. So your answer was %s\n", goQuestion)

	fmt.Println("Thank you", name, "for your answers.", "We will see on the next lesson.")
}
