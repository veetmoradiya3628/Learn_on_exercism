package darts

import "math"

func Score(x, y float64) int {
	dartLocation := math.Sqrt(x*x + y*y)
    switch {
    case dartLocation <= 1.0:
        return 10
    case dartLocation <= 5.0:
    	return 5
    case dartLocation <= 10.0:
    	return 1
    default:
    	return 0
    }
}
