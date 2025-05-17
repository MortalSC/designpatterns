package templatmethod

// ---------- Abstract template ----------
type BeverageTemplate struct {
	// Embed（嵌入） interfaces to achieve polymorphism（多态）
	// (Go has no inheritance（继承） and uses composition simulation（模拟）)
	BrewBeverage
}

// Template method: Define the algorithm skeleton
func (b *BeverageTemplate) Prepare() {
	// 1. Boil water
	b.BoilWater()
	// 2. Brew the beverage
	b.Brew()
	// 3. Pour into cup
	b.PourInCup()
	// 4. Add condiments
	if b.NeedCondiments() { // Hook method (Optional step)
		// 5. Add condiments
		b.AddCondiments()
	}
}

func (b *BeverageTemplate) BoilWater() {
	println("Boiling water")
}

func (b *BeverageTemplate) PourInCup() {
	println("Pouring into cup")
}

// ---------- Abstract interface (defining the methods that subclasses need to implement) ----------
type BrewBeverage interface {
	// Brew the beverage
	Brew()
	// Add condiments
	AddCondiments()
	// Hook method: Determine whether to add condiments
	NeedCondiments() bool
}

// Default hook method implementation (Optional override for subclasses)
func (b *BeverageTemplate) NeedCondiments() bool {
	return true
}

// ---------- Specific subclasses: Coffee ----------
type Coffee struct {
	BeverageTemplate
}

func (c *Coffee) Brew() {
	println("Dripping coffee through filter")
}

func (c *Coffee) AddCondiments() {
	println("Adding sugar and milk")
}

// Override the hook method
func (c *Coffee) NeedCondiments() bool {
	// Coffee does need condiments
	return false
}

// ---------- Specific subclasses: Tea ----------
type Tea struct {
	BeverageTemplate
}

func (t *Tea) Brew() {
	println("Steeping the tea")
}

func (t *Tea) AddCondiments() {
	println("Adding lemon")
}

// Override the hook method
func (t *Tea) NeedCondiments() bool {
	// Tea does not need condiments
	return false
}

// ---------- Specific subclasses: Milk ----------
type Milk struct {
	BeverageTemplate
}

func (m *Milk) Brew() {
	println("Heating the milk")
}

func (m *Milk) AddCondiments() {
	println("Adding sugar")
}

// Override the hook method
func (m *Milk) NeedCondiments() bool {
	// Milk does not need condiments
	return false
}

// ---------- Example usage ----------
// func main() {
// 	// Create a coffee instance
// 	coffee := &Coffee{}
// 	// Prepare the coffee
// 	coffee.Prepare()

// 	println("")

// 	// Create a tea instance
// 	tea := &Tea{}
// 	// Prepare the tea
// 	tea.Prepare()

// 	println("")

// 	// Create a milk instance
// 	milk := &Milk{}
// 	// Prepare the milk
// 	milk.Prepare()
// }
