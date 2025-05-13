package builder

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func (i IglooBuilder) setWindowType() {
	i.windowType = "igloo window"
}

func (i IglooBuilder) setDoorType() {
	i.doorType = "igloo door"
}

func (i IglooBuilder) setNumFloor() {
	i.floor = 3
}

func (i IglooBuilder) getHouse() House {
	return House{
		WindowType: i.windowType,
		DoorType:   i.doorType,
		Floor:      i.floor,
	}
}

func newIglooBuilder() IBuilder {
	return IglooBuilder{}
}
