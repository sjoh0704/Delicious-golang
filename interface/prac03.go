// 포함된 인터페이스 

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

// 주목 주목

// Interface In Interface 
type ReadeWriter interface{
	Reader // Reader의 메서드 집합을 포함한다. 
	Writer // Wrtier의 메서드 집합을 포함한다. 
}
// 즉 Read(), Write(), Close() 메서드를 모두 갖고 있어야 이 인터페이스를 사용할 수 있음


