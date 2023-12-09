/*
Generator of images which only in text for stamps using in Discord and Slack.
*/

package stamp

import (
	"bytes"
	"image"
	"image/png"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type Generator struct {
	Font      font.Face
	FontColor *image.Uniform
	Width     int
	Height    int
}

func NewGenerator(fontSize float64, width int, height int, fontColor string, externalFontPath string) (*Generator, error) {
	font, err := LoadFont(fontSize, externalFontPath)
	if err != nil {
		return nil, err
	}
	fontColorRGBA, err := ParseHexColor(fontColor)
	if err != nil {
		return nil, err
	}
	return &Generator{
		Font:      font,
		Width:     width,
		Height:    height,
		FontColor: image.NewUniform(fontColorRGBA),
	}, nil
}

func (g *Generator) Generate(text string) ([]byte, error) {
	a := font.MeasureString(g.Font, text)
	height := g.Font.Metrics().Height

	rect := image.Rect(0, 0, a.Ceil(), height.Ceil())
	img := image.NewRGBA(rect)

	y := g.Font.Metrics().Ascent

	drawer := &font.Drawer{
		Dst:  img,
		Src:  g.FontColor,
		Face: g.Font,
		Dot:  fixed.Point26_6{Y: y},
	}
	drawer.DrawString(text)
	buffer := new(bytes.Buffer)
	png.Encode(buffer, img)

	return buffer.Bytes(), nil
}
