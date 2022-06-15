// 채널을 이용한 생산자/소비자 패턴
// 차체 생산 => 바퀴 설치 => 도색(2개의 채널일 필요)
// 채널을 사용한다면 10초
// 채널을 사용하지 않고, 동기적으로 했다면 20초 
package main

import (
	"fmt"
	"sync"
	"time"
)

type Car struct{
	Body string
	Tire string
	Color string
}

var wg sync.WaitGroup
var startTime = time.Now()

func main(){
	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	fmt.Println("Start Factory")

	wg.Add(3) // 3개의 고루틴으로 진행 

	go MakeBody(tireCh)
	go InstallTire(tireCh, paintCh)
	go PaintCar(paintCh)



	wg.Wait()

}

// 차체 생산
func MakeBody(tireCh chan *Car){
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)
	
	for{
		select{
		case <- tick: // I/O
			car := &Car{}
			car.Body = "Sports Car"
			tireCh <- car

		case <- after:
			close(tireCh)
			wg.Done()
			return 
		}
	}
}

// 바퀴 설치 
func InstallTire(tireCh chan *Car, paintCh chan *Car){
	
	for car := range tireCh{
		time.Sleep(time.Second)
		car.Tire = "Nexen tire"
		paintCh <- car
	}
	close(paintCh)
	wg.Done()
}


// 도색 
func PaintCar(paintCh chan *Car){

	for car := range paintCh{
		car.Color = "Red"
		duration := time.Now().Sub(startTime)
		fmt.Println("완성: ", *car, "생산 시간:", duration)
	}
	wg.Done()
}