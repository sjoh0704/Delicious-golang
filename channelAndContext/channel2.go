package main

import (
	"fmt"
	"time"
)


// 일반적으로 메인 고루틴이 더 빨리 끝나면서 서브 고루틴은 실행되기도 전에 끝나고
// Never print가 출력된다.  
func main(){

	//1. 보관함이 없는 경우  
	ch := make(chan int) // 하지만 보관함이 없는 경우 
	go square() 
	ch <- 9  // 수신자가 없으므로 수신할때 까지 대기
	fmt.Println("Never print")


	//2. 버퍼를 가진 채널(보관함이 있는 경우) 
	// ch := make(chan int, 2) // 보관함 2개 생성 
	// go square()
	// ch <- 9 // 채널에 공간이 있으므로 데이터를 보관하고 다음일을 한다.    
	// fmt.Println("print!") 
}

func square(){
	for {
		time.Sleep(time.Second)
		fmt.Println("Sleep")
	}
}