package phone

import "fmt"

type adapter struct {}

func (a *adapter) insertTypeCPort(p typeCPort) string {
	p.insertInto()
	// convert type-c to micro-usb
	fmt.Println("micro-usb insert into")
	return "micro-usb"
}
