package adapter

import "fmt"

type WindowsAdapter struct {
	WindowsMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPost() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.WindowsMachine.insertIntoUSBPort()

}
