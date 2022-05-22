package main

import (
	"fmt"
)

type Stringer interface{
	String() string
}

type Student struct{
	Name string
	Age int
}

func (s Student) String() string{
	return fmt.Sprintf("안녕 나는 %d살 %s라고 해", s.Age, s.Name)
}

func (s Student) String2() string{
	return fmt.Sprintf("안녕 나는 %d살 %s라고 해", s.Age, s.Name)
}


func main(){

	stu := Student{"제이", 12}
	
	var stringer Stringer

	stringer = stu

	fmt.Println(stringer.String())

	// fmt.Println(stringer.String2()) // 이거는 안된다. 

}