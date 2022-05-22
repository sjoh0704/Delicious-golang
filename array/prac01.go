package main

import (
	"fmt"
)

func main(){

	var a int = 10
	fmt.Println(a)

	var b [3]int = [3]int{1, 2, 3}
	fmt.Println(b)


	var c [6]int = [6]int{2:1, 5:100}
	fmt.Println(c)


	d := [3]int{1, 2, 3} // 배열
	fmt.Println(d)

	e := [...]int{1, 2, 3} // 배열
	fmt.Println(e)


	f := []int{1, 2, 3} // 슬라이스
	fmt.Println(f)
	
}