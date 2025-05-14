package abstract_factory

type Nike struct{}

type NikeShoe struct {
	Shoe
}

func (n *Nike) MakeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "nike",
			size: 14,
		},
	}
}

type NikeShirt struct {
	Shirt
}

func (n *Nike) MakeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			logo: "nike",
			size: 2,
		},
	}
}
