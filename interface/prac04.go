// 빈 인터페이스

package main

import (
	"fmt"
)

type Student struct{
	Name string
}

// 빈 인터페이스를 사용하는 경우 모든 타입을 의미한다. 
func PrintVal(v interface{}){

	switch t:= v.(type){
	case int:
		fmt.Println("int형입니다.", t)
	case float64:
		fmt.Println("float64형입니다.", t)
	case string:
		fmt.Println("string입니다.", t)
	default:
		fmt.Println("Not supported type", t)
	}
}
// printf의 경우 빈 인터페이스를 받는다. 

func main(){
	PrintVal(12)
	PrintVal(12.3)
	PrintVal("hi")
	PrintVal(Student{"지원"})
}

