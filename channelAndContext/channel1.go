package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var wg sync.WaitGroup
	ch := make(chan int)
	
	wg.Add(1)
	go square(&wg, ch)
	ch <- 10
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int){
	var v = <-ch
	time.Sleep(time.Second)
	fmt.Println(v*v)
	wg.Done()
}