// 버퍼 관리 예제 
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){

	// 버퍼를 관리하기 위해 사용
	stdin := bufio.NewReader(os.Stdin)

	// scanln
	var str string	
	n2, err := fmt.Scanln(&str)

	if err != nil{
		// 에러가 발생했을 때 \n까지 읽는다(\n까지 입력 버퍼를 비운다)
		fmt.Println("Error:", err)
		stdin.ReadString('\n')
	}else{
		fmt.Println(str)
		fmt.Println(n2)
	}

	fmt.Println("--------------------------------")

	// 위에서 에러가 발생해서 버퍼가 남아있다면, 
	// 다음 코드에서 에러가 발생할 수 있다. 
	// scan
	var a, b int
	n1, _ := fmt.Scan(&a, &b)
	fmt.Println(a, b)
	fmt.Println(n1)

}