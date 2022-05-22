package main

import (
	"fmt"
	"os"
)

func main(){
	f, err := os.Create("test.txt")
	defer f.Close()

	if err != nil{
		fmt.Println("error")
	}

	fmt.Fprintln(f, "쓰기")


}