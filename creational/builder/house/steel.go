package house

type steelBuilder struct {
	house
}

func (s *steelBuilder) buildWall() {
	s.house.wallType = "steel"
}

func (s *steelBuilder) buildWindow() {
	s.house.windowType = "steel"
}

func (s *steelBuilder) buildFloor() {
	s.house.floor = 1
}

func (s *steelBuilder) buildDoor() {
	s.house.doorType = "steel"
}

func (s *steelBuilder) getHouse() house {
	return s.house
}
