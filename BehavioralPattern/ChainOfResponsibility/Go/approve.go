package chainofresponsibility

// IApprover is the interface for the approver in the chain of responsibility pattern.
type IApprover interface {
	// SetNext sets the next approver in the chain.
	SetNext(approver IApprover) IApprover
	// Approve processes the request for approval.
	Approve(amount float64)
}

// Approver is the base struct for the approver in the chain of responsibility pattern.
type Approver struct {
	limit float64
	next  IApprover
}

// NewApprover creates a new approver with the given limit.
func NewApprover(limit float64) *Approver {
	return &Approver{
		limit: limit,
	}
}

func (a *Approver) SetNext(approver IApprover) IApprover {
	a.next = approver
	return a
}

// Approve processes the request for approval.
func (a *Approver) Approve(amount float64) {
	//Approve empty implements the operation of approving the amount, and the specific logic is implemented by the subclass
}

// -------------------------------------------------------

// The Manager structure represents the approver of the manager position
type Manager struct {
	*Approver
}

func NewManager(limit float64) *Manager {
	return &Manager{
		Approver: NewApprover(limit),
	}
}

func (m *Manager) Approve(amount float64) {
	if amount <= m.limit {
		println("Manager approved the amount:", amount)
	} else if m.next != nil {
		m.next.Approve(amount)
	} else {
		println("No one can approve the amount:", amount)
	}
}

// -------------------------------------------------------

// The Director structure represents the approver of the director position
type Director struct {
	*Approver
}

func NewDirector(limit float64) *Director {
	return &Director{
		Approver: NewApprover(limit),
	}
}

func (d *Director) Approve(amount float64) {
	if amount <= d.limit {
		println("Director approved the amount:", amount)
	} else if d.next != nil {
		d.next.Approve(amount)
	} else {
		println("No one can approve the amount:", amount)
	}
}

// -------------------------------------------------------

// CFO structure represents the approver of the CFO position
type CFO struct {
	*Approver
}

func NewCFO(limit float64) *CFO {
	return &CFO{
		Approver: NewApprover(limit),
	}
}

func (c *CFO) Approve(amount float64) {
	if amount <= c.limit {
		println("CFO approved the amount:", amount)
	} else if c.next != nil {
		c.next.Approve(amount)
	} else {
		println("No one can approve the amount:", amount)
	}
}

// -------------------------------------------------------

// How to use the chain of responsibility pattern
// Simple example of how to use the chain of responsibility pattern
// func main() {
// 	manager := NewManager(1000)
// 	director := NewDirector(5000)
// 	cfo := NewCFO(10000)
// 	manager.SetNext(director).SetNext(cfo)
// 	// Test the chain of responsibility pattern
// 	manager.Approve(500)   // Manager approved the amount: 500
// 	manager.Approve(2000)  // Director approved the amount: 2000
// 	manager.Approve(7000)  // CFO approved the amount: 7000
// 	manager.Approve(15000) // No one can approve the amount: 15000
// }

// Also you can see the test file for more examples
