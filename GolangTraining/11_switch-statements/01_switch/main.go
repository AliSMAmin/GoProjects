package main

import "fmt"

func main() {
	fmt.Println("Who are your friends?")
	n := "Daniel"
	switch "Daniel" {
	case n:
		fmt.Println("Aligator is your friend")
		fallthrough
	case "Buddle":
		fmt.Println("Buddle is your friend")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	default:
		fmt.Println("Have you no friends?")
	}
}

/*
  no default fallthrough
  fallthrough is optional
  -- you can specify fallthrough by explicitly stating it
  -- break isn't needed like in other languages
*/
