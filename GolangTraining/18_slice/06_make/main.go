package main

import "fmt"

func main() {
	x:= make ([]int, 10,12)
	x[0] = 42
	x[9]=999
	x = append(x, 3423)
	x = append(x, 3424)
	x = append(x, 3425)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
