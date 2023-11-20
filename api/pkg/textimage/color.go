package textimage

import (
	"image"
	"image/color"
)

type Color int

const (
	black Color = iota + 1
	red
	blue
	green
	yellow
	white
)

func NewColor(color string) *Color {
	var ret Color
	switch color {
	case "black":
		ret = black
	case "red":
		ret = red
	case "blue":
		ret = blue
	case "green":
		ret = green
	case "yellow":
		ret = yellow
	case "white":
		ret = white
	default:
		ret = black
	}
	return &ret
}

func (u *Color) Uniform() *image.Uniform {
	switch *u {
	case black:
		return image.Black
	case red:
		return image.NewUniform(color.RGBA{255, 0, 0, 255})
	case blue:
		return image.NewUniform(color.RGBA{0, 0, 255, 255})
	case green:
		return image.NewUniform(color.RGBA{0, 255, 0, 255})
	case yellow:
		return image.NewUniform(color.RGBA{255, 255, 0, 255})
	case white:
		return image.White
	default:
		return image.Black
	}
}

func (u *Color) String() string {
	switch *u {
	case black:
		return "black"
	case red:
		return "red"
	case blue:
		return "blue"
	case green:
		return "green"
	case yellow:
		return "yellow"
	case white:
		return "white"
	default:
		return "black"
	}
}
