package builder

type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func (n NormalBuilder) setWindowType() {
	n.windowType = "normal window"
}

func (n NormalBuilder) setDoorType() {
	n.doorType = "normal door"
}

func (n NormalBuilder) setNumFloor() {
	n.floor = 2
}

func (n NormalBuilder) getHouse() House {
	return House{
		WindowType: n.windowType,
		DoorType:   n.doorType,
		Floor:      n.floor,
	}
}

func newNormalBuilder() IBuilder {
	return NormalBuilder{}
}
