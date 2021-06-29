package transport

import (
	"errors"
)

type Transport interface {
	Deliver()
}

func CreateTransportWith(s string) (Transport, error) {
	switch s {
	case "road":
		return truck{}, nil // if initial logic are complex, use function wrap it. eg: NewTruck()
	case "sea":
		return NewShip(), nil
	}
	return nil,errors.New("unknown logistics")
}
