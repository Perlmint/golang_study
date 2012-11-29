package main

import (
	"image"
	"tour/pic"
	"image/color"
)

type Image struct{
	width int
	height int
	foregroundColor color.Color
	backgroundColor color.Color
	drawfunc func(x, y int) bool
}

func (self Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (self Image) At(x, y int) color.Color {
	var ret color.Color
	if self.drawfunc(x, y) {
		ret = self.foregroundColor
	} else {
		ret = self.backgroundColor
	}
	return ret
}

func (self Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, self.width, self.height)
}

func main() {
	m := Image{50, 50, color.RGBA{0, 0, 0, 255}, color.RGBA{255, 255, 255, 255}, func(x, y int) bool {
		if x == y {
			return true
		}
		return false
	}}
	pic.ShowImage(m)
}

