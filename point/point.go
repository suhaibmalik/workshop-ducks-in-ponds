package point

import "math"

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

// Distance2D calculates the 2D (X & Y) distance between two vector points
//  using the equation: sqrt((x_1 - x_2)^2 + (y_1 - y_2)^2)
func (v1 Vec3) Distance2D(v2 Vec3) float64 {
	return math.Sqrt(math.Pow(v1.X-v2.X, 2) + math.Pow(2, v1.Y-v2.Y))
}

// todo: A better 3D distance calculation
// func (v1 Vec3) Distance(v2 Vec3) float64 {
// }
