package main

import (
	"fmt"
	_ "fmt"
	"sort"
	_ "sort"
)

type person struct {
	first string
	age   int
}

type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }

func main() {
	p1 := person{"James", 32}
	p2 := person{"Money", 19}

	people := []person{p1, p2}

	sort.Sort(ByAge(people))
	fmt.Println(people)

}
