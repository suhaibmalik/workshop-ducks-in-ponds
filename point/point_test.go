package point

import (
	"math"
	"testing"
)

func TestDistance2D(t *testing.T) {
	cases := []struct {
		v1       Vec3
		v2       Vec3
		expected float64
	}{
		{Vec3{0, 0, 0}, Vec3{0, 0, 0}, 0},
		{Vec3{0, 0, 0}, Vec3{1, 1, 1}, math.Sqrt(2)},
		{Vec3{0, 0, 0}, Vec3{1, 1, 0}, math.Sqrt(2)},
		{Vec3{2, 0, 0}, Vec3{1, 0, 0}, 1},
	}

	for _, c := range cases {
		actual := c.v1.Distance2D(c.v2)
		if actual != c.expected {
			t.Fatalf("expected distance of: %v got: %v for a: %v and b %v",
				c.expected, actual, c.v1, c.v2)
		}
	}

}
