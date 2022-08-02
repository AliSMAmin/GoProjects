package main

import (
	"fmt"
	_ "rsc.io/quote"
)

var y = 42

// Declares variable 'z' as type 'string'. Writing string is not necessary.
//Since Go is a static language, variables hold values of a certain type.
//Dynamic languages are different
var z string = "Al-manhaj al-nabawi"

type mumin string

var iman bool

func main() {
	var ali mumin
	iman = true
	fmt.Println("Hello world!", y, z, iman)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
	fmt.Printf("%b\n", z)
	fmt.Printf("%x\n", z)
	fmt.Printf("%#x\n%b\n%x\n", z, z, z)
	s, _ := fmt.Printf("%#x\n%b\n%x\n", z, z, z)
	fmt.Println(s)
	fmt.Printf("%T\n", ali)
}
