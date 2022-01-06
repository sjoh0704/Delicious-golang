package main

import (
	"fmt"
	"math/rand" // 슬래시 뒷부분이 코드에서 사용할 수 있는 패키지명이 된다. 
	// htemplate "html/template" // 별칭 짓기 
	_ "strings" // 일반적으로 strings를 사용해야 에러가 나지 않지만 언더바를 이용하면 에러가 나지 않는다. 
				// 주로 패키지 초기화에 따른 부가효과를 위해서 사용한다. 
)

func main(){
	fmt.Println(rand.Int())
}