package phone

import "fmt"

type typeC struct {

}

func (receiver typeC) insertInto() string {
	fmt.Println("type-c insert into")
	return "type-c"
}

