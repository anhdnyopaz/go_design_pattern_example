package builder

type Director struct {
	builder IBuilder
}

func NewDirector(builder IBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) SetBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) BuildHouse() House {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}
