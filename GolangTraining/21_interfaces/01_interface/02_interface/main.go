package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secretagent struct {
	person
	ltk bool
}

// func (r receiver) identifier(parameters) (return(s)) {code}

func (s secretagent) speak() {
	fmt.Println("I am", s.first, s.last)
}

type human interface {
	speak()
}

func main() {
	sa1 := secretagent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}
	sa2 := secretagent{
		person: person{
			"Miss",
			"Moneypenny",
		},
	}
	p1 := person{
		first: "Dr.",
		last:  "No",
	}
	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
	fmt.Println(p1)
}
