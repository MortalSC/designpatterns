package AbstractFactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AbstractFactory_Exam(t *testing.T) {
	assert := assert.New(t)

	{
		// Create a modern UI factory
		modernUIFactory := &ModernUIFactory{}

		// Create a modern button and text box
		modernButton := modernUIFactory.CreateButton()
		modernTextBox := modernUIFactory.CreateTextBox()

		// Use the modern button and text box
		assert.Equal(modernButton.RenderButton(), "Rendering a modern button")
		assert.Equal(modernTextBox.RenderTextBox(), "Rendering a modern text box")
	}

	{
		// Create a classic UI factory
		classicUIFactory := &ClassicUIFactory{}

		// Create a classic button and text box
		classicButton := classicUIFactory.CreateButton()
		classicTextBox := classicUIFactory.CreateTextBox()

		// Use the classic button and text box
		assert.Equal(classicButton.RenderButton(), "Rendering a classic button")
		assert.Equal(classicTextBox.RenderTextBox(), "Rendering a classic text box")
	}
}
