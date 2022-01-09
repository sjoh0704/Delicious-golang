package main

import (
	"fmt"
)

type Product struct{
	Name string
	Price int
}

func main(){
	
	m := make(map[string]string)

	m["seung"] = "강동"
	m["호옹이"] = "강서"
	m["백두"] = "강북"

	fmt.Println(m)

	m2 := make(map[int]Product)

	m2[3] = Product{"볼펜", 1000}
	m2[4] = Product{"지우개", 1500}
	m2[2] = Product{"샤프", 3000}
	
	// 넣은 순서대로 나오지 않는다. => 해쉬 맵을 사용하기 때문 
	for k, v := range m2{
		fmt.Println(k, v)
	}

	// 삭제하기 
	delete(m2, 2)// 2번 키 값 삭제하기 

	m2[5] = Product{"자", 2000}
	fmt.Println(m2[6]) // 없는 값을 출력하면 디폴트가 나옴
					   //값이 있어서 0인지 없어서 0인지 모른다. 

	for i:=0; i<10; i++{
		if v, ok := m2[i]; ok{ // 값이 있는지 확인하기 
			fmt.Println(v)
		}
	}

	 
}