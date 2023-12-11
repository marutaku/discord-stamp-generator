package stamp

import (
	"crypto/rand"
	"fmt"
	"image/color"
	"math/big"
)

func GenerateRandomHexColor() (string, error) {
	red, err := rand.Int(rand.Reader, big.NewInt(255))
	if err != nil {
		return "", err
	}
	green, err := rand.Int(rand.Reader, big.NewInt(255))
	if err != nil {
		return "", err
	}
	blue, err := rand.Int(rand.Reader, big.NewInt(255))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("#%02x%02x%02x", red, green, blue), nil
}

// https://stackoverflow.com/questions/54197913/parse-hex-string-to-image-color
func ParseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")

	}
	return
}
