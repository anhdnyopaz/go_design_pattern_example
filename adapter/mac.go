package adapter

import "fmt"

type Mac struct {
}

func (m *Mac) InsertIntoLightningPost() {
	fmt.Println("Lighting connector is plugged into mac machine.")
}
