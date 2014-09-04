// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cp

/*
#include "chipmunk/include/chipmunk/chipmunk.h"
*/
import "C"

import (
	"runtime"
	"unsafe"
)

// Done by hand for their nice commentary:

// Nearest point query info struct.
type PointQueryInfo struct {
	// The nearest shape, nil if no shape was within range.
	Shape *Shape

	// The closest point on the shape's surface. (in world space coordinates)
	Point Vect

	// The distance to the point. The distance is negative if the point is
	// inside the shape.
	Distance float64

	// The gradient of the signed distance function.
	//
	// The same as info.p/info.d, but accurate even for very small values of
	// info.d.
	Gradient Vect
}

// Segment query info struct.
type SegmentQueryInfo struct {
	// The shape that was hit, nil if no collision occured.
	Shape *Shape

	// The point of impact.
	Point Vect

	// The normal of the surface hit.
	Normal Vect

	// The normalized distance along the query segment in the range [0, 1].
	Alpha float64
}

type ShapeFilter struct {
	Group      Group
	Categories Bitmask
	Mask       Bitmask
}

var (
	SHAPE_FILTER_ALL  = ShapeFilter{NO_GROUP, ALL_CATEGORIES, ALL_CATEGORIES}
	SHAPE_FILTER_NONE = ShapeFilter{NO_GROUP, 0 &^ ALL_CATEGORIES, 0 &^ ALL_CATEGORIES}
)

func ShapeFilterNew(group Group, categories, mask Bitmask) ShapeFilter {
	return ShapeFilter{group, categories, mask}
}

// The Shape struct defines the shape of a rigid body.
type Shape struct {
	c        *C.cpShape
	spaceRef *Space
	userData interface{}
}

func goShape(c *C.cpShape, optional *Space) *Shape {
	data := C.cpShapeGetUserData(c)
	g := (*Shape)(data)
	g.spaceRef = optional
	return g
}

// Allocate and initialize a circle shape.
func (b *Body) CircleShapeNew(radius float64, offset Vect) *Shape {
	s := new(Shape)
	s.c = C.cpCircleShapeNew(b.c, C.cpFloat(radius), offset.c())
	if s.c == nil {
		return nil
	}
	C.cpShapeSetUserData(s.c, C.cpDataPointer(unsafe.Pointer(s)))
	runtime.SetFinalizer(s, finalizeShape)
	return s
}

// Allocate and initialize a segment shape.
func (bd *Body) SegmentShapeNew(a, b Vect, radius float64) *Shape {
	s := new(Shape)
	s.c = C.cpSegmentShapeNew(bd.c, a.c(), b.c(), C.cpFloat(radius))
	if s.c == nil {
		return nil
	}
	C.cpShapeSetUserData(s.c, C.cpDataPointer(unsafe.Pointer(s)))
	runtime.SetFinalizer(s, finalizeShape)
	return s
}

func finalizeShape(s *Shape) {
	if s.c != nil {
		s.c = nil
		C.cpShapeFree(s.c)
	}
}

// Free is deprecated. Do not use it, it is no-op.
func (s *Shape) Free() {
}

// Update, cache and return the bounding box of a shape based on the body it's
// attached to.
func (s *Shape) CacheBB() BB {
	return goBB(C.cpShapeCacheBB(s.c))
}

// Update, cache and return the bounding box of a shape with an explicit transformation.
func (s *Shape) Update(transform Transform) BB {
	return goBB(C.cpShapeUpdate(s.c, transform.c()))
}

// Perform a nearest point query. It finds the closest point on the surface of
// shape to a specific point.
//
// The value returned is the distance between the points. A negative distance
// means the point is inside the shape.
func (s *Shape) PointQuery(p Vect) (out *PointQueryInfo, d float64) {
	out = new(PointQueryInfo)
	d = float64(C.cpShapePointQuery(
		s.c,
		p.c(),
		(*C.cpPointQueryInfo)(unsafe.Pointer(&out)),
	))
	return
}

// Perform a segment query against a shape.
func (s *Shape) SegmentQuery(a, b Vect, radius float64) (info *SegmentQueryInfo, ret bool) {
	info = new(SegmentQueryInfo)
	ret = goBool(C.cpShapeSegmentQuery(
		s.c,
		a.c(),
		b.c(),
		C.cpFloat(radius),
		(*C.cpSegmentQueryInfo)(unsafe.Pointer(&info)),
	))
	return
}

// Return contact information about two shapes.
func (a *Shape) ShapesCollide(b *Shape) ContactPointSet {
	ret := C.cpShapesCollide(a.c, b.c)
	return *(*ContactPointSet)(unsafe.Pointer(&ret))
}

// The cpSpace this body is added to.
func (s *Shape) Space() *Space {
	return goSpace(C.cpShapeGetSpace(s.c))
}

// The cpBody this shape is connected to.
func (s *Shape) Body() *Body {
	return goBody(C.cpShapeGetBody(s.c), s.Space())
}

// Set the cpBody this shape is connected to.
//
// Can only be used if the shape is not currently added to a space.
func (s *Shape) SetBody(b *Body) {
	s.bodyRef = b
	C.cpShapeSetBody(s.c, b.c)
}

