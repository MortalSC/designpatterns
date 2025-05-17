package templatmethod

import "testing"

func Test_TemplateMethod_For_Coffee(t *testing.T) {
	coffee := &Coffee{}
	coffee.BeverageTemplate.BrewBeverage = coffee
	coffee.Prepare()
}

func Test_TemplateMethod_For_Tea(t *testing.T) {
	tea := &Tea{}
	tea.BeverageTemplate.BrewBeverage = tea
	tea.Prepare()
}

func Test_TemplateMethod_For_Milk(t *testing.T) {
	milk := &Milk{}
	milk.BeverageTemplate.BrewBeverage = milk
	milk.Prepare()
}
