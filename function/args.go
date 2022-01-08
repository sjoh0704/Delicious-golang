package main

import (
	"fmt"
)

func sum(nums ...int) int{
	ret := 0
	for _, v := range nums{
		ret += v
	}
	return ret
}


func main(){
	defer fmt.Println("This is last")
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4, 5, 6))
	fmt.Println(sum(1))
}