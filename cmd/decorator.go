package main

import (
	"fmt"
	"github.com/anhdnyopaz/go_design_pattern/decorator"
)

func main() {
	pizza := &decorator.VeggieMania{}

	pizzaWithCheese := &decorator.CheeseTopping{
		Pizza: pizza,
	}
	// Add tomato topping

	pizzaWithCheeseAndTomato := &decorator.TomatoTopping{
		Pizza: pizzaWithCheese,
	}

	fmt.Printf("Price is: %d $", pizzaWithCheeseAndTomato.GetPrice())
}
