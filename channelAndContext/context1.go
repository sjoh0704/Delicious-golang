
package main

import (
	"fmt"
	"time"
	"context"
	"sync"
)

var wg sync.WaitGroup

func main(){

	wg.Add(1)
	// context.Background()는 기본 컨텍스트를 생성한다. 
	// 이후 기본 컨텍스트에 cancel 기능을 가진 context를 덮어쓰기하여 새로운 컨텍스트 생성 
	ctx, cancel := context.WithCancel(context.Background())
	// ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second) // 3초 뒤에 cancel
	// ctx = context.WithValue(ctx, "keyword", "Lilly")
	go PrintEverySecond(ctx)

	time.Sleep(time.Second*5)

	// ctx.Done channel 생성 
	cancel()

	wg.Wait()
}

func PrintEverySecond(ctx context.Context){
	tick := time.Tick(1*time.Second)

	for{
		select{
		case <-tick:
			fmt.Println("tick")
			time.Sleep(time.Second)
		case <-ctx.Done(): // cancel()이 실행되었을 때 ctx.Done()은 채널을 리턴
			fmt.Println("Exit")
			wg.Done()
			return 
		}
	}
}