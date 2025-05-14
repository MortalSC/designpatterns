package simplefactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SingleFactroy(t *testing.T) {
	assert := assert.New(t)

	{
		circle := NewShape(ShapeTypeCircle)
		assert.Equal(circle.Draw(), "Drawing a Circle")
	}

	{
		square := NewShape(ShapeTypeSquare)
		assert.Equal(square.Draw(), "Drawing a Square")
	}
}
