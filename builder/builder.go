package builder

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func GetBuilder(builderType string) IBuilder {

	switch builderType {
	case "normal":
		return newNormalBuilder()
	case "igloo":
		return newIglooBuilder()
	default:
		return nil

	}
}
