package simplefactory

type ShapeType string

const (
	ShapeTypeCircle ShapeType = "circle"
	ShapeTypeSquare ShapeType = "square"
	// ... add more shape types as needed
)

type Shape interface {
	Draw() string
}

type Circle struct{}

func (c *Circle) Draw() string {
	return "Drawing a Circle"
}

type Square struct{}

func (s *Square) Draw() string {
	return "Drawing a Square"
}

// ShapeFactory is a simple factory for creating shapes
func NewShape(shapeType ShapeType) Shape {
	switch shapeType {
	case ShapeTypeCircle:
		return &Circle{}
	case ShapeTypeSquare:
		return &Square{}
	default:
		return nil
	}
}
