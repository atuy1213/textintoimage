package textimage

import (
	"os"
	"path"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

type Font int

const (
	Regular Font = iota + 1
	Light
	Bold
)

func NewFont(font string) *Font {
	var ret Font
	switch font {
	case "regular":
		ret = Regular
	case "light":
		ret = Light
	case "bold":
		ret = Bold
	default:
		ret = Regular
	}
	return &ret
}

func (u *Font) Path() string {
	switch *u {
	case Regular:
		return "koruri/Koruri-Regular.ttf"
	case Light:
		return "koruri/Koruri-Light.ttf"
	case Bold:
		return "koruri/Koruri-Bold.ttf"
	default:
		return "koruri/Koruri-Regular.ttf"
	}
}

func (u *Font) Face(size float64) (*font.Face, error) {
	ftBinary, err := os.ReadFile(path.Join("asset", u.Path()))
	if err != nil {
		return nil, err
	}
	ft, err := truetype.Parse(ftBinary)
	if err != nil {
		return nil, err
	}
	opt := &truetype.Options{
		Size: size,
	}
	face := truetype.NewFace(ft, opt)
	return &face, nil
}
