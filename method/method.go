package main

import (
	"fmt"
)

type Account struct{
	Balance int
}

type myInt int

// function
func withdrawFunc(a *Account, amount int){
	a.Balance -= amount
}

// method
// Account 타입에 포함된 method 생성 
func (a *Account) withdrawMethod(amount int){
	a.Balance -= amount
}

func (m myInt) add(a int) myInt{
	return myInt(int(m) + a)
}

func main(){
	a := &Account{100}

	withdrawFunc(a, 10)
	fmt.Println(a)

	// method 사용하기 
	a.withdrawMethod(30)
	fmt.Println(a)

	b := myInt(10)
	b = b.add(20) // 값으로 받았기 때문에 
	fmt.Println(b)
}