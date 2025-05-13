package main

import "fmt"

// 基础接口：形状
type Shape interface {
	Area() float64
}

// 长方形
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 正方形（继承长方形，但强制宽高相等）
type Square struct {
	Rectangle // 嵌入而非继承
}

func NewSquare(side float64) Square {
	return Square{
		Rectangle: Rectangle{
			Width:  side,
			Height: side,
		},
	}
}

// 使用组合而非继承，避免违反 LSP
func PrintArea(s Shape) {
	fmt.Printf("面积: %.2f\n", s.Area())
}

func main() {
	rect := Rectangle{Width: 3, Height: 4}
	square := NewSquare(5)

	PrintArea(rect)   // 输出: 面积: 12.00
	PrintArea(square) // 输出: 面积: 25.00
}
