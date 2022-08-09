package main

import (
	"fmt"
)

func main() {
	fmt.Println("How old is James Bond?")
	m := map[string]int{
		"James":      32,
		"Miss Money": 18,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	var v, ok = m["Charles"]
	fmt.Println(m["Charles"])
	if v, ok = m["Charles"]; ok {
		fmt.Println(m["Charles"])
		fmt.Println(v)
		fmt.Println(ok)
	}
}
