package main

import "fmt"

func main() {
	x:=2; if x==2 {
		fmt.Println("01")
	}
	if x!=2 {
		fmt.Println("02")
	}
	if x==2 {
		fmt.Println("03")
	}
	if !(x!=2) {
		fmt.Println("04"); fmt.Println("05");
	}
}
