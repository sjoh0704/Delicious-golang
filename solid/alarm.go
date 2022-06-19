// 의존관계 역전 원칙2 예제
// 메일이라는 구체화된 모듈과 알람이라는 구체화된 모듈을 분리하기 
package main

import (
	"fmt"
)

type Mail struct{
	listener EventListener
}

func (m *Mail) Register(listener EventListener){
	m.listener = listener
	fmt.Println("Mail 이벤트 리스너가 등록되었습니다.")
}

func (m *Mail) OnRecv(){
	m.listener.OnFire()
}

type Event interface{
	Register(EventListener)	
}

type EventListener interface{
	OnFire()
}

type Alarm struct{
}

func (a *Alarm) OnFire(){
	fmt.Println("메일이 왔습니다.")
}


func main(){

	var listener EventListener = &Alarm{}
	var event Event = &Mail{}

	event.Register(listener)
	listener.OnFire()
}