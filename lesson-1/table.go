package main

import (
	"fmt"
)

const (
	id        string = "ID"
	name      string = "NAME"
	cost      string = "COST($)"
	separator string = "----------"
)

func main() {
	// row 1
	cl1 := "1"
	cl2 := "Car"
	var cl3 float32 = 350.15

	// row 2
	cl4 := "2"
	cl5 := "Robot"
	var cl6 float32 = 12.5456

	// row 3

	cl7 := "3"
	cl8 := "Phone"
	var cl9 float64 = 600.000

	// separator

	fmt.Printf("|%+10s|%+10s|%+10s|\n", id, name, cost)
	fmt.Printf("|%-10s|%-10s|%-10s|\n", separator, separator, separator)
	fmt.Printf("|%-10s|%-10s|%-10.2f|\n", cl1, cl2, cl3)
	fmt.Printf("|%-10s|%-10s|%-10.2f|\n", cl4, cl5, cl6)
	fmt.Printf("|%-10s|%-10s|%-10.2f|\n", cl7, cl8, cl9)

}
