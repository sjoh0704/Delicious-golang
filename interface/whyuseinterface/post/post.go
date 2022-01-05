package post

import (
	"fmt"
)

type Post struct{
	//...
}

func (p *Post) Send(s string) string{
	return fmt.Sprintf("Post sends %s parcel", s)
}
