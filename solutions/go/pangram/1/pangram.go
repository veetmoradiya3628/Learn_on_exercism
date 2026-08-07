package pangram
import (
	"strings"
)
func IsPangram(input string) bool {
	input = strings.ToLower(input)
    seen := make(map[rune]bool)
    count := 0
    for _, r := range input {
        if r >= 'a' && r <= 'z' {
            if !seen[r] {
                seen[r] = true
                count++
            }
        }
    }
    return count == 26
}
