//go:build !js

package main

const (
	turnServerAddr = "localhost:17342"
	turnServerURL  = "turn:" + turnServerAddr + "?transport=tcp"
	turnUsername   = "turn_username"
	turnPassword   = "turn_password"
)

func main() {

	turnServer := newTURNServer()
	defer turnServer.Close()

	setupAnsweringAgent()

	setupOfferingAgent()

	select {}
}
