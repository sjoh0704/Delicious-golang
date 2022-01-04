package main

import (
	"fmt"
)

type account struct{
	balance int 
	firstname string
	lastname string
}


// 이 경우에는 리시버가 값이 아닌 포인터여야 한다. 
func (a *account) withdraw(amount int){
	a.balance -= amount
}


func main(){
	
	a := &account{10, "seung", "oh"}
	a.withdraw(3)
	fmt.Println(a)

}