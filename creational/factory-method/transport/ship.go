package transport

type ship struct{}

func (ship) Deliver() {
	// ship transport logic
}

func NewShip() Transport {
	return ship{}
}
