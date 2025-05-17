package chainofresponsibility

import "testing"

func Test_Chain(t *testing.T) {
	// Create the chain of approvers
	manager := NewManager(1000)
	director := NewDirector(5000)
	cfo := NewCFO(10000)

	// Set the next approver in the chain
	manager.SetNext(director).SetNext(cfo)
	// <=>
	// manager.SetNext(director)
	// director.SetNext(cfo)

	// Test the chain of responsibility pattern
	amounts := []float64{500, 1500, 6000, 12000}

	// expectedResults := []string{
	// 	"Manager approved the amount: 500",
	// 	"Director approved the amount: 1500",
	// 	"CFO approved the amount: 6000",
	// 	"No one can approve the amount: 12000",
	// }

	for _, amount := range amounts {
		t.Logf("Requesting approval for amount: %.2f", amount)
		manager.Approve(amount)
	}
}
