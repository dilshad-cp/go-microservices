package geometry

import "math"

func Area(len, wid float64) float64 {
	return len * wid
}

func Diagonal(len, wid float64) float64 {
	return math.Sqrt((len * len) + (wid * wid))
}
