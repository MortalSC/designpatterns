package builder

import "testing"

func Test_Builder_House(t *testing.T) {
	// Create a wooden house builder
	woodenHouseBuilder := &WoodenHouseBuilder{}

	// Create a house director with the wooden house builder
	houseDirector := HouseDirector{builder: woodenHouseBuilder}

	// Construct the house
	houseDirector.ConstructHouse()

	// Get the constructed house
	house := woodenHouseBuilder.GetHouse()

	// Use the constructed house
	if house.Walls != "木墙" {
		t.Errorf("Expected walls to be '木墙', got '%s'", house.Walls)
	}
	if house.Door != "木门" {
		t.Errorf("Expected door to be '木门', got '%s'", house.Door)
	}
	if house.Windows != "木窗" {
		t.Errorf("Expected windows to be '木窗', got '%s'", house.Windows)
	}
	if !house.HasGarage {
		t.Error("Expected garage to be built")
	}
}
