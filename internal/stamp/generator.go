/*
Generator of images which only in text for stamps using in Discord and Slack.
*/

package stamp

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"math"
	"strings"
	"unicode/utf8"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

var MARGIN_PER_LINE = 24
var MIN_CHAR_NUM_IN_LINE = 3

func splitTextIntoLines(text string) []string {
	words := strings.Split(text, "")
	lengthOfText := len(words)
	if lengthOfText <= MIN_CHAR_NUM_IN_LINE {
		return []string{text}
	}
	middleOfWords := int(math.Ceil(float64(lengthOfText) / 2))
	return []string{
		strings.Join(words[:middleOfWords], ""),
		strings.Join(words[middleOfWords:], ""),
	}
}

func drawLines(drawer *font.Drawer, lines []string, startY int, lineHeight int, imgWith int) {
	for i, line := range lines {
		y := startY + lineHeight*i + MARGIN_PER_LINE*i
		fmt.Printf("y: %d\n", y)
		textWidth := drawer.MeasureString(line).Round()
		x := (imgWith - textWidth) / 2
		drawer.Dot = fixed.Point26_6{X: fixed.I(x), Y: fixed.I(y)}
		drawer.DrawString(line)
	}
}

func decideStartY(imgHeight int, fontHeight int, lineNum int) int {
	imageCenterY := imgHeight / 2
	totalTextHeight := fontHeight * lineNum
	margin := (MARGIN_PER_LINE / 2) * (lineNum - 1)
	fmt.Printf("imageCenterY: %d, totalTextHeight: %d, margin: %d\n", imageCenterY, totalTextHeight, margin)
	return imageCenterY - totalTextHeight/2 + fontHeight - margin
}

func Generate(text string, width int, height int, fontColor string, externalFontPath string) ([]byte, error) {
	lines := splitTextIntoLines(text)
	fontColorRGBA, err := ParseHexColor(fontColor)
	if err != nil {
		return nil, err
	}
	fontSize := CalculateFontSize(utf8.RuneCountInString(lines[0]), width)
	fontFace, err := LoadFont(fontSize, externalFontPath)
	if err != nil {
		return nil, err
	}
	// Initialize image
	rect := image.Rect(0, 0, width, height)
	img := image.NewRGBA(rect)
	drawer := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(fontColorRGBA),
		Face: fontFace,
	}
	fontHeight := fontFace.Metrics().CapHeight.Ceil()
	startY := decideStartY(height, fontHeight, len(lines))
	drawLines(drawer, lines, startY, fontHeight, width)
	buffer := new(bytes.Buffer)
	png.Encode(buffer, img)
	return buffer.Bytes(), nil
}
