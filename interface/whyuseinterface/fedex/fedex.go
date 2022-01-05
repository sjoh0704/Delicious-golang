package fedex

import (
	"fmt"
)

type Fedex struct {
	// ... 
}

func (f *Fedex) Send(s string) string{
	return fmt.Sprintf("Fedex sends %s parcel", s)
}

