package main

import "fmt"

const (
	_ = iota
	k = 1 << (iota * 10)
	m = 1 << (iota * 10)
	g = 1 << (iota * 10)
)

func main() {
	var x = 2
	fmt.Printf("%d\t\t%b\n", x, x)
	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)
	kb := 1024
	mb := kb * 1024
	gb := mb * 1024
	//	fmt.Printf("%d - %b - %x \n", 42, 42, 42)
	//	fmt.Printf("%d - %b - %#x \n", 42, 42, 42)
	//	fmt.Printf("%d - %b - %#X \n", 42, 42, 42)
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)
	fmt.Printf("%d\t\t\t%b\n\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n\n", mb, mb)
	fmt.Printf("%d\t\t%b\n\n", gb, gb)
	fmt.Printf("%d\t\t\t%b\n\n", k, k)
	fmt.Printf("%d\t\t\t%b\n\n", m, m)
	fmt.Printf("%d\t\t%b\n\n", g, g)
}
