package textimage

import (
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

type FontType string

const (
	KoruriBold FontType = "koruri/Koruri-Bold.ttf"
)

func (u *FontType) String() string {
	return string(*u)
}

func PrepareFontFace(fontType FontType, size float64) (*font.Face, error) {
	ftBinary, err := os.ReadFile(fontType.String())
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
