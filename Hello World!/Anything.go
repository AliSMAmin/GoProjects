package main

import "fmt"

func main() {
	var (
		x int = 33
		y int = x + 1
	)
	fmt.Println("Allahu Akbar!")
	Tasbih()
	Tahlil()
	fmt.Println("Alhamdulillah", y)
	for i := 0; i < 100; i++ {
		if i%2 == 0 {

		}
		fmt.Println(i)
	}
}
func Tasbih() {
	fmt.Println("SubhanAllah")
}

func Tahlil() {
	fmt.Println("La ilaha ila Allah")
}

//control flow: (1) sequence
//(2) loop or iterative
//(3) conditional
