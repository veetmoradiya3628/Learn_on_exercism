package armstrongnumbers

import (
    "strconv"
    "math"
)

func IsNumber(n int) bool {
	pow_len := len(strconv.Itoa(n))
	original_num := n
    ans := 0
    for n != 0 {
        ans += int(math.Pow(float64(n % 10), float64(pow_len)))
        n = n / 10
    }
    return ans == original_num
}
