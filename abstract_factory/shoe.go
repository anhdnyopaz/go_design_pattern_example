package abstract_factory

type IShoe interface {
	setLogo(logo string)
	GetLogo() string
	setSize(size int)
	GetSize() int
}

type Shoe struct {
	logo string
	size int
}

func (a *Shoe) setLogo(logo string) {
	a.logo = logo
}

func (a *Shoe) GetLogo() string {
	return a.logo
}

func (a *Shoe) setSize(size int) {
	a.size = size
}

func (a *Shoe) GetSize() int {
	return a.size
}
