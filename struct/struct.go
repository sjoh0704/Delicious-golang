// Go는 클래스가 없고, 구조체만 지원한다. 
// 구조체를 이용하면 의존성을 낮추고, 응집성은 높일 수 있다. 

package main

import (
	"fmt"
	"unsafe" // 안전하지 않은 내장 모듈
)

type User struct {
	Name string
	Gender string
	Age int32
	Height float32
}

type Student struct{
	UserInfo User
	Grade int32
	GPA float32
}

type MemoryTest struct{
	A int8 // 1
	C int  // 8
	B int8 // 1
	D int //  8
	// 총 8 x 4 = 32 바이트의 크기를 소모
}

// 메모리 공간 아껴쓰기 
type MemoryTest2 struct{
	A int8 // 1
	C int16  // 2
	B int8 // 1
	D int //  8
	// 총 8 + 8 = 16바이트가 메모리를 차지한다.
	// 초기 1, 2, 1 바이트의 변수들이 모두 8바이트 안에 들어가므로 총 16바이트의 공간을 사용한다.  
	// 즉 팁이라고 한다면 8바이트보다 작은 필드는 앞으로 밀어서 배치한다. 
}

func main(){

	// User struct의 오브젝트 생성하기 
	var u1 User
	u1.Name = "제임스"
	u1.Gender = "남"
	u1.Age = 20
	u1.Height = 178.3
	fmt.Println(u1)

	// 임베디드 struct의 object 생성하기 
	s1 := Student{
		User{"승주", "남", 26, 180},
		3,
		4.3, // 마지막 콤마도 필요(규칙성)
	}
	fmt.Println(s1)

	// struct의 메모리 공간 

	fmt.Println(unsafe.Sizeof(u1)) // u1이 차지하는 메모리 공간
	// 메모리의 성능을 위해서 변수의 크기 + padding을 이용해서 배치시킨다. 
	// padding 때문에 1바이트 짜리도 8바이트의 공간을 차지하게 된다. 
	// 메모리가 낭비된다. 

	
	m1 := MemoryTest{1, 2, 3, 4}
	fmt.Println(unsafe.Sizeof(m1)) // 32 바이트
	
	m2 := MemoryTest2{1, 2, 3, 4}
	fmt.Println(unsafe.Sizeof(m2)) // 16 바이트
	
}