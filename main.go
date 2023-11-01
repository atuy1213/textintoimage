package main

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

func main() {
	text()
	textIntoImage()
}

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

func text() {
	// フォントファイルを読み込み
	ftBinary, err := os.ReadFile("koruri/Koruri-Bold.ttf")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	ft, err := truetype.Parse(ftBinary)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	opt := truetype.Options{
		Size:              90,
		DPI:               0,
		Hinting:           0,
		GlyphCacheEntries: 0,
		SubPixelsX:        0,
		SubPixelsY:        0,
	}

	imageWidth := 100
	imageHeight := 100
	textTopMargin := 90
	text := "れいぞうこ"

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

	buf := &bytes.Buffer{}
	err = png.Encode(buf, img)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	file, err := os.Create(`src/text.png`)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	file.Write(buf.Bytes())
}
