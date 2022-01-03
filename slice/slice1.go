package main

import (
	"fmt"
)

func main(){
	// 슬라이스 선언하기 
	//1
	var a []int
	// a := []int{}
	fmt.Println(a)

	//2 
	if len(a) == 0{
		fmt.Println("슬라이스가 비어 있습니다.")
	}else{
		fmt.Println(a)
	}

	//3 주의
	c1 := []int{1, 2, 3} // 슬라이스
	// c2 := [...]int{1, 2, 3}
	// c3 := [3]int{1, 2, 3}
	fmt.Println(c1)

	// 4. make
	d := make([]int, 10) // 10개의 요소를 갖는 int 타입 슬라이스
	fmt.Println(d)

	// 5. 슬라이스 순회 
	e := []int{1, 2, 3, 4}
	for _, v := range e{
		fmt.Println(v)
	}

	// 6. append를 통해서 슬라이스에 값 추가하기 

	e = append(e, 1, 2, 3)
	for _, v := range e{
		fmt.Println(v)
	}

	// 7. 함수 
	t1 := [5]int{1, 2, 3, 4, 5} // 8x5 = 40 바이트 
	t2 := make([]int, 3, 5) // struct의 크기가 24 byte (포인터일뿐)

	changeArrayWrong(t1)  // 안바뀜
	fmt.Println(t1)

	changeArrayRight(&t1)  // 바뀜 
	fmt.Println(t1)

	changeSlice(t2)  // 바뀜 
	fmt.Println(t2)

	//8. append 심화

	//case1 공간이 충분할 때 

	s1 := make([]int, 3, 5) // 2개의 공간이 남는다. 
	s2 := append(s1, 4, 5) // 2개의 공간이 남아있으므로 2개의 요소를 넣기에 충분하다. 
							// 따라서 기존 배열을 사용 
	s1[1] = 100	// s1과 s2는 같은 배열을 공유하게 되고, s1의 요소를 바꾸면 s2의 요소 또한 바뀐다.  
	fmt.Println(s1)  // 0 100 0 
	fmt.Println(s2)  // 0 100 0 4 5

	//case2 공간이 부족할 때 
	
	s3 := make([]int, 3, 4)
	s4 := append(s3, 4, 5) // 공간이 부족하기 때문에 새로운 배열을 생성한다. 
	s3[1] = 100  // s3와 s4는 다른 배열을 사용하기 때문에 한 값을 변경해도 다른 값에 영향을 끼치지 않는다. 
	fmt.Println(s3) // 0 100 0
	fmt.Println(s4) // 0 0 0 4 5


}

func changeArrayWrong(arr [5]int){ // 값이 들어가므로 바뀌지 않음. 이부분은 C++과 다르다.  
	arr[1] = 100 
}

func changeArrayRight(arr *[5]int){ // 값이 들어가므로 바뀌지 않음. 
	arr[1] = 100
}


func changeSlice(slice []int){ // 주소가 들어가므로 값이 바뀜 
	slice[1] = 100
}
