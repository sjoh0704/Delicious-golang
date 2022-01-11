// 서브 고루틴이 종료될 때까지 기다리기 

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func SumAtoB(a, b int){
	sum:= 0
	for i:=a; i < b+1; i++{
		sum += i
	}

	fmt.Println(sum)

	wg.Done() // 작업이 끝났다는 것을 알려주는 것
}

func main(){

	wg.Add(10) // 10개의 작업을 넣기 

	for i:=0; i<10; i++{ // 10개의 작업을 고루틴으로 만든다. 
		go SumAtoB(1, 100000000) 
	}
	
	wg.Wait() // 10개의 작업이 끝날 때까지 기다린다.
}

      