package main

import (
	"fmt"
)

type Reader interface{
	Read() (n int, err error)
	Close() error
}

type Writer interface{
	Write(n int, err error)
	Close() error
}

// Interface In Interface 
type ReadeWriter interface{
	Reader // Reader의 메서드 집합을 포함한다. 
	Writer // Wrtier의 메서드 집합을 포함한다. 
}

