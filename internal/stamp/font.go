package stamp

import (
	_ "embed"
	"fmt"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// go:embed NotoSansJP-Regular.ttf
var notoSansJPFontFile []byte

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
