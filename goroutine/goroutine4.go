// 영역을 나누어서 겹치지 않게 하는 방법
// 작업을 나누고 서로에게 터치하지 않는 방식 
package main

import (
	"fmt"
	"sync"
	"time"
)

type Job interface{
	Do()
}

type SquareJob struct{
	index int
}

func (job *SquareJob) Do(){
	fmt.Printf("%d 작업 시작\n", job.index)
	time.Sleep(time.Second)
	fmt.Printf("%d 작업 완료 => 결과: %d\n", job.index, job.index*job.index)
}


func main(){

	var	jobList [10]Job

	for i:=0; i<10; i++{
		jobList[i] = &SquareJob{i}

	}

	var wg sync.WaitGroup
	wg.Add(10)

	for i:=0; i<10; i++{
		j := jobList[i] // 유의 => 새로 변수 할당을 하지 않으면
						// 익명 함수내에서 Pointer로 사용된다. 
		go func(){
			j.Do()
			wg.Done()
		}()
	}

	wg.Wait()

}