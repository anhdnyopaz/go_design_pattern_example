package main

import (
	"fmt"
	"github.com/anhdnyopaz/go_design_pattern/factory"
)

func main() {
	ak47, _ := factory.GetGun("ak47")
	musket, _ := factory.GetGun("musket")

	printDetail(ak47)
	printDetail(musket)
}

func printDetail(g factory.IGun) {
	fmt.Printf("Gun: %s \n", g.GetName())
	fmt.Printf("Power: %d \n", g.GetPower())
}
