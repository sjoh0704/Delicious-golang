// 에러 
// 에러를 사용자단으로 올려서 사용자가 처리할 수 있게 한다. 
package main

import (
	"bufio"
	"fmt"
	"os"
)


func WriteFile(filename string, line string) error {
	file, err := os.Create(filename)
	if err != nil{
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, line)
	return err
}

func ReadFile(filename string)(string, error){
	file, err := os.Open(filename)
	if err != nil{
		return "", err
	}

	defer file.Close()
	rd := bufio.NewReader(file)
	line, _ := rd.ReadString('\n')

	return line, nil
}
func main(){

	filename := "data.txt"
	line, err := ReadFile(filename)
	if err != nil{
		err = WriteFile(filename, "This is the first line")
		if err != nil{
			fmt.Println("Error: Creating file fails")
			return
		}
		return 
	}

	line, err = ReadFile(filename)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println("파일 내용:", line)
}