package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Bonjour!",
	}
	delete(myGreeting, "Tim")
	myGreeting["Harleen"] = "Howdy"
	fmt.Println(myGreeting)
	myGreeting["Harleen"] = "Gidday"
	fmt.Println(myGreeting)
}
