package scrabblescore
import (
    "unicode"
)
func Score(word string) int {
	ans := 0
	for idx := 0; idx < len(word); idx++ {
        c := byte(unicode.ToUpper(rune(word[idx])))
        switch c {
            case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
            	ans += 1
            	break
        	case 'D', 'G':
            	ans += 2
            	break
            case 'B', 'C', 'M', 'P':
            	ans += 3
            	break
            case 'F', 'H', 'V', 'W', 'Y':
            	ans += 4
            	break
            case 'K':
            	ans += 5
            	break
            case 'J', 'X':
            	ans += 8
            	break
            case 'Q', 'Z':
            	ans += 10
            	break
            default:
            	break
        }
    }
    return ans
}
