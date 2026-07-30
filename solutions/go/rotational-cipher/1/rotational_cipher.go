package rotationalcipher

import "strings"

func RotationalCipher(plain string, shiftKey int) string {
	var result strings.Builder
	result.Grow(len(plain))

	shiftKey = shiftKey % 26

	for _, r := range plain {
		if r >= 'a' && r <= 'z' {
			cipherRune := 'a' + (r - 'a' + rune(shiftKey)) % 26
			result.WriteRune(cipherRune)
		} else if r >= 'A' && r <= 'Z' {
			cipherRune := 'A' + (r - 'A' + rune(shiftKey)) % 26
			result.WriteRune(cipherRune)
		} else {
			result.WriteRune(r)
		}
	}

	return result.String()
}