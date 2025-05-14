package abstract_factory

type IShirt interface {
	setLogo(logo string)
	setSize(size int)
	GetLogo() string
	GetSize() int
}

type Shirt struct {
	logo string
	size int
}

func (a *Shirt) setLogo(logo string) {
	a.logo = logo
}
func (a *Shirt) setSize(size int) {
	a.size = size
}
func (a *Shirt) GetLogo() string {
	return a.logo
}
func (a *Shirt) GetSize() int {
	return a.size
}
