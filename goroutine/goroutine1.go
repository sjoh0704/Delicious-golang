// 고루틴 기본!

package main

import (
	"fmt" 
	"time"
)

// 고루틴으로 실행시킬 함수 1
func PrintHan(){
	han := []rune{'가', '나', '다', '라', '마', '바'}
	for _, v:= range han{
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", v)
	}
}

// 고루틴으로 실행시킬 함수 2
func PrintNum(){
	for i:= 1; i<6; i++{
		time.Sleep(350 * time.Millisecond)
		fmt.Printf("%d ", i)
	}

}

// main goroutine에서 2개의 go routine을 실행시킨다. 
func main(){
	go PrintHan()
	go PrintNum()

	// 메인 고루틴이 먼저 종료되면 내부의 다른 고루틴 또한 먼저 종료되므로 
	// 내부 고루틴이 종료될 때 까지 기다려 준다.  
	time.Sleep(3 * time.Second)

}

// 결과적으로 3개의 고루틴이 실행되며 코어의 수에 따라 컨텍스트 스위칭의 유무가 달라진다.
// 현재 논리 코어가 3개이상이므로 컨텍스트 스위칭이 일어나지 않는다. 

