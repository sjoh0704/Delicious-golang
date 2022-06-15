// 채널 닫아주기 
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
	}

	close(ch) // 좀비고루틴이 되는 것을 막기 위해 채널을 닫는다. 
			  // 채널에 보낼 것들을 다 보냈다면 채널을 닫아주자. 

	wg.Wait()


}

func square(wg *sync.WaitGroup, ch chan int){

	for i:= range ch{ // 채널로부터 무한히 받는다.
		fmt.Println(i*i)
		time.Sleep(time.Second)		
	}
	// 위에서 채널로부터 무한히 받으므로 Done을 할 수가 없게 되고, 좀비 고루틴이 된다.
	
	wg.Done()
}