package decorator

type TomatoTopping struct {
	Pizza IPizza
}

func (c *TomatoTopping) GetPrice() int {
	pizzaPrice := c.Pizza.GetPrice()

	return pizzaPrice + 7
}
