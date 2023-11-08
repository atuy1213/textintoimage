package pkg

import (
	"image"
	_ "image/jpeg"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type TextToImageInput struct {
	MainText string
	SubText  string
}

type TextToImageOutput struct {
	MainTextImage image.Image
	SubTextImage  image.Image
}

func TextToImage(input *TextToImageInput) (*TextToImageOutput, error) {
	// フォントファイルを読み込み
	ftBinary, err := os.ReadFile("koruri/Koruri-Bold.ttf")
	if err != nil {
		return nil, err
	}

	ft, err := truetype.Parse(ftBinary)
	if err != nil {
		return nil, err
	}

	opt := truetype.Options{
		Size:              90,
		DPI:               0,
		Hinting:           0,
		GlyphCacheEntries: 0,
		SubPixelsX:        0,
		SubPixelsY:        0,
	}

	imageWidth := 300
	imageHeight := 100
	textTopMargin := 90
	text := input.MainText

	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	face := truetype.NewFace(ft, &opt)
	dr := &font.Drawer{
		Dst:  img,
		Src:  image.Black,
		Face: face,
		Dot:  fixed.Point26_6{},
	}
	dr.Dot.X = (fixed.I(imageWidth) - dr.MeasureString(text)) / 2
	dr.Dot.Y = fixed.I(textTopMargin)
	dr.DrawString(text)

	ret := &TextToImageOutput{
		MainTextImage: img,
	}

	return ret, nil
}
