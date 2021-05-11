package utils

import "math"

func RoundFloor(n float64, numPlaces int) float64 {
	return math.Floor(n*math.Pow10(numPlaces)) / math.Pow10(numPlaces)
}
