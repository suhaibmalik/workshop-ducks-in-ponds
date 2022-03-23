package pond

import (
	"ducks-in-ponds/duck"
	"ducks-in-ponds/point"
)

type Pond struct {
	Name   string     `json:"name"`
	Radius float64    `json:"radius"`
	Center point.Vec3 `json:"center"`
}

// todo: Create rule(s) for a duck being inside a pond
//   Are ponds spherical?
//   Does your pond have a miniscus?
//   Are ponds even real?
func (p Pond) IsInside(d duck.Duck) bool {
	return false
}
