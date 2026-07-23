package collatzconjecture
import (
    "math"
    "errors"
    )
func CollatzConjecture(n int) (int, error) {
    if n <= 0 {
        return 0, errors.New("Negative number detected")
    }
	ans := 0
    for n != 1 {
        if n % 2 == 0 {
            n = n / 2
        } else {
            if n > (math.MaxInt-1)/3 {
				return 0, errors.New("integer overflow occurred during calculation")
			}
            n = 3 * n + 1
        }
        ans++
    }
    return ans, nil
}
