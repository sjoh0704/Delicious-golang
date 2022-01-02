// 대문자는 공개, 소문자는 비공개를 의미한다. 

package custompkg

import (
	"fmt"
)

type Student struct {
	Name string  // 공개
	Age int // 공개 
	score int // 비공개
} 

func PrintCustom2(){
	PrintCustom()
	fmt.Println("custom Print2")
}
