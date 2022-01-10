// custom error 만들기 
package main

import (
	"fmt"
	"math"
)

func sqrt(f float64) (float64, error){
	if f < 0{
		return 0, fmt.Errorf("제곱근은 양수여야 합니다. f:%g", f)
	}

	return math.Sqrt(f), nil

}

func main(){


	res, err := sqrt(-2)
	if err != nil{
		fmt.Printf("%v\n", err) // %v는 객체의 값을 출력시킨다. 
		return 
	}
	fmt.Println(res)
} 