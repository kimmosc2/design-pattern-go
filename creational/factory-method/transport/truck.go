package transport

type truck struct {
}

func (t truck) Deliver() {
	// truck deliver logic
}

func NewTruck() Transport {
	// here can addition some complex logic
	return truck{}
}
