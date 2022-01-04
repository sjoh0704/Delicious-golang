package main

import (
	"fmt"
	"sort"
)

func main(){

	// 1. 슬라이스를 매개변수로 사용하기 
	s1 := []int{1, 2, 3} //len:3, cap: 3	
	addNum1(s1) // 잘못된 경우 
	fmt.Println(s1)
	addNum2(&s1) // 옳은 경우 
	fmt.Println(s1)
	s1 = addNum3(s1) // 옳은 경우 
	fmt.Println(s1)


	// 2. 슬라이싱 
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr[2:4])

	s2 := arr[2:4]  // 슬라이스 생성 
	fmt.Println(s2, len(s2), cap(s2)) // cap은 시작 인덱스 부터 끝까지 

	s2 = append(s2, 10, 20) // 공간이 부족하므로 새로운 배열을 생성 
	fmt.Println(s2, len(s2), cap(s2)) // cap이 모자라므로 2배 증가한다. 
	
	// 3. 배열을 슬라이스로 바꾸고 싶을 때
	s3 := arr[:]
	fmt.Println(s3, len(s3), cap(s3)) // 단, 슬라이싱을 통해서 만들어진 슬라이스이기 때문에 cap과 len이 동일하다. 


	// 4. 슬라이싱 주의점
	
	arr2 := [20]int{1:1, 2:2, 10:10, 14:14}
	s4 := arr2[1:10]
	s5 := s4[2:19] // s4를 슬라이싱 할때 s4가 아닌 arr2를 쳐다본다. 
	fmt.Println(s4) 
	fmt.Println(s5) //arr2에서 슬라이싱한 결과가 나온다. 

	// 5. cap 사이즈 지정하기 
	s6 := arr2[1:3:10] // 마지막 인덱스는 cap사이즈를 지정한다. 
	fmt.Println(s6, len(s6), cap(s6))


	// 6. 복사?
	s7 := []int{1, 2, 3, 4, 5}
	s8 := s7[:] // 새로운 배열이 생성되는 것이 아니라 s8은 s7이 가르키는 배열을 가르킨다. 
	s8[1] = 100
	fmt.Println(s7) // s8과 s7은 동일한 배열을 가르키므로 s8 요소 수정시 s7도 수정된다. 

	// 복사하기 위해서는
	// 1. 순회를 돌면서 복사하기 
	s9 := make([]int, len(s7))
	for i:=0; i<len(s7); i++{
		s9[i] = s7[i]
	}
	s7[1] = 2 // s7을 바뀌어도 s9가 변경되지 않는다(독립적)
	fmt.Println(s9)

	// 2. append를 이용한 방법(사실상 1번 방법과 동일)
	s10 := append([]int{}, s7...)  // 가변인자를 쓸 때 ...을 사용
	fmt.Println(s10)

	// 3. append를 이용한 방법 => 이 방법이 가장 명확하다. 
	s11 := make([]int, len(s7))
	copy(s11, s7)  // 좌측이 dst, 우측이 src 
	fmt.Println(s11)

	
	// 7. 요소 삭제하기 
	s12 := []int{1, 2, 3, 4, 5, 6}

	// case1 

	delIdx := 2

	// delIdx 이후 값을 앞으로 한칸씩 땡겨주기 
	for i:= delIdx+1; i<len(s12); i++{  
		s12[i-1] = s12[i] 
	}
	// 마지막 값 슬라이싱하기 
	s12 = s12[:len(s12)-1]
	fmt.Println(s12)

	// case2 => 이게 좀 더 잘 보인다. 
	s12 = append(s12[:delIdx], s12[delIdx+1:]...)
	fmt.Println(s12)

	// 8. 요소 삽입

	insertIdx := 2
	insertValue := 100

	// case1
	s12 = append(s12, 0) // 먼저 한칸 추가

	// 뒤로 한칸씩 옮기기 
	for i:=insertIdx; i < len(s12)-1; i++{
		s12[i+1] = s11[i] 
	}

	// 삽입하기 
	s12[insertIdx] = insertValue

	fmt.Println(s12) 

	// case2 한줄로 줄이기 
	s12 = append(s12[:insertIdx+1], append([]int{insertValue}, s12[insertIdx+1:]...)...)
	fmt.Println(s12)

	//case3 copy이용하기
	s12 = append(s12, 0)
	copy(s12[insertIdx+1:], s12[insertIdx:])
	s12[insertIdx] = insertValue
	fmt.Println(s12)


	// 9. 정렬하기 
	// int 값 정렬
	s13 := []int{4, 2, 3, 1, 2, 10, 12}
	sort.Ints(s13)
	fmt.Println(s13)

	// 특정 필드 값으로 정렬하기 
	st1 := Student{"A", 13}
	st2 := Student{"B", 20}
	st3 := Student{"D", 17}
	st4 := Student{"C", 21}
	school := Students{st1, st2, st3, st4}

	sort.Sort(school) // 인터페이스를 배운 후
	fmt.Println(school)
}

// 흔히 하는 실수 
func addNum1(slice []int){
	slice = append(slice, 4)
}

func addNum2(slice *[]int){
	*slice = append(*slice, 4)
}

func addNum3(slice []int) []int{
	
	return append(slice, 4)
}

// 구조체 슬라이스 정렬
type Student struct{
	Name string
	Age int

}

type Students []Student