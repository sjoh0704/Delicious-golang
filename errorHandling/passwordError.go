package main

import (
	"fmt"
)

type PasswordError struct{
	Len int
	RequiredLen int
}

// error interface가 되기 위해서 필요한 메서드
func (passwordError PasswordError) Error() string{
	return "암호 길이가 짧습니다."
}

func RegisterAccount(id string, password string) error {
	if len(password) < 8{
		// error interface가 되기 위해서는 Error() string이 필요 
		// return fmt.Errorf("암호 길이가 짧습니다. %d", 8) => 이러한 방식도 있지만 디테일한 정보를 얻기 힘들다. 
		return PasswordError{len(password), 8}
	}else{
		return nil
	}

}

func main(){

	err := RegisterAccount("myid", "1234")
	if err != nil{
		// fmt.Printf("%v\n", err)
		// 인터페이스 타입 변환을 이용해서 자세한 정보를 가져 올 수 있다. 
		if errInfo, ok := err.(PasswordError); ok{
			fmt.Printf("%v, Len: %d, RequiredLen: %d\n", errInfo, errInfo.Len, errInfo.RequiredLen)
			return 
		}
		return 
	}else{
		fmt.Println("회원가입 성공입니다.")
	}

}