package main

import "github.com/anhdnyopaz/go_design_pattern/adapter"

func main() {
	client := &adapter.Client{}

	mac := &adapter.Mac{}

	client.InsertLightingConnectorIntoComputer(mac)

	windowsMachine := &adapter.Windows{}

	windowsMachineAdapter := &adapter.WindowsAdapter{
		WindowsMachine: windowsMachine,
	}

	client.InsertLightingConnectorIntoComputer(windowsMachineAdapter)

}
