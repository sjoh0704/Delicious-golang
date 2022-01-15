// 생성자, 소비자 패턴
// 차체 생산 => 바퀴 설치 => 도색
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
		case <- tick:
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
		fmt.Println(car, duration)
	}
	wg.Done()
}