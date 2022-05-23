// interface 타입 변환 
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

	// 몬스터에 대한 정보 또한 얻어오고 싶을 때 ! => 타입 변환을 이용 

	// 인터페이스에서 객체의 정보를 가져오는 방법 => 타입변환을 한다. 
	// fmt.Println(att.Lv) => 이런거 안됌 

	// 하지만 이 방식은 좋지 않음 ㅠㅠ => 만약 Monster가 아닌 Player가 공격한다면! Player는 Lv라는 속성을 갖고 있을지 보장 받을 수 없음 
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