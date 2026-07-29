// Package bob provides a simulation of a lackadaisical teenager's responses.
package bob

import "strings"

// Hey returns Bob's response to various types of remarks.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}

	isQuestion := strings.HasSuffix(remark, "?")
	isShouting := hasLetters(remark) && strings.ToUpper(remark) == remark

	if isShouting && isQuestion {
		return "Calm down, I know what I'm doing!"
	}

	if isShouting {
		return "Whoa, chill out!"
	}

	if isQuestion {
		return "Sure."
	}

	return "Whatever."
}

func hasLetters(s string) bool {
	return strings.ContainsAny(strings.ToLower(s), "abcdefghijklmnopqrstuvwxyz")
}
