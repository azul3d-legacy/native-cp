package cp

import (
	"testing"
)

func TestVelocityUpdateFunc(t *testing.T) {
	space := SpaceNew()
	shape := BodyNew(10, MomentForCircle(1, 0, 32, V(0, 0)))
	space.AddBody(shape)
	shape.SetVelocityUpdateFunc(func(body *Body, gravity Vect, damping, dt float64) {
	})
	space.Step(1)

	space.RemoveBody(shape)
	shape.Free()
	space.Free()
}

func TestPositionUpdateFunc(t *testing.T) {
	space := SpaceNew()
	shape := BodyNew(10, MomentForCircle(1, 0, 32, V(0, 0)))
	space.AddBody(shape)
	shape.SetPositionUpdateFunc(func(body *Body, dt float64) {
	})
	space.Step(1)

	space.RemoveBody(shape)
	shape.Free()
	space.Free()
}
