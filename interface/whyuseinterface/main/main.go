package main

import (
	"fmt"
	"interfacepractice/fedex"
	"interfacepractice/post"
)

type Sender interface{ // 내부에 정의된 메서드만을 갖는 struct만 해당된다. 더 많거나 적으면 안된다. 
	Send(string) string  //내부에 메서드가 들어간다. 
}

// 인터페이스를 이용하면 내부에 정의된 함수를 수정하지 않고도 사용할 수 있다. => 추상화 
func sendParcel(s Sender, parcel string){
	fmt.Println(s.Send(parcel))
}

func main(){

	var sender1 Sender = &fedex.Fedex{} // 인터페이스는 인스턴스의 주솟값을 가르킨다. 
	sendParcel(sender1, "장난감")
	sendParcel(sender1, "노트북")
	sendParcel(sender1, "핸드폰")

	var sender2 Sender = &post.Post{}
	sendParcel(sender2, "장난감")
	sendParcel(sender2, "노트북")
	sendParcel(sender2, "핸드폰")

}