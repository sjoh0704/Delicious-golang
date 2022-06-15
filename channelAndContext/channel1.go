package main

import (
	"fmt"
	"sync"
	"time"
)

// main 고루틴과 square 고루틴이 채널을 통해서 값을 주고 받는다. 
// func main(){
// 	var wg sync.WaitGroup
// 	ch := make(chan int)
	
// 	wg.Add(1)
// 	go square(&wg, ch)
// 	ch <- 10
// 	wg.Wait()
// }

// func square(wg *sync.WaitGroup, ch chan int){
// 	var v = <-ch
// 	time.Sleep(time.Second)
// 	fmt.Println(v*v)
// 	wg.Done()
// }

func main(){

	num := 3
	ch := make(chan int, 3) 
	var wg sync.WaitGroup
	wg.Add(num)
	for i:=0; i<num; i++{
		idx := i
		// ch <- idx // 채널의 크기가 0인 경우 수신자가 대기하고 있어야 하기 때문에 에러가 남
		go func (){
			Test(ch)
			wg.Done()
		}()
		ch <- idx
	}
	wg.Wait()

}

func Test(ch chan int){
	idx := <- ch
	fmt.Println("Test 시작: ", idx)
	time.Sleep(time.Second)
	fmt.Println("Test 끝: ", idx)
}