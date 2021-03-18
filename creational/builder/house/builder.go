package house

import (
	"errors"
)

type iBuilder interface {
	buildWall()
	buildWindow()
	buildFloor()
	buildDoor()
	getHouse() house
}

func getBuilder(texture string) (iBuilder,error) {
	switch texture {
	case "wood":
		return &woodenBuilder{},nil
	case "steel":
		return &steelBuilder{},nil
	default:
		return nil,errors.New("no specified builder")
	}
}