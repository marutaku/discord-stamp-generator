package stamp

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"image/color"
)

// GenerateRandomHexColor generates a random hex color.
func GenerateRandomHexColor() (string, error) {
	colorBytes := make([]byte, 3)
	_, err := rand.Read(colorBytes)
	if err != nil {
		return "", err
	}
	return "#" + hex.EncodeToString(colorBytes), nil
}

// ParseHexColor parses a hex color string to a color.RGBA.
func ParseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff
	if s[0] != '#' {
		return c, fmt.Errorf("invalid format, must start with #")
	}
	hexColor := s[1:]
	switch len(hexColor) {
	case 6:
		_, err = fmt.Sscanf(hexColor, "%02x%02x%02x", &c.R, &c.G, &c.B)
	case 3:
		_, err = fmt.Sscanf(hexColor, "%1x%1x%1x", &c.R, &c.G, &c.B)
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")
	}
	return
}
