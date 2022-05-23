// 인터페이스 사이즈

package main

import (
	"fmt"
	"unsafe"
)

type Student struct{
	Name string
}

func (s *Student) String() string {
	return fmt.Sprintf("%s", s.Name)
}


type User struct{
	Name string
}

func (u User) String() string {
	return fmt.Sprintf("%s", u.Name)
}


type Stringer interface{
	String() string
}


func main(){

	s1 := &Student{"sseung"}
	var i1 Stringer

	fmt.Printf("%T, %d\n", i1, unsafe.Sizeof(i1))
	i1 = s1


	fmt.Printf("%T, %d\n", i1, unsafe.Sizeof(i1))


	u1 := User{"ppeng"}
	var i2 Stringer


	fmt.Printf("%T, %d\n", i2, unsafe.Sizeof(i2))
	i2 = u1

	fmt.Printf("%T, %d\n", i2, unsafe.Sizeof(i2))

}

// 인터페이스는 인터페이스 메모리 주소 + 타입정보를 저장한다. 