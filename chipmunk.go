// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cp is a wrapper for the Chipmunk 2D Physics Engine.
//
// More information about Chipmunk can be found at:
//  http://chipmunk-physics.net/
//
// This package does not attempt to hide the manual memory management of the C
// library, as such you must pay attention as you would if you where using the
// C library itself.
//
// This package uses the Chipmunk library, and as such it is also bound by it's
// license which can be found at:
//  https://github.com/slembcke/Chipmunk2D/blob/master/LICENSE.txt
//
package cp

/*
#cgo linux,amd64 LDFLAGS: /usr/lib/x86_64-linux-gnu/libm.a
#cgo !linux LDFLAGS: -lm
#cgo windows LDFLAGS: libchipmunk_windows_amd64.a
#include "chipmunk/chipmunk.h"
*/
import "C"

import "unsafe"

var (
	// Version string.
	VersionString = C.GoString(C.cpVersionString)
)

// Opaque Chipmunk types

type (
	Array               C.cpArray
	HashSet             C.cpHashSet
	CircleShape         C.cpCircleShape
	SegmentShape        C.cpSegmentShape
	PolyShape           C.cpPolyShape
	PinJoint            C.cpPinJoint
	SlideJoint          C.cpSlideJoint
	PivotJoint          C.cpPivotJoint
	GrooveJoint         C.cpGrooveJoint
	DampedSpring        C.cpDampedSpring
	DampedRotarySpring  C.cpDampedRotarySpring
	RotaryLimitJoint    C.cpRotaryLimitJoint
	RatchetJoint        C.cpRatchetJoint
	GearJoint           C.cpGearJoint
	SimpleMotorJoint    C.cpSimpleMotorJoint
	ContactBufferHeader C.cpContactBufferHeader
)

// Calculate the moment of inertia for a circle.
//
// r1 and r2 are the inner and outer diameters. A solid circle has an inner diameter of 0.
func MomentForCircle(m, r1, r2 float64, offset Vect) float64 {
	return float64(C.cpMomentForCircle(
		C.cpFloat(m),
		C.cpFloat(r1),
		C.cpFloat(r2),
		*(*C.cpVect)(unsafe.Pointer(&offset)),
	))
}

// Calculate area of a hollow circle.
//
// r1 and  r2 are the inner and outer diameters. A solid circle has an inner diameter of 0.
func AreaForCircle(r1, r2 float64) float64 {
	return float64(C.cpAreaForCircle(
		C.cpFloat(r1),
		C.cpFloat(r2),
	))
}

// Calculate the moment of inertia for a line segment.
//
// Beveling radius is not supported.
func MomentForSegment(m float64, a, b Vect, radius float64) float64 {
	return float64(C.cpMomentForSegment(
		C.cpFloat(m),
		*(*C.cpVect)(unsafe.Pointer(&a)),
		*(*C.cpVect)(unsafe.Pointer(&b)),
		C.cpFloat(radius),
	))
}

// Calculate the area of a fattened (capsule shaped) line segment.
func AreaForSegment(a, b Vect, radius float64) float64 {
	return float64(C.cpAreaForSegment(
		*(*C.cpVect)(unsafe.Pointer(&a)),
		*(*C.cpVect)(unsafe.Pointer(&b)),
		C.cpFloat(radius),
	))
}

// Calculate the moment of inertia for a solid polygon shape assuming it's
// center of gravity is at it's centroid. The offset is added to each vertex.
func MomentForPoly(m float64, verts []Vect, offset Vect, radius float64) float64 {
	return float64(C.cpMomentForPoly(
		C.cpFloat(m),
		C.int(len(verts)),
		(*C.cpVect)(unsafe.Pointer(&verts[0])),
		*(*C.cpVect)(unsafe.Pointer(&offset)),
		C.cpFloat(radius),
	))
}

// Calculate the signed area of a polygon. A Clockwise winding gives positive
// area.
//
// This is probably backwards from what you expect, but matches Chipmunk's the
// winding for poly shapes.
func AreaForPoly(verts []Vect, radius float64) float64 {
	return float64(C.cpAreaForPoly(
		C.int(len(verts)),
		(*C.cpVect)(unsafe.Pointer(&verts[0])),
		C.cpFloat(radius),
	))
}

// Calculate the natural centroid of a polygon.
func CentroidForPoly(verts []Vect) Vect {
	ret := C.cpCentroidForPoly(
		C.int(len(verts)),
		(*C.cpVect)(unsafe.Pointer(&verts[0])),
	)
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Calculate the moment of inertia for a solid box.
func MomentForBox(m, width, height float64) float64 {
	return float64(C.cpMomentForBox(
		C.cpFloat(m),
		C.cpFloat(width),
		C.cpFloat(height),
	))
}

// Calculate the moment of inertia for a solid box.
func MomentForBox2(m float64, box BB) float64 {
	return float64(C.cpMomentForBox2(
		C.cpFloat(m),
		*(*C.cpBB)(unsafe.Pointer(&box)),
	))
}

// Calculate the convex hull of a given set of points. Returns the count of points in the hull.
//
//  result must be a pointer to a  cpVect array with at least  count elements.
//  If  verts ==  result, then  verts will be reduced inplace.
//
//  first is an optional pointer to an integer to store where the first vertex
//  in the hull came from (i.e. verts[first] == result[0])
//
//  tol is the allowed amount to shrink the hull when simplifying it. A
//  tolerance of 0.0 creates an exact hull.
func ConvexHull(count int, verts, result *Vect, first *int, tol float64) int {
	return int(C.cpConvexHull(
		C.int(count),
		(*C.cpVect)(unsafe.Pointer(verts)),
		(*C.cpVect)(unsafe.Pointer(result)),
		(*C.int)(unsafe.Pointer(first)),
		C.cpFloat(tol),
	))
}

// Returns the closest point on the line segment ab, to the point p.
func ClosetPointOnSegment(p, a, b Vect) Vect {
	ret := C.cpClosetPointOnSegment(
		*(*C.cpVect)(unsafe.Pointer(&p)),
		*(*C.cpVect)(unsafe.Pointer(&a)),
		*(*C.cpVect)(unsafe.Pointer(&b)),
	)
	return *(*Vect)(unsafe.Pointer(&ret))
}
