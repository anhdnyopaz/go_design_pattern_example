package adapter

import "fmt"

type Client struct {
}

func (c Client) InsertLightingConnectorIntoComputer(com Computer) string {
	fmt.Println("Client inserts Lighting Connector into computer")
	com.InsertIntoLightningPost()
}
