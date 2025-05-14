package factory

type Gun struct {
	name  string
	power int
}

func (m Gun) SetName(name string) {
	m.name = name
}

func (m Gun) SetPower(power int) {
	m.power = power

}

func (m Gun) GetName() string {
	return m.name
}

func (m Gun) GetPower() int {
	return m.power
}
