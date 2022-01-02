package main

import (
	"fmt"
	"test/custompkg"
	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

func main(){
	custompkg.PrintCustom2()
	expkg.PrintSample()
	data := []float64{3, 4, 5, 6, 9, 7, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
	
	// 값의 공개와 비공개 
	s := custompkg.Student{}
	s.Name = "tester"
	s.Age = 10
	// s.score = 20 // 비공개 값이므로 참조가 불가능 
	fmt.Println(s)

}