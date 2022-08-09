package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}
	type secretAgent struct {
		person
		ltk bool
}
func main() {
	sa1 := secretAgent{
		person:	person{
		first: "James",
		last: "Bond",
		age: 32,
	},
		ltk: true,
	}
	p2 := person{
		first: "Miss",
		last: "Money",
		age: 18,
	}
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
	fmt.Println(p2)

}