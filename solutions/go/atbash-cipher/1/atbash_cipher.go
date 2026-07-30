package atbashcipher

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	var result strings.Builder
	s = strings.ToLower(s)
	count := 0
	for _, r := range s {
		var cipherRune rune
		if r >= 'a' && r <= 'z' {
			cipherRune = 'z' - (r - 'a')
		} else if unicode.IsDigit(r) {
			cipherRune = r
		} else {
			continue
		}
		if count > 0 && count%5 == 0 {
			result.WriteRune(' ')
		}
		result.WriteRune(cipherRune)
		count++
	}
	return result.String()
}