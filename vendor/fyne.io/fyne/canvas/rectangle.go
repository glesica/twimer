package canvas

import (
	"image/color"

	"fyne.io/fyne"
)

// Declare conformity with CanvasObject interface
var _ fyne.CanvasObject = (*Rectangle)(nil)

// Rectangle describes a coloured rectangle primitive in a Fyne canvas
type Rectangle struct {
	baseObject

	FillColor   color.Color // The rectangle fill colour
	StrokeColor color.Color // The rectangle stroke colour
	StrokeWidth float32     // The stroke width of the rectangle
}

// NewRectangle returns a new Rectangle instance
func NewRectangle(color color.Color) *Rectangle {
	return &Rectangle{
		FillColor: color,
	}
}
