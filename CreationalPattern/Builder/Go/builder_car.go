package builder

// ICar
type ICar interface {
	Speed() int
	Brand() string
	Brief()
}

// ICarBuilder
type ICarBuilder interface {
	Wheel(wheel int) ICarBuilder
	Engine(engine string) ICarBuilder
	Speed(seats int) ICarBuilder
	Brand(brand string) ICarBuilder
	Build() ICar
}

// CarProto
type CarProto struct {
	Wheel     int
	Engine    string
	MaxSpeed  int
	BrandName string
}

func (c *CarProto) Speed() int {
	return c.MaxSpeed
}

func (c *CarProto) Brand() string {
	return c.BrandName
}

func (c *CarProto) Brief() {
	println("Brand: ", c.BrandName)
	println("MaxSpeed: ", c.MaxSpeed)
	println("Engine: ", c.Engine)
	println("Wheel: ", c.Wheel)
}

// CarStudio
type CarStudio struct {
	protoType CarProto
}

func NewCarStudio() *CarStudio {
	return &CarStudio{}
}

func (c *CarStudio) Wheel(wheel int) ICarBuilder {
	c.protoType.Wheel = wheel
	return c
}

func (c *CarStudio) Engine(engine string) ICarBuilder {
	c.protoType.Engine = engine
	return c
}

func (c *CarStudio) Speed(speed int) ICarBuilder {
	c.protoType.MaxSpeed = speed
	return c
}
func (c *CarStudio) Brand(brand string) ICarBuilder {
	c.protoType.BrandName = brand
	return c
}

// Build
func (c *CarStudio) Build() ICar {
	return &CarProto{
		Wheel:     c.protoType.Wheel,
		Engine:    c.protoType.Engine,
		MaxSpeed:  c.protoType.MaxSpeed,
		BrandName: c.protoType.BrandName,
	}
}
