package main

import (
	"fmt"
)

const M int = 10


func hash(num int) int{
	return num % M
}

func main(){

	// 해쉬를 이용한 배열 생성

	var nums [M]int
	value := 35
	nums[hash(value)] = value
	fmt.Println(nums)
	 
}