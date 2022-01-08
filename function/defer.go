package main

import (
	"fmt"
	"os"
)

func main(){
	f, err := os.Create("test.txt") 
	if err != nil{
		fmt.Println("error 발생")
	} 

	defer f.Close() // os 자원을 돌려준다. 
	fmt.Fprintln(f, "file 쓰기") // 출력 스트림을 파일로 변경 

} 

