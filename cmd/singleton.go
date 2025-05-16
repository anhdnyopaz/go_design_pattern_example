package main

import (
	"fmt"
	"github.com/anhdnyopaz/go_design_pattern/singleton"
)

func main() {

	for i := 0; i < 30; i++ {
		go singleton.GetInstance()
	}

	fmt.Scanln()

}
