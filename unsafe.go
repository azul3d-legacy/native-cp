// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cp

/*
#include "chipmunk/include/chipmunk/chipmunk.h"
#include "chipmunk/include/chipmunk/chipmunk_unsafe.h"
*/
import "C"

import "unsafe"

// Set the radius of a circle shape.
//
// This function is used for mutating collision shapes. Chipmunk does not have
// any way to get velocity information on changing shapes, so the results will
// be unrealistic. This function is considered 'unsafe' by Chipmunk.
func (shape *Shape) CircleShapeSetRadius(radius float64) {
	C.cpCircleShapeSetRadius(
		(*C.cpShape)(unsafe.Pointer(shape)),
		C.cpFloat(radius),
	)
}

// Set the offset of a circle shape.
//
// This function is used for mutating collision shapes. Chipmunk does not have
// any way to get velocity information on changing shapes, so the results will
// be unrealistic. This function is considered 'unsafe' by Chipmunk.
func (shape *Shape) CircleShapeSetOffset(offset Vect) {
	C.cpCircleShapeSetOffset(
		(*C.cpShape)(unsafe.Pointer(shape)),
		offset.c(),
	)
}

// Set the endpoints of a segment shape.
//
// This function is used for mutating collision shapes. Chipmunk does not have
// any way to get velocity information on changing shapes, so the results will
// be unrealistic. This function is considered 'unsafe' by Chipmunk.
func (shape *Shape) SegmentShapeSetEndpoints(a, b Vect) {
	C.cpSegmentShapeSetEndpoints(
		(*C.cpShape)(unsafe.Pointer(shape)),
		a.c(),
		b.c(),
	)
}

// Set the radius of a segment shape.
//
// This function is used for mutating collision shapes. Chipmunk does not have
// any way to get velocity information on changing shapes, so the results will
// be unrealistic. This function is considered 'unsafe' by Chipmunk.
func (shape *Shape) SegmentShapeSetRadius(radius float64) {
	C.cpSegmentShapeSetRadius(
		(*C.cpShape)(unsafe.Pointer(shape)),
		C.cpFloat(radius),
	)
}

// Set the vertexes of a poly shape.
//
// This function is used for mutating collision shapes. Chipmunk does not have
// any way to get velocity information on changing shapes, so the results will
// be unrealistic. This function is considered 'unsafe' by Chipmunk.
func (shape *Shape) PolyShapeSetVerts(verts []Vect, transform Transform) {
	C.cpPolyShapeSetVerts(
		(*C.cpShape)(unsafe.Pointer(shape)),
		C.int(len(verts)),
		(*C.cpVect)(unsafe.Pointer(&verts[0])),
		*(*C.cpTransform)(unsafe.Pointer(&transform)),
	)
}

// Set the vertexes of a poly shape.
//
// This function is used for mutating collision shapes. Chipmunk does not have
// any way to get velocity information on changing shapes, so the results will
// be unrealistic. This function is considered 'unsafe' by Chipmunk.
func (shape *Shape) PolyShapeSetVertsRaw(verts []Vect) {
	C.cpPolyShapeSetVertsRaw(
		(*C.cpShape)(unsafe.Pointer(shape)),
		C.int(len(verts)),
		(*C.cpVect)(unsafe.Pointer(&verts[0])),
	)
}

// Set the radius of a poly shape.
//
// This function is used for mutating collision shapes. Chipmunk does not have
// any way to get velocity information on changing shapes, so the results will
// be unrealistic. This function is considered 'unsafe' by Chipmunk.
func (shape *Shape) PolyShapeSetRadius(radius float64) {
	C.cpPolyShapeSetRadius(
		(*C.cpShape)(unsafe.Pointer(shape)),
		C.cpFloat(radius),
	)
}
