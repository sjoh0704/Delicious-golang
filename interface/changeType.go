package main

import (
	"fmt"
)

type Monster struct{
	Lv int
}

func (m *Monster) Attack(){
	fmt.Printf("Lv: %d Monster Attack!\n", m.Lv)
}

type Attacker interface{
	Attack()
}

func DoAttack(att Attacker){

	if att == nil{
		fmt.Println("indexing error!")
		return
	}
	att.Attack()

	// 인터페이스에서 객체의 정보를 가져오는 방법 => 타입변환을 한다. 
	// fmt.Println(att.Lv) => 이런거 안됌 

	monster := att.(*Monster)
	if monster.Lv > 10{
		fmt.Println("강한 공격!!")
	}else{
		fmt.Println("약한 공격...")
	}
}

func main(){
	// Error가 발생한다. nil pointer error
	// => panic: runtime error: invalid memory address or nil pointer dereference
	var at1 Attacker
	DoAttack(at1)

	var at2 Attacker = &Monster{20}
	DoAttack(at2)

	var at3 Attacker = &Monster{10}
	DoAttack(at3)
	

}