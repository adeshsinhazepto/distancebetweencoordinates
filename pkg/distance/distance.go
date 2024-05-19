package distance

import (
	"math"
)

func Distance(p1, p2 Point) float64 {
	return math.Sqrt(float64((p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y)))
}
