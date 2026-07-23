package luhn
import (
    "strings"
)
func Valid(id string) bool {
	strId := strings.ReplaceAll(id, " ", "")
    if len(strId) <= 1 {
		return false
	}
	sum := 0
    should_double := false
	for i := len(strId) - 1; i >= 0; i-- {
        if (strId[i] < '0' || strId[i] > '9'){
            return false;
        }
        digit := int(strId[i] - '0')
        if should_double {
            digit *= 2
            if digit > 9 {
                digit -= 9
            }
        }
        sum += digit
        should_double = !should_double
    }
    return sum % 10 == 0
}
