package solution

import (
	"math"
)

type Num int

const SidesTriangle = 3
const SidesSquare = 4
const SidesCircle = 0

func CalcSquare(sideLen float64, sidesNum Num) float64 {
	switch sidesNum {
	case SidesCircle:
		return math.Pi * math.Pow(sideLen, 2)
	case SidesTriangle:
		return math.Sqrt(SidesTriangle) / 4 * math.Pow(sideLen, 2)
	case SidesSquare:
		return math.Pow(sideLen, 2)
	default:
		return 0
	}
}
