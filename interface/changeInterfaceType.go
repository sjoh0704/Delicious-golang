// 다른 인터페이스로 변환하기 

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
	// reader에 close() 함수가 없어도 빌드시에 에러가 나지 않는다. 
	// 하지만 close()를 사용하면 error 발생  
	// 그래서 두번째 인자를 통해서 성공 여부를 확인할 수 있다. 
	
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