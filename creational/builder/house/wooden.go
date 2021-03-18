package house

type woodenBuilder struct {
	house
}

func (w *woodenBuilder) buildWall() {
	w.house.wallType = "wood"
}

func (w *woodenBuilder) buildWindow() {
	w.house.windowType = "wood"
}

func (w *woodenBuilder) buildFloor() {
	w.house.floor = 1
}

func (w *woodenBuilder) buildDoor() {
	w.house.doorType = "wood"
}

func (w *woodenBuilder) getHouse() house {
	return w.house
}
