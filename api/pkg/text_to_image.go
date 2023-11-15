package pkg

import (
	"image"
	"image/color"
	_ "image/jpeg"
	"os"
	"path"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type Color image.Uniform

type TextToImageInput struct {
	Text          string
	Size          float64
	ImageWidth    int
	ImageHeight   int
	TextTopMargin int
	Color         string
	Font          string
}

type TextToImageOutput struct {
	MainTextImage image.Image
	SubTextImage  image.Image
}

func TextToImage(input *TextToImageInput) (*TextToImageOutput, error) {
	// フォントファイルを読み込み

	var fontName string
	if input.Font == "Bold" {
		fontName = "Koruri-Bold.ttf"
	} else if input.Font == "Light" {
		fontName = "Koruri-Light.ttf"
	} else {
		fontName = "Koruri-Regular.ttf"
	}
	ftBinary, err := os.ReadFile(path.Join("..", "koruri", fontName))
	if err != nil {
		return nil, err
	}

	ft, err := truetype.Parse(ftBinary)
	if err != nil {
		return nil, err
	}

	opt := truetype.Options{
		Size: input.Size,
	}

	imageWidth := input.ImageWidth
	imageHeight := input.ImageHeight
	textTopMargin := input.TextTopMargin
	text := input.Text

	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	face := truetype.NewFace(ft, &opt)

	var src *image.Uniform
	if input.Color == "Black" {
		src = image.Black
	} else if input.Color == "Red" {
		// RED
		src = image.NewUniform(color.RGBA{255, 0, 0, 255})
	} else if input.Color == "Blue" {
		// BLUE
		src = image.NewUniform(color.RGBA{0, 0, 255, 255})
	} else if input.Color == "Green" {
		// GREEN
		src = image.NewUniform(color.RGBA{0, 255, 0, 255})
	} else if input.Color == "Yellow" {
		// YELLOW
		src = image.NewUniform(color.RGBA{255, 255, 0, 255})
	} else {
		src = image.White
	}

	dr := &font.Drawer{
		Dst:  img,
		Src:  src,
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
