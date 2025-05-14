package AbstractFactory

import "testing"

func Test_AbstractFactory_Door(t *testing.T) {
	// Create a wooden door factory
	woodenDoorFactory := &WoodenDoorFactory{}

	// Create a wooden door and door handle
	door := woodenDoorFactory.CreateDoor()
	handle := woodenDoorFactory.CreateDoorHandle()

	// Use the door and handle
	door.Open()
	handle.Press()
	door.Close()
}
