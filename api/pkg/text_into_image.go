package pkg

import (
	"image"
	"image/draw"
	_ "image/jpeg"
)

type TextIntoImageInput struct {
	SrcImage      image.Image
	MainTextImage image.Image
	SubTextImage  image.Image
}

type TextIntoImageOutput struct {
	Image image.Image
}

func TextIntoImage(input *TextIntoImageInput) (*TextIntoImageOutput, error) {
	// デコードしてイメージオブジェクトを準備
	srcImg := input.MainTextImage
	dstImg := input.SrcImage

	// 書き出し用のイメージを準備
	outRect := image.Rectangle{image.Pt(0, 0), dstImg.Bounds().Size()}
	out := image.NewRGBA(outRect)

	// 描画する
	// 元画像をまず描く
	dstRect := image.Rectangle{image.Pt(0, 0), dstImg.Bounds().Size()}
	draw.Draw(out, dstRect, dstImg, image.Pt(0, 0), draw.Src)
	// 上書きする
	srcRect := image.Rectangle{image.Pt(0, 0), srcImg.Bounds().Size()}
	draw.Draw(out, srcRect, srcImg, image.Pt(0, 0), draw.Over)

	ret := &TextIntoImageOutput{
		Image: out,
	}

	return ret, nil
}
