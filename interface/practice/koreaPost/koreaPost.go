package koreaPost

import (
	"fmt"
)

type KoreaPost struct {

}

func (k *KoreaPost) Send(parcel string){
	fmt.Println("koreapost sends", parcel)
}

