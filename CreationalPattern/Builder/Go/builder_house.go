package builder

import "fmt"

// House Product
type House struct {
	Walls     string
	Door      string
	Windows   string
	HasGarage bool
}

func (h House) Show() {
	fmt.Printf("房屋结构：墙壁=%s, 门=%s, 窗户=%s, 车库=%v\n",
		h.Walls, h.Door, h.Windows, h.HasGarage)
}

// HouseBuilder Interface
type HouseBuilder interface {
	BuildWalls()
	InstallDoor()
	InstallWindows()
	BuildGarage()
	GetHouse() House
}

// WoodenHouseBuilder Concrete Builder
type WoodenHouseBuilder struct {
	house House
}

func (b *WoodenHouseBuilder) BuildWalls() {
	b.house.Walls = "木墙"
}

func (b *WoodenHouseBuilder) InstallDoor() {
	b.house.Door = "木门"
}

func (b *WoodenHouseBuilder) InstallWindows() {
	b.house.Windows = "木窗"
}

func (b *WoodenHouseBuilder) BuildGarage() {
	b.house.HasGarage = true
}

func (b *WoodenHouseBuilder) GetHouse() House {
	return b.house
}

// BrickHouseBuilder Concrete Builder
type BrickHouseBuilder struct {
	house House
}

func (b *BrickHouseBuilder) BuildWalls() {
	b.house.Walls = "砖墙"
}

func (b *BrickHouseBuilder) InstallDoor() {
	b.house.Door = "铁门"
}

func (b *BrickHouseBuilder) InstallWindows() {
	b.house.Windows = "玻璃窗"
}

func (b *BrickHouseBuilder) BuildGarage() {
	b.house.HasGarage = false
}

func (b *BrickHouseBuilder) GetHouse() House {
	return b.house
}

// HouseDirector Director
type HouseDirector struct {
	builder HouseBuilder
}

func (d *HouseDirector) SetBuilder(builder HouseBuilder) {
	d.builder = builder
}

func (d *HouseDirector) ConstructHouse() House {
	d.builder.BuildWalls()
	d.builder.InstallDoor()
	d.builder.InstallWindows()
	d.builder.BuildGarage() // Optional
	return d.builder.GetHouse()
}
