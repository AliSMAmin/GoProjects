package main

import "fmt"

func main() {
	x:=[]int{4, 5, 6, 8, 42}
	fmt.Println(len(x));
	x=append(x, 77, 88, 99, 1014)
	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])

	for i, v:=range x {
	fmt.Println(i, v)}

	y:=[]int{234, 456, 678, 987}
	x=append(x, y...)
	}