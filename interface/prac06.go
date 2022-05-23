// 다른 인터페이스로 변환하기 
// 타입 변환시 문제점을 해결하는 방법(prac05 참고) 
package main

import (
	"fmt"
)

type Reader interface{
	Read()
}

type Closer interface{
	Close()
}

type File struct{
}

func (f *File) Read(){
	fmt.Println("Read!")
} 

// 만약 Closer가 없다면 
// func (f *File) Close(){
// 	fmt.Println("Close!")
// }

func ReadFile(reader Reader){
	
	// 다음과 같은 방법으로 타입변환을 해도 패닉을 피해갈 수 있음
	c, ok := reader.(Closer) // Closer로 형변환 
	if ok {
		c.Close()  
	}

	// 한줄로 표현하면 
	// if c, ok := reader.(Closer), ok {
	// 	c.Close()  
	// }


}

func main(){
	file := &File{}
	ReadFile(file)
}