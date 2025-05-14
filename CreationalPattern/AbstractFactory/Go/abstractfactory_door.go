package AbstractFactory

// DoorFactory is an interface that defines the methods for creating door and door handle objects.
type DoorFactory interface {
	CreateDoor() Door
	CreateDoorHandle() DoorHandle
}

// Door is an interface that defines the methods for a door object.
type Door interface {
	Open()
	Close()
}

// DoorHandle is an interface that defines the methods for a door handle object.
type DoorHandle interface {
	Press()
}

// WoodenDoorFactory is a concrete implementation of the DoorFactory interface that creates wooden doors and door handles.
type WoodenDoorFactory struct{}

// CreateDoor creates a wooden door.
func (wdf *WoodenDoorFactory) CreateDoor() Door {
	return &WoodenDoor{}
}

func (wdf *WoodenDoorFactory) CreateDoorHandle() DoorHandle {
	return &WoodenDoorHandle{}
}

// WoodenDoor is a concrete implementation of the Door interface that represents a wooden door.
type WoodenDoor struct{}

func (wd *WoodenDoor) Open() {
	println("Opening wooden door")
}

func (wd *WoodenDoor) Close() {
	println("Closing wooden door")
}

// WoodenDoorHandle is a concrete implementation of the DoorHandle interface that represents a wooden door handle.
type WoodenDoorHandle struct{}

func (wdh *WoodenDoorHandle) Press() {
	println("Pressing wooden door handle")
}

// -------------------------------------------------------------------------

// You can add more concrete factories and products here, such as:
// - MetalDoorFactory: A factory that creates metal doors and door handles.
// - MetalDoor: A concrete implementation of the Door interface that represents a metal door.
// - MetalDoorHandle: A concrete implementation of the DoorHandle interface that represents a metal door handle.

// - GlassDoorFactory: A factory that creates glass doors and door handles.
// - GlassDoor: A concrete implementation of the Door interface that represents a glass door.
// - GlassDoorHandle: A concrete implementation of the DoorHandle interface that represents a glass door handle.
