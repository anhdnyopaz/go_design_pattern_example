package abstract_factory

import "fmt"

type ISportFactory interface {
	MakeShoe() IShoe
	MakeShirt() IShirt
}

func GetSportFactory(brand string) (ISportFactory, error) {
	switch brand {
	case "adidas":
		return &Adidas{}, nil
	case "nike":
		return &Nike{}, nil
	}
	return nil, fmt.Errorf("wrong factory name: %s", brand)
}
