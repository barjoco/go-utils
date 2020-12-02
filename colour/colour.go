package colour

import (
	"errors"
)

var errInvalidFormat = errors.New("invalid format")

// HexToRGB ...
func HexToRGB(s string) (r, g, b int) {
	if s[0] != '#' {
		panic(errInvalidFormat)
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		panic(errInvalidFormat)

	}

	switch len(s) {
	case 7:
		r = int(hexToByte(s[1])<<4 + hexToByte(s[2]))
		g = int(hexToByte(s[3])<<4 + hexToByte(s[4]))
		b = int(hexToByte(s[5])<<4 + hexToByte(s[6]))
	case 4:
		r = int(hexToByte(s[1]) * 17)
		g = int(hexToByte(s[2]) * 17)
		b = int(hexToByte(s[3]) * 17)
	default:
		panic(errInvalidFormat)
	}

	return
}
