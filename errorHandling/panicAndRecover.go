// panic은 살리는 방법이 아닌 죽이는 기능 
// recover는 panic을 복구하는 기능 
package main

import (
	"fmt"
)

func f(){
	fmt.Println("f()함수 시작!")
	
	// panic을 복구
	defer func(){
		if r:=recover(); r!=nil{
			fmt.Println("Recover: panic 복구-", r)
		}
	}()

	g()
	fmt.Println("f()함수 끝!")

}

func g(){
	h(9, 3)
	h(9, 0)

}


func h(a int, b int){
	if b == 0{
		panic("두번째 인자가 0일 수 없습니다.") // 이 시점에서 프로그램이 죽는다.  
	}
	fmt.Printf("Result: %d\n", a/b)
}

func main(){
	f()
	fmt.Println("다음 프로그램이 실행됩니다. ")
} 