package main

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	"log"
	"math"
	"strings"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
)

var font *truetype.Font

func init() {
	var err error
	font, err = truetype.Parse(MustAsset("文泉驿等宽正黑.ttf"))
	if err != nil {
		log.Fatal(err)
	}

}

func WaterMark(img []byte, text string, size float64, scale int) ([]byte, error) {
	if scale <= 1 {
		return nil, errors.New("scale must gte 1")
	}
	im, _, err := image.Decode(bytes.NewReader(img))
	if err != nil {
		return nil, err
	}
	w, h := im.Bounds().Max.X, im.Bounds().Max.Y
	diagonal := int(math.RoundToEven(math.Sqrt(float64(w*w+h*h)))) * 2
	waterMarkDc := gg.NewContext(diagonal, diagonal)
	waterMarkDc.SetFontFace(truetype.NewFace(font, &truetype.Options{Size: size}))
	waterMarkDc.Clear()
	waterMarkDc.SetRGB(105, 105, 105)
	var texts []string
	tw, th := waterMarkDc.MeasureString(text)
	for i := 0; i < diagonal/(scale*int(tw)); i++ {
		texts = append(texts, text)
	}
	text = strings.Join(texts, strings.Repeat(" ", len(text)))
	for i := 0; i < diagonal; i += scale * int(th) {
		waterMarkDc.DrawStringAnchored(text, float64(diagonal/2), float64(i), 0.5, 0.5)
	}

	dc := gg.NewContext(w, h)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.DrawImage(im, 0, 0)
	dc.RotateAbout(gg.Radians(45), float64(w/2), float64(h/2))
	dc.DrawImageAnchored(waterMarkDc.Image(), w/2, h/2, 0.5, 0.5)

	waterMarkedImg := new(bytes.Buffer)
	err = jpeg.Encode(waterMarkedImg, dc.Image(), nil)
	if err != nil {
		return nil, err
	}
	return waterMarkedImg.Bytes(), nil
}
