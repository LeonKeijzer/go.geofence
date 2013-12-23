package geofence

type Radius interface {
	Inside(p Point) bool
}
