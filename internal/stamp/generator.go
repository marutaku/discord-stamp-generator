/*
Generator of images which only in text for stamps using in Discord and Slack.
*/

package stamp

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"strings"

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

func splitTextIntoLines(text string, fontWidth int, maxWidth int) []string {
	words := strings.Fields(text)
	var lines []string
	var currentLine string
	for _, word := range words {
		fmt.Println((len(currentLine) + len(word) + 1))
		if (len(currentLine) + len(word) + 1) > maxWidth {
			fmt.Println("Append new line")
			lines = append(lines, currentLine)
			currentLine = word
		} else {
			if currentLine != "" {
				currentLine += " "
			}
			currentLine += word
		}
	}
	if currentLine != "" {
		lines = append(lines, currentLine)
	}
	return lines
}

func (g *Generator) Generate(text string) ([]byte, error) {
	// Initialize image
	rect := image.Rect(0, 0, g.Width, g.Height)
	img := image.NewRGBA(rect)
	drawer := &font.Drawer{
		Dst:  img,
		Src:  g.FontColor,
		Face: g.Font,
	}
	imageCenterY := g.Height / 2
	fmt.Printf("imageCenterY: %d\n", imageCenterY)
	lines := splitTextIntoLines(text, g.Font.Metrics().Height.Ceil(), g.Width)
	totalTextHeight := g.Font.Metrics().Height.Ceil() * len(lines)
	fmt.Printf("totalTextHeight: %d\n", totalTextHeight)
	// FIXME: I don't know why I need to divide by 4
	startY := imageCenterY + totalTextHeight/4
	fmt.Printf("startY: %d\n", startY)
	for i, line := range lines {
		fmt.Printf("line: %s\n", line)
		y := startY + g.Font.Metrics().Height.Ceil()*i
		fmt.Printf("y: %d\n", y)
		textWidth := drawer.MeasureString(line).Round()
		fmt.Printf("textWidth: %d\n", textWidth)
		x := (g.Width - textWidth) / 2
		drawer.Dot = fixed.Point26_6{X: fixed.I(x), Y: fixed.I(y)}
		drawer.DrawString(line)
	}

	buffer := new(bytes.Buffer)
	png.Encode(buffer, img)
	return buffer.Bytes(), nil
}
