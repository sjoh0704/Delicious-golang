package main

import (
	"fmt"
)

type DuckInterface interface {
	// func 없이 method를 적는다. 
	Fly()
	Walk(distance int) int
}

func main(){

}