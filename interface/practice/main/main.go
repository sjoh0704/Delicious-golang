package main

import (
	"goprojects/test/fedex"
	"goprojects/test/koreaPost"
)

type Sender interface{
	Send(string)
}


func SendBook(name string, sender Sender){
	sender.Send(name)
}

func main(){
	// sender := &fedex.FedexSender{}
	// sender := &koreaPost.KoreaPost{}
	
	var sender Sender = &fedex.FedexSender{}
	var sender2 Sender = &koreaPost.KoreaPost{}


	SendBook("어린 왕자", sender)
	SendBook("그리스 로마 신화", sender)

	SendBook("어린 왕자", sender2)
	SendBook("그리스 로마 신화", sender2)
}