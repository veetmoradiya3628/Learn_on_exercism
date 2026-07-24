package anagram

import (
	"slices"
	"strings"
)

func sortString(s string) string {
	runes := []rune(strings.ToLower(s))
	slices.Sort(runes)
	return string(runes)
}

func Detect(subject string, candidates []string) []string {
	sortedSubject := sortString(subject)
	ans := []string{}

	for _, candidate := range candidates {
		if strings.EqualFold(subject, candidate) {
			continue
		}

		if sortedSubject == sortString(candidate) {
			ans = append(ans, candidate)
		}
	}

	return ans
}