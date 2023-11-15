package textimage

import (
	"context"
	"image"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type GenerateTextImageInput struct {
	FontType FontType
	Text     string
	Size     float64
	Width    int
	Height   int
	TopMagin int
}

func GenerateTextImage(ctx context.Context, input *GenerateTextImageInput) (image.Image, error) {
	face, err := PrepareFontFace(input.FontType, input.Size)
	if err != nil {
		return nil, err
	}

	img := image.NewRGBA(image.Rect(0, 0, input.Width, input.Height))

	dr := &font.Drawer{
		Dst:  img,
		Src:  image.Black,
		Face: *face,
		Dot:  fixed.Point26_6{},
	}
	dr.Dot.X = (fixed.I(input.Width) - dr.MeasureString(input.Text)) / 2
	dr.Dot.Y = fixed.I(input.TopMagin)
	dr.DrawString(input.Text)

	return img, nil
}
