package etl

import (
    "strings"
)

func Transform(in map[int][]string) map[string]int {
	ans := make(map[string]int)

    for k, v := range in {
        for _, item := range v {
            lower_item := strings.ToLower(item)
            ans[lower_item] = k
        }
    }
	return ans    
}
