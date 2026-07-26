package wordcount
import (
	"regexp"
    "strings"
)
type Frequency map[string]int

func WordCount(phrase string) Frequency {
	re := regexp.MustCompile(`[a-zA-Z0-9]+('[a-zA-Z0-9]+)?`)
    words := re.FindAllString(phrase, -1)

    ans := make(Frequency)
    for _, val := range words {
        lowerCase := strings.ToLower(val)
        ans[lowerCase]++
    }
    return ans
}
