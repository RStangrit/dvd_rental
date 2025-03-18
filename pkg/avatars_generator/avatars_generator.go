package avatars_generator

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"strings"

	"github.com/golang/freetype"
	"golang.org/x/image/font"
)

const (
	imageSize = 50
	fontSize  = 20
	fontFile  = "assets/fonts/nimbussanl_boldcond.ttf"
)

func CreateAvatar(initials string) ([]byte, error) {

	img := image.NewRGBA(image.Rect(0, 0, imageSize, imageSize))

	bgColor := color.RGBA{0, 122, 255, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{bgColor}, image.Point{}, draw.Src)

	fontBytes, err := os.ReadFile(fontFile)
	if err != nil {
		return nil, fmt.Errorf("font loading error: %v", err)
	}
	fontParsed, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return nil, fmt.Errorf("font parsing error: %v", err)
	}

	context := freetype.NewContext()
	context.SetDPI(72)
	context.SetFont(fontParsed)
	context.SetFontSize(fontSize)
	context.SetClip(img.Bounds())
	context.SetDst(img)
	context.SetSrc(image.White)
	context.SetHinting(font.HintingFull)

	pt := freetype.Pt(imageSize/4, imageSize/2+fontSize/3)
	_, err = context.DrawString(strings.ToUpper(initials), pt)
	if err != nil {
		return nil, fmt.Errorf("text rendering error: %v", err)
	}

	var buf bytes.Buffer
	err = png.Encode(&buf, img)
	if err != nil {
		return nil, fmt.Errorf("PNG encoding error: %v", err)
	}

	return buf.Bytes(), nil
}
