package main

import (
	"fmt"
)

type User struct{
	Name string
	Age int
}

func (u *User) String() string{
	return fmt.Sprintf("name: %s\n age: %d", u.Name, u.Age)
}

func (v *VIPUser) getVIPLevel() int{
	return v.VIPLevel
}


type VIPUser struct{
	*User  // embedded field, 변수 없이 사용해도 된다. 
	VIPLevel int
}

func main(){

	u1 := &User{"sseung", 20}
	fmt.Println(u1.String())

	vip1 := &VIPUser{&User{"cindy", 20}, 3}

	fmt.Println(vip1.getVIPLevel())
	// String 메서드는 User에 있음에도 불구하고 VIP의 임베디드 필드이기 때문에
	// VIP에서 사용할 수 있다. 
	// 상속과 비슷하지만 상속은 아니다! => 단지 편리를 위해서 만든 것
	// 만약 VIP를 기준으로 String 메서드를 만든다면, User의 메서드와는 무관하다. 
	fmt.Println(vip1.String())


}