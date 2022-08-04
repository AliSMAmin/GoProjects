package main

import "fmt"

func main() {
	a := (42 == 43)
	b := (42 <= 43)
	c := (42 >= 43)
	d := (42 != 43)
	e := (42 < 43)
	f := (42 > 43)
	//Loops are for init; condition; post {}
	fmt.Println(a, b, c, d, e, f)
	for i := 0; i <= 100; i++ {
		fmt.Printf("The outer loop: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t \t The Inner Loop %d\n", j)
		}
	}
}
