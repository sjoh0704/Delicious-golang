package main

import (
	"fmt"
)

func main(){

	// 1. " "와 ` `비교하기 
	s1 := "hi\tI am seung\nhow about you"

	// ``의 경우에는 \n, \t가 반영되지 않고, 그대로 출력된다. 
	s2 := `hi	I am seung\n
how about you`
	fmt.Println(s1)
	fmt.Println(s2)

	//2. len은 바이트 길이를 나타낸다. 문자의 갯수가 아니다!

	s3 := "hello 월드"
	for i:=0; i<len(s3); i++{  
		fmt.Printf("%d번째 값: %T, %d, %c \n", i+1, s3[i], s3[i], s3[i]) // 8개의 값이 아닌 12개의 값이 출력된다. 
	}
	//"월"과 "드"가 3바이트의 크기를 갖고 있기 때문에 1바이트식 출력하면 깨진다. 

	//3. len대신 rune 사용하기 => 문자 순회를 할때 쓴다.  
	s4 := []rune(s3)  // 슬라이스 rune 타입(int32 별칭)으로 만들기 
 	for i:=0; i<len(s4); i++{
		fmt.Printf("%d번째 값 %T, %d, %c\n", i+1, s4[i], s4[i], s4[i])  // 8개의 값이 차례대로 나온다. 
	}
	// int32와 제대로된 글자가 출력된다.(한 글자당 4바이트를 차지하면서 한 글자당 한 자료형에 들어가게 된다) 

	//4. int는 8byte이다. 그렇다면 string은 몇 바이트일까?
	// 다음 str1과 str2는 타입이 같지만 크기가 다른데 어떻게 대입이 가능할까? 
	var str1 string = "hello" 
	var str2 string
	str2 = str1
	fmt.Println(str2)
	
}