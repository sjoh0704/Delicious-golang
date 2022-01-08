package main

import (
	"fmt"
	"os"
)

type Writer func(string)

func writeHello(writer Writer){
	writer("hello world")
}

// 인터페이스를 이용하기
type writerInterface interface{
	writeHello(Writer)
}

func writeHello2(writer writerInterface){
	writer.writerHello("hello world2")
}


func main(){

	file, err := os.Create("test.txt")
	if err != nil{
		fmt.Println("error")
		return 
	}
	
	defer file.Close()


	// 외부에서 로직을 주입한다. 
	writeHello(func(msg string){
		fmt.Fprintln(file, msg)
	})

	writeHello(func(msg string){
		fmt.Println(msg)
	})

}