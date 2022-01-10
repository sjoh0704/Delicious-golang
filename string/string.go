package main

import (
	"fmt"
	"unsafe"
	"reflect"
)

func main(){

	// 1. " "와 ` `비교하기 
	s1 := "hi\tI am seung\nhow about you"

	// ``의 경우에는 \n, \t가 반영되지 않고, 그대로 출력된다. 
	s2 := `hi	I am seung\n
how about you`
	fmt.Println(s1)
	fmt.Println(s2)

	// 2. len은 바이트 길이를 나타낸다. 문자의 갯수가 아니다!

	s3 := "hello 월드"
	for i:=0; i<len(s3); i++{  
		fmt.Printf("%d번째 값: %T, %d, %c \n", i+1, s3[i], s3[i], s3[i]) // 8개의 값이 아닌 12개의 값이 출력된다. 
	}
	//"월"과 "드"가 3바이트의 크기를 갖고 있기 때문에 1바이트식 출력하면 깨진다. 

	// 3. len대신 rune 사용하기 => 문자 순회를 할때 쓴다.  
	s4 := []rune(s3)  // 슬라이스 rune 타입(int32 별칭)으로 만들기 
 	for i:=0; i<len(s4); i++{
		fmt.Printf("%d번째 값 %T, %d, %c\n", i+1, s4[i], s4[i], s4[i])  // 8개의 값이 차례대로 나온다. 
	}
	// int32와 제대로된 글자가 출력된다.(한 글자당 4바이트를 차지하면서 한 글자당 한 자료형에 들어가게 된다) 

	// 4. int는 8byte이다. 그렇다면 string은 몇 바이트일까?
	// 다음 str1과 str2는 타입이 같지만 크기가 다른데 어떻게 대입이 가능할까? 
	var str1 string = "hello"    // Data 필드에 주소, Len 필드에 문자열의 길이를 저장한다. ex) 100/6 
	var str2 string
	str2 = str1
	fmt.Println(str2)
	
	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))
	fmt.Println(stringHeader1) // &{4932410 5}
	fmt.Println(stringHeader2) // &{4932410 5}

	// 5.  string 변환하기. string은 불변이다!
	
	var s5 string = "hello world"
	var s6 []byte = []byte(s5)

	// s5[2] = "a" // Error
	s6[2] = 'a' // 정상

	fmt.Println(s5)
	fmt.Printf("%s\n", s6)

	// 문자열이 불변이면 좋은점 => 중간에 문자열을 변경할 수 없으므로 여러개의 인스턴스가 문자열을 공유해도 된다.
	s7 := "hello!!"
	s8 := s7 // s8이라는 공간을 생성하지만, s9이라는  
	s9 := s7
	stringHeader7 := (*reflect.StringHeader)(unsafe.Pointer(&s7))
	stringHeader8 := (*reflect.StringHeader)(unsafe.Pointer(&s8))
	stringHeader9 := (*reflect.StringHeader)(unsafe.Pointer(&s9))
	
	// stringheader 7, 8, 9는 모두 동일한 값을 갖는다. 
	fmt.Println(stringHeader7)
	fmt.Println(stringHeader8)
	fmt.Println(stringHeader9)
	


	// 6. 문자열을 더할때 새로운 문자열이 생성된다. 
	

}

// string은 문자열 자체를 저장하는 것이 아니라 문자열의 주소와 길이를 저장한다. 
// String Header struct는 reflect와 unsafe 패키지를 통해서 사용할 수 있다. 
type StringHeader struct{
	Data uintptr  
	Len int
}
// 참고로 슬라이스는 Cap이라는 필드가 하나 더 추가된다.