// Get the mass of the shape if you are having Chipmunk calculate mass properties for you.
func (s *Shape) Mass() float64 {
	return float64(C.cpShapeGetMass(s.c))
}

// Set the mass of this shape to have Chipmunk calculate mass properties for you.
func (s *Shape) SetMass(mass float64) {
	C.cpShapeSetMass(
		s.c,
		C.cpFloat(mass),
	)
}

// Get the density of the shape if you are having Chipmunk calculate mass properties for you.
func (s *Shape) Density() float64 {
	return float64(C.cpShapeGetDensity(s.c))
}

// Set the density  of this shape to have Chipmunk calculate mass properties for you.
func (s *Shape) SetDensity(density float64) {
	C.cpShapeSetDensity(
		s.c,
		C.cpFloat(density),
	)
}

// Get the calculated moment of inertia for this shape.
func (s *Shape) Moment() float64 {
	return float64(C.cpShapeGetMoment(s.c))
}

// Get the calculated area of this shape.
func (s *Shape) Area() float64 {
	return float64(C.cpShapeGetArea(s.c))
}

// Get the centroid of this shape.
func (s *Shape) CenterOfGravity() Vect {
	return goVect(C.cpShapeGetCenterOfGravity(s.c))
}

// Get the bounding box that contains the shape given it's current position and angle.
func (s *Shape) BB() BB {
	return goBB(C.cpShapeGetBB(s.c))
}

// Get if the shape is set to be a sensor or not.
func (s *Shape) Sensor() bool {
	return goBool(C.cpShapeGetSensor(s.c))
}

// Set if the shape is a sensor or not.
func (s *Shape) SetSensor(sensor bool) {
	var cbool C.cpBool = C.cpTrue
	if !sensor {
		cbool = C.cpFalse
	}
	C.cpShapeSetSensor(s.c, cbool)
}

// Get the elasticity of this shape.
func (s *Shape) Elasticity() float64 {
	return float64(C.cpShapeGetElasticity(s.c))
}

// Set the elasticity of this shape.
func (s *Shape) SetElasticity(elasticity float64) {
	C.cpShapeSetElasticity(
		s.c,
		C.cpFloat(elasticity),
	)
}

// Get the friction of this shape.
func (s *Shape) Friction() float64 {
	return float64(C.cpShapeGetFriction(s.c))
}

// Set the friction of this shape.
func (s *Shape) SetFriction(friction float64) {
	C.cpShapeSetFriction(
		s.c,
		C.cpFloat(friction),
	)
}

// Get the surface velocity of this shape.
func (s *Shape) SurfaceVelocity() Vect {
	return goVect(C.cpShapeGetSurfaceVelocity(s.c))
}

// Set the surface velocity of this shape.
func (s *Shape) SetSurfaceVelocity(surfaceVelocity Vect) {
	C.cpShapeSetSurfaceVelocity(
		s.c,
		surfaceVelocity.c(),
	)
	return
}

// Get the user definable data interface of this shape.
func (s *Shape) UserData() interface{} {
	return s.userData
}

// Set the user definable data pointer of this shape.
func (s *Shape) SetUserData(i interface{}) {
	s.userData = i
}

// Get the collision type of this shape.
func (s *Shape) CollisionType() CollisionType {
	return CollisionType(C.cpShapeGetCollisionType(s.c))
}

// Set the collision type of this shape.
func (s *Shape) SetCollisionType(collisionType CollisionType) {
	C.cpShapeSetCollisionType(
		s.c,
		C.cpCollisionType(collisionType),
	)
}

// Get the collision filtering parameters of this shape.
func (s *Shape) Filter() ShapeFilter {
	ret := C.cpShapeGetFilter(s.c)
	return *(*ShapeFilter)(unsafe.Pointer(&ret))
}

// Set the collision filtering parameters of this shape.
func (s *Shape) SetFilter(filter ShapeFilter) {
	C.cpShapeSetFilter(
		s.c,
		*(*C.cpShapeFilter)(unsafe.Pointer(&filter)),
	)
}

// Get the offset of a circle shape.
func (s *Shape) CircleOffset() Vect {
	return goVect(C.cpCircleShapeGetOffset(s.c))
}

// Get the radius of a circle shape.
func (s *Shape) CircleRadius() float64 {
	return float64(C.cpCircleShapeGetRadius(s.c))
}

// Let Chipmunk know about the geometry of adjacent segments to avoid colliding with endcaps.
func (s *Shape) SegmentSetNeighbors(prev, next Vect) {
	C.cpSegmentShapeSetNeighbors(
		s.c,
		prev.c(),
		next.c(),
	)
}

// Get the first endpoint of a segment shape.
func (s *Shape) SegmentA() Vect {
	return goVect(C.cpSegmentShapeGetA(s.c))
}

// Get the second endpoint of a segment shape.
func (s *Shape) SegmentB() Vect {
	return goVect(C.cpSegmentShapeGetB(s.c))
}

// Get the normal of a segment shape.
func (s *Shape) SegmentNormal() Vect {
	return goVect(C.cpSegmentShapeGetNormal(s.c))
}

// Get the first endpoint of a segment shape.
func (s *Shape) SegmentRadius() float64 {
	return float64(C.cpSegmentShapeGetRadius(s.c))
}
