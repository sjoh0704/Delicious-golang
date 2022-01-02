package main

import (
	"fmt"
	"test/exinit"
	"test/custompkg"
	"github.com/guptarohit/asciigraph"  // github에서 패키지 가져오기 
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

func main(){
	// 1. custompkg에서 함수 호출하기 
	custompkg.PrintCustom()
	custompkg.PrintCustom2()
 
	// 2. 공개와 비공개 
	s := custompkg.Student{}
	s.Name = "tester"
	s.Age = 10
	// s.score = 20 // 비공개 값이므로 참조가 불가능 
	fmt.Println(s)
	
	// 3. 외부 패키지 사용하기 
	expkg.PrintSample()
	data := []float64{3, 4, 5, 6, 9, 7, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
	
	//4. 패키지 변수의 초기화
	exinit.PrintD()

}