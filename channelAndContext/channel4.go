// select문 사용하기(close 없이) 
package main

import (
	"fmt"
	"time"
	"sync"
)

func main(){

	var wg sync.WaitGroup

	ch := make(chan int)
	wg.Add(1)

	go square(&wg, ch)
	for i:=0; i<10; i++{
		ch <- i
		time.Sleep(time.Second)
	}
	// close(ch) // close 없이  
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int){
	tick := time.Tick(time.Second) // 일정 간격으로 신호를 주는 채널을 리턴 
	terminate := time.After(10*time.Second) // 일정 시간 대기 후 신호를 주는 채널을 리턴
	for{
		// 게임에서 이런 방식을 많이 쓴다. (애니메이션 + 입력)
		select{
		case <- tick:
			// 애니메이션 
			fmt.Println("Tick")
		case <- terminate:
			fmt.Println("Terminate")
			wg.Done()
			return 
		case n := <- ch:
			// 입력 
			fmt.Println(n*n)
		}
	}
}