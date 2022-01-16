// 인터페이스 분리 
// 생각보다 복잡하고 어렵다. 
package main

import (
	"fmt"
)

type Mail struct{
	listener EventListener
}

func (m *Mail) Register(listener EventListener){
	m.listener = listener
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

	mail := &Mail{}
	mail.Register(listener)

	mail.OnRecv() 
}