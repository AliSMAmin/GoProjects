package main

import "fmt"

func main() {
	z := 1
	// Short declaration operator, z and 1 are operands to declare a value
	//Variable declaration and value assignation simultaneously
	var (
		x int = 33 + z
		y int = x - 1
	)
	fmt.Println("Allahu Akbar!", "x", y)
	Tasbih()
	Tahlil()
	fmt.Println("Alhamdulillah", "x", y)
	for i := 0; i < 100; i++ {
		if i%2 == 0 {

		}
		fmt.Println(i)
	}
}
func Tasbih() {
	fmt.Println("SubhanAllah", "x", "33")
}

func Tahlil() {
	fmt.Println("La ilaha ila Allah")
}

//control flow: (1) sequence
//(2) loop or iterative
//(3) conditional
