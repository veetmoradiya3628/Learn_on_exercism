package raindrops
import (
    "strings"
    "strconv"
)
func Convert(number int) string {
	var ans strings.Builder;
    if number % 3 == 0 {
        ans.WriteString("Pling")
    }
    if number % 5 == 0 {
        ans.WriteString("Plang")
    }
    if number % 7 == 0 {
        ans.WriteString("Plong")
    }

	if ans.Len() == 0 {
        return strconv.Itoa(number)
    }
    
	return ans.String()
}
