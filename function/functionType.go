package main

import (
	"fmt"
)

type OpFn func(int, int) int 

func add(a, b int) int{
	return a + b
}

func mul(a, b int) int{
	return a * b
} 

func getOperator(op string) OpFn{
	if op == "+"{
		return add
	}else if op == "*"{
		return mul
	}else{
		return nil
	}
}


func main(){
	var operator func(int, int) int
	operator = getOperator("+")
	fmt.Println(operator(10, 20))

	operator2 := getOperator("*")
	fmt.Println(operator2(10, 20))

	// 별칭 타입 적어주기 
	var operator3 OpFn
	operator3 = getOperator("*")
	fmt.Println(operator3(10, 11))
}