package exinit

import (
	"fmt"
)

var (
	a = c + b // 3
	b = f()  // 1
	c = f()  // 2
	d = 3  //4
) 

func init(){ //5
	d++
	fmt.Println("init 호출", d)
}

func f() int{
	d++
	fmt.Println("f() 호출", d)
	return d
}

func PrintD(){
	fmt.Println("d:", d)
}