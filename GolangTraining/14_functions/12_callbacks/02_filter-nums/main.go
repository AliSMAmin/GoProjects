package main
import (
	"fmt"
)
func main(){
	x:=0
	fmt.Println(x)
	x++
	fmt.Println(x)
	x+=42
		fmt.Println(x)
	x-=3
	fmt.Println(x)
	ii:= []int{1,2,3,4,5,6,7,8,9}
	s:=sum(ii...)
	fmt.Println(s)
}

func sum(xi ...int) int{
	fmt.Printf("%T\n",xi)
	total :=0
	for _, v:=range xi {
		total +=v
	}
	return total
}

func even (f func(xi ...int) int, vi ...int) int{
	var yi []int
	for _, v:=range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	f(yi ...)
}
}
}