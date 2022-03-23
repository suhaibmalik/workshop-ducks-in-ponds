package duck

import "ducks-in-ponds/point"

type Duck struct {
	Name     string     `json:"name"`
	Position point.Vec3 `json:"position"`
}
