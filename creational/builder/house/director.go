package house

import (
	"errors"
)

type director struct {
	builder iBuilder
}

func newDirector(builder iBuilder) *director {
	return &director{builder: builder}
}

func (d *director) setBuilder(builder iBuilder) {
	d.builder = builder
}

func (d *director) buildHouse() (house, error) {
	if d.builder != nil {
		d.builder.buildDoor()
		d.builder.buildFloor()
		d.builder.buildWall()
		d.builder.buildWindow()
		return d.builder.getHouse(), nil
	}
	return house{}, errors.New("no specify builder")
}
