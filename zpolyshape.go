// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cp

/*
#include "chipmunk/chipmunk.h"
*/
import "C"

import (
	"unsafe"
)

// Allocate and initialize a polygon shape with rounded corners.
//
// A convex hull will be created from the vertexes.
func (b *Body) PolyShapeNew(verts []Vect, transform Transform, radius float64) *Shape {
	s := new(Shape)
	s.c = C.cpPolyShapeNew(
		b.c,
		C.int(len(verts)),
		(*C.cpVect)(unsafe.Pointer(&verts[0])),
		*(*C.cpTransform)(unsafe.Pointer(&transform)),
		C.cpFloat(radius),
	)
	if s.c == nil {
		return nil
	}
	C.cpShapeSetUserData(s.c, C.cpDataPointer(unsafe.Pointer(s)))
	return s
}

// Allocate and initialize a box shaped polygon shape.
func (b *Body) BoxShapeNew(width, height, radius float64) *Shape {
	s := new(Shape)
	s.c = C.cpBoxShapeNew(
		b.c,
		C.cpFloat(width),
		C.cpFloat(height),
		C.cpFloat(radius),
	)
	if s.c == nil {
		return nil
	}
	C.cpShapeSetUserData(s.c, C.cpDataPointer(unsafe.Pointer(s)))
	return s
}

// Allocate and initialize an offset box shaped polygon shape.
func (b *Body) BoxShapeNew2(box BB, radius float64) *Shape {
	s := new(Shape)
	s.c = C.cpBoxShapeNew2(
		b.c,
		*(*C.cpBB)(unsafe.Pointer(&box)),
		C.cpFloat(radius),
	)
	if s.c == nil {
		return nil
	}
	C.cpShapeSetUserData(s.c, C.cpDataPointer(unsafe.Pointer(s)))
	return s
}

// Get the number of verts in a polygon shape.
func (s *Shape) PolyCount() int {
	return int(C.cpPolyShapeGetCount(s.c))
}

// Get the ith vertex of a polygon shape.
func (s *Shape) PolyVert(index int) Vect {
	ret := C.cpPolyShapeGetVert(s.c, C.int(index))
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Get the radius of a polygon shape.
func (s *Shape) PolyRadius() float64 {
	return float64(C.cpPolyShapeGetRadius(s.c))
}
