// 여러 고루틴이 한 값에 접근하는 경우 
// 뮤텍스를 이용해서 해결하기 
package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct{
	Balance int
}

func depositAndWithdraw(ac *Account){

	// mutex를 이용해서 lock을 건다. 
	mutex.Lock()
	defer mutex.Unlock()

	if ac.Balance < 0{
		panic(fmt.Sprintf("negative value is not possible: %d", ac.Balance))
	}
	ac.Balance -= 1000
	time.Sleep(time.Millisecond)
	ac.Balance += 1000
}

var mutex sync.Mutex

func main(){

	var wg sync.WaitGroup
	wg.Add(10)

	// 힙 영역에 객체 하나 생성
	myAccount := &Account{10}

	for i:=0; i<10; i++{
		go func(){
			for{
				depositAndWithdraw(myAccount)
			}
			wg.Done()
		}()
	}

	wg.Wait()

}