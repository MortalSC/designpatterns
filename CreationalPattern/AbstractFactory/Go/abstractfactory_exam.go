package AbstractFactory

// Producter interface
// Producter is an interface that defines the methods for creating UI components.
type Button interface {
	RenderButton() string
}

type TextBox interface {
	RenderTextBox() string
}

// Modern Products
type ModernButton struct{}

func (mb ModernButton) RenderButton() string {
	return "Rendering a modern button"
}

type ModernTextBox struct{}

func (mtb ModernTextBox) RenderTextBox() string {
	return "Rendering a modern text box"
}

// Classic Products
type ClassicButton struct{}

func (cb ClassicButton) RenderButton() string {
	return "Rendering a classic button"
}

type ClassicTextBox struct{}

func (ctb ClassicTextBox) RenderTextBox() string {
	return "Rendering a classic text box"
}

// UIFactory interface
type UIFactory interface {
	CreateButton() Button
	CreateTextBox() TextBox
}

// ModernUIFactory is a concrete implementation of the UIFactory interface that creates modern UI components.
type ModernUIFactory struct{}

func (m *ModernUIFactory) CreateButton() Button {
	return &ModernButton{}
}

func (m *ModernUIFactory) CreateTextBox() TextBox {
	return &ModernTextBox{}
}

// ClassicUIFactory is a concrete implementation of the UIFactory interface that creates classic UI components.
type ClassicUIFactory struct{}

func (c *ClassicUIFactory) CreateButton() Button {
	return &ClassicButton{}
}

func (c *ClassicUIFactory) CreateTextBox() TextBox {
	return &ClassicTextBox{}
}
