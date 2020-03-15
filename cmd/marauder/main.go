package main

import (
	"marauders-map-client-desktop/internal/wsclient"
)

func main() {

	// Deploy for persistence
	// this setups home directory folder for the program
	// folder strcuture & persist mechanism
	// deploy.Deploy()

	// Start connection and communication with server
	wsclient.StartCommunications()

}
