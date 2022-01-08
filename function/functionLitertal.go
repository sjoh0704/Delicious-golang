package main

import (
	"fmt"
)

func CaptureLoop(){
	loop := make([]func(), 3)
	for i:=0; i<3; i++{
		loop[i] = func(){
			fmt.Println(i) // 이 경우 loop 0, 1, 2는 모두 같은 주소를 참조하므로 
						   // 모두 동일한 가르키고 결국 3이 출력된다. 
		}
	}
	for i:=0; i<3; i++{
		loop[i]() // 3 3 3 출력
	}
}

func CaptureLoop2(){

	loop := make([]func(), 3)

	for i:=0; i<3; i++{	
		v := i  // 매 루프마다 새로운 변수를 할당해줌으로 
		loop[i] = func(){
			fmt.Println(v)  // 모두 다른 값이 출력되는 것. 
		}
	}

	for i:=0; i<3; i++{
		loop[i]() // 0 1 2 출력  
	}
}

func main(){
	// 기본적인 함수 리터럴 
	add := func(a int, b int) int{
		return a + b
	}

	fmt.Println(add(10, 20))

	i := 0 
	f := func(){
		i++	// 외부의 변수를 내부에서 상태처럼 사용할 수 있다.
			// 같은 메모리 공간을 참조한다.  
	}
	f()
	fmt.Println(i)

	// 익명 함수에서 변수 참조를 어떻게 하냐의 차이 
	CaptureLoop()
	CaptureLoop2()

}