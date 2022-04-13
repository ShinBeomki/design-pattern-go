package FactoryPattern

import "fmt"

type Shape interface {
	Draw()
}

type Circle struct{}

type Square struct{}

type Rectangle struct{}

func NewCircle() *Circle {
	return &Circle{}
}

func (c *Circle) Draw() {
	fmt.Println("Circle Draw() method.")
}

func NewSquare() *Square {
	return &Square{}
}

func (s *Square) Draw() {
	fmt.Println("Square Draw() method.")
}

func NewRectangle() *Rectangle {
	return &Rectangle{}
}

func (r *Rectangle) Draw() {
	fmt.Println("Rectangle Draw() method.")
}

type ShapeFactory struct{}

func NewShapeFactory() *ShapeFactory {
	return &ShapeFactory{}
}

func (sf *ShapeFactory) GetShape(shapeType string) Shape {
	switch shapeType {
	case "circle":
		return NewCircle()
	case "square":
		return NewSquare()
	case "rectangle":
		return NewRectangle()
	default:
		return nil
	}
}
