package arm

import (
	"errors"
)

type Arm interface {
	GetName() string
	GetPower() int
	SetName(name string)
	SetPower(power int)
}

func NewGun(name string) (Arm, error) {
	switch name {
	case "ak47":
		return NewAK47(), nil
	case "m4a1":
		return NewM4A1(), nil
	default:
		return nil, errors.New("no specified gun")
	}
}

type gun struct {
	name  string
	power int
}

func (g *gun) SetName(name string) {
	(*g).name = name
}

func (g *gun) SetPower(power int) {
	(*g).power = power
}

func (g gun) GetName() string {
	return g.name
}

func (g gun) GetPower() int {
	return g.power
}
