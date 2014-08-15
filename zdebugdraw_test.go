package cp

import (
	"testing"
)

func TestDebugDrawNil(t *testing.T) {
	space := SpaceNew()

	body := BodyNew(10, MomentForCircle(1, 0, 32, V(0, 0)))
	space.AddBody(body)
	space.AddShape(body.CircleShapeNew(32, V(0, 0)))

	// Fill in debugDraw parameters as desired, in this case we don't fill in
	// any.
	debugDraw := new(SpaceDebugDrawOptions)
	debugDraw.Flags = SPACE_DEBUG_DRAW_SHAPES | SPACE_DEBUG_DRAW_CONSTRAINTS | SPACE_DEBUG_DRAW_COLLISION_POINTS
	space.DebugDraw(debugDraw)
}

func TestDebugDraw(t *testing.T) {
	space := SpaceNew()

	body := BodyNew(10, MomentForCircle(1, 0, 32, V(0, 0)))
	space.AddBody(body)
	space.AddShape(body.CircleShapeNew(32, V(0, 0)))

	debugDraw := new(SpaceDebugDrawOptions)
	debugDraw.Flags = SPACE_DEBUG_DRAW_SHAPES | SPACE_DEBUG_DRAW_CONSTRAINTS | SPACE_DEBUG_DRAW_COLLISION_POINTS

	// Simply fed into each callback below. No need to use it, could just use
	// closures like we do below and access data directly instead of using the
	// data parameter.
	debugDraw.Data = "Hello World!"

	debugDraw.DrawCircle = func(pos Vect, angle, radius float64,
		outlineColor, fillColor SpaceDebugColor, data interface{}) {
		t.Logf("DrawCircle(pos=%v, angle=%v, radius=%v, outlineColor=%v, fillColor=%v, data=%v)\n", pos, angle, radius, outlineColor, fillColor, data)
	}
	debugDraw.DrawSegment = func(a, b Vect, color SpaceDebugColor,
		data interface{}) {
		t.Logf("DrawSegment(a=%v, b=%v, color=%v, data=%v)\n", a, b, color, data)
	}
	debugDraw.DrawFatSegment = func(a, b Vect, radius float64,
		outlineColor, fillColor SpaceDebugColor, data interface{}) {
		t.Logf("DrawFatSegment(a=%v, b=%v, radius=%v, outlineColor=%v, fillColor=%v, data=%v)\n", a, b, radius, outlineColor, fillColor, data)
	}
	debugDraw.DrawPolygon = func(verts []Vect, radius float64,
		outlineColor, fillColor SpaceDebugColor, data interface{}) {
		t.Logf("DrawPolygon(verts[:2]=%v, radius=%v, outlineColor=%v, fillColor=%v, data=%v)\n", verts, radius, outlineColor, fillColor, data)
	}
	debugDraw.DrawDot = func(size float64, pos Vect,
		color SpaceDebugColor, data interface{}) {
		t.Logf("DrawDot(size=%v, pos=%v, color=%v, data=%v)\n", size, pos, color, data)
	}
	debugDraw.ColorForShape = func(shape *Shape, data interface{}) SpaceDebugColor {
		t.Logf("ColorForShape(shape=%v, data=%v)\n", shape, data)
		return SpaceDebugColor{}
	}

	space.DebugDraw(debugDraw)
}
