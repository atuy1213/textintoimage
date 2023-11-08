package pkg

import (
	"image"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	"os"
)

func textIntoImage() {
	src, err := os.Open("src/text.png")
	if err != nil {
		panic(err)
	}
	defer src.Close()
	dst, err := os.Open("src/sample.jpeg")
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	// デコードしてイメージオブジェクトを準備
	srcImg, _, err := image.Decode(src)
	if err != nil {
		panic(err)
	}
	dstImg, _, err := image.Decode(dst)
	if err != nil {
		panic(err)
	}

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

	// 書き出し用ファイル準備
	outfile, _ := os.Create("dist/out.png")
	defer outfile.Close()
	// 書き出し
	png.Encode(outfile, out)
}
