package main

import (
	"fmt"
)

func main() {
	foo()
	bar("James")
	st := woo("Moneypenny")
	fmt.Println(st)
	x, y := mouse("Ian", " Flemming")
	fmt.Println(x)
	fmt.Println(y)

}

// func (r receiver) identifier(parameters) (return(s)) {...}

func foo() {
	fmt.Println("hello from foo")
}

func bar(s string) {
	fmt.Println("Hello,", s)
}

func woo(st string) string {
	return fmt.Sprint("Hello from woo", st)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, ", says 'Hello'")
	b := false
	return a, b

}
