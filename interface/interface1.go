// adapter pattern
package main

import (
	"fmt"
)

// A database와 B database는 다른 메서드를 갖는다. 

// A database
type ADatabase struct{
	Cost int 
}

func (a *ADatabase) Get() int{
	return a.Cost
}

func (a *ADatabase) Set(cost int){
	a.Cost = cost
}

// B database
type BDatabase struct{
	Cost int
}

func (b *BDatabase) GetCost() int{
	return b.Cost
}

func (b *BDatabase) SetCost(cost int){
	b.Cost = cost
}

// Database Interface
type Database interface{
	Get() int 
	Set(int)
}

// Database의 cost를 증가시키는 함수 
func IncreaseCostDatabase(db Database, increase int){
	
	fmt.Println("Current Cost:", db.Get())
	db.Set(db.Get() + increase)
	fmt.Println("After increasing:", db.Get())
}


//Wrapper: Adapter 역할을 한다. 
type Wrapper struct{
	*BDatabase
}

func (w *Wrapper) Get() int{
	return w.GetCost()
}

func (w *Wrapper) Set(cost int){
	w.SetCost(cost)
}


func main(){
	var db1 Database = &ADatabase{10}
	IncreaseCostDatabase(db1, 20)

	// 메서드가 다르기 때문에 Error 
	// var dbB Database = &BDatabase{20}  
	
	// wrapper이용하기 
	var db2 Database = &Wrapper{&BDatabase{20}}
	IncreaseCostDatabase(db2, 50)
}