package stamp

import (
	_ "embed"
	"fmt"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed NotoSansJP-ExtraBold.ttf
var notoSansJPFontFile []byte

// CalculateFontSize calculates the font size based on the maximum number of characters in one line and the width.
func CalculateFontSize(maxChatNumInOneLine int, width int) float64 {
	return float64(width) / float64(maxChatNumInOneLine)
}

// LoadFont loads the font based on the font size and the external font file path.
func LoadFont(fontSize float64, externalFontFilePath string) (font.Face, error) {
	fontFile := notoSansJPFontFile
	if externalFontFilePath != "" {
		var err error
		fontFile, err = loadExternalFontFile(externalFontFilePath)
		if err != nil {
			return nil, err
		}
	}
	tt, err := opentype.Parse(fontFile)
	if err != nil {
		return nil, err
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return nil, err
	}

	return face, nil
}

// loadExternalFontFile loads the external font file.
func loadExternalFontFile(fontFilePath string) ([]byte, error) {
	if fontFilePath == "" {
		return nil, nil
	}
	fontFile, err := os.ReadFile(fontFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read font file: %w", err)
	}
	return fontFile, nil
}
