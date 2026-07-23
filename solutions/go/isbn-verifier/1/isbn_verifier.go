package isbnverifier
import (
    "strings"
)
func IsValidISBN(isbn string) bool {
    strId := strings.ReplaceAll(isbn, "-", "")
    if len(strId) != 10 {
        return false
    }
    sum := 0
    for i := 0; i < 10; i++ {
        char := strId[i]
		weight := 10 - i

		if i == 9 && (char == 'X' || char == 'x') {
			sum += 10 * weight
		} else if char >= '0' && char <= '9' {
			digit := int(char - '0')
			sum += digit * weight
		} else {
			return false
		}
    }
    return sum % 11 == 0
}
