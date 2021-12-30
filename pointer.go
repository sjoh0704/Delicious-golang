package main

import (
	"fmt"
)

type Data struct{
	key string
	value string
}

func main(){
	// 1. pointer 기본 
	var p *int // pointer 변수 선언, pointer 변수의 기본값을 nil이다. 
	var a int  // 변수 공간 생성 
	p = &a  // p에 a의 주소를 대입
	*p = 100  // p의 주소에 있는 값을 100으로 변경 

	fmt.Println(p)
	fmt.Println(*p)


	//2. pointer는 같은 값을 가르킬 수 있다.
	b := 10
	var p1, p2 *int
	p1 = &b
	p2 = &b
	*p1 = 20
	fmt.Println(*p2) // 20
	// p1을 변경했음에도 p2가 바뀌는 것은 p1과 p2가 같은 값을 가르키기 때문이다. 


	// 3. 조심해야할점
	data := Data{"k", "v1"}
	fmt.Println(data)
	
	// 잘못된 경우
	setDataWrong(data, "kk", "vv") // 실패
	fmt.Println(data) // 값이 변경되지 않은채로 출력 => data의 주소가 아닌 값을 넣었기 때문 

	// 올바른 경우 
	setDataRight(&data, "kk", "vv")
	fmt.Println(data) // 올바르게 값이 변경되어 출력 
	

	// 4. 인스턴스 생성1
	var p3 *Data = &Data{} // 구조체의 빈 인스턴스 생성 
	var p4 *Data = p3 // 다른 포인터가 같은 인스턴스를 가르키는 경우 
	fmt.Println(*p3)
	(*p4).key = "kk"
	p4.value = "vv"
	fmt.Println(*p3)

	//인스턴스 생성2
	var p5 *Data = new(Data) // new를 사용할때는 무조건 기본값으로 초기화된다. 
	p5.key = "k5"
	p5.value = "v5"
	fmt.Println(*p5)
	// 인스턴스는 아무도 찾지 않을때 사라진다. 힙에 만들어진 경우 다음 가비지콜렉터 타임때 사라진다. 


	// 
	p6 := newData("k6", "v6") // newData가 *Data타입을 반환하므로 p6 또한 자연스럽게 포인터가 된다. 
	fmt.Println(*p6)

}

//객체 생성함수 
func newData(key string, value string) *Data{
	data := Data{key, value} // 이상해보일 수 있다. => 일반적으로 지역변수는 stack에 저장되고 함수가 끝나면 pop되어 사라진다.  
							// Go에서는 탈출 분석을 하여 stack 대신 heap에 저장시키며 쓰임이 다하면 사라진다.  
	return &data
}

// 잘못된 경우
func setDataWrong(data Data, key string, value string){
	data.key = key
	data.value = value
}

// 올바른 경우 
func setDataRight(data *Data, key string, value string){
	(*data).key = key // 이렇게 하는 게 옳지만
	data.value = value // go에서는 편의상 좌측과 같이도 쓴다. 
}