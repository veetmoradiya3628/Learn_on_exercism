package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	ans := initial
    for _, v := range s {
        ans = fn(ans, v)
    }
    return ans
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	ans := initial
    for i := len(s) - 1; i >= 0; i-- {
        ans = fn(s[i], ans)
    }
    return ans
}

func (s IntList) Filter(fn func(int) bool) IntList {
	var result IntList
    for _, v := range s {
        if fn(v) {
            result = append(result, v)
        }
    }
    return result
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	result := make(IntList, len(s))
	for i, v := range s {
		result[i] = fn(v)
	}
	return result
}

func (s IntList) Reverse() IntList {
	result := make(IntList, len(s))
	for i, v := range s {
		result[len(s)-1-i] = v
	}
	return result
}

func (s IntList) Append(lst IntList) IntList {
	result := make(IntList, 0, len(s)+len(lst))
	result = append(result, s...)
	result = append(result, lst...)
	return result
}

func (s IntList) Concat(lists []IntList) IntList {
	totalLen := len(s)
	for _, l := range lists {
		totalLen += len(l)
	}

	result := make(IntList, 0, totalLen)
	result = append(result, s...)
	for _, l := range lists {
		result = append(result, l...)
	}
	return result
}
