package fedex

import "fmt"

type FedexSender struct {

}

func (f *FedexSender) Send(parcel string){
	fmt.Println("fedex sends", parcel)
}










