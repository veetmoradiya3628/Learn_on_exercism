package hamming
import "fmt"
func Distance(a, b string) (int, error) {
    if len(a) != len(b){
        return 0, fmt.Errorf("different length")
    }
	ans := 0;
    for idx := 0; idx < len(a); idx++ {
        if a[idx] != b[idx] {
            ans++
        }
    }
    return ans, nil
}
