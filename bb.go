// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cp

/*
#include "chipmunk/include/chipmunk/chipmunk.h"
*/
import "C"

import "unsafe"

// Convenience constructor for BB structs.
func BBNew(l, b, r, t float64) BB {
	return BB{
		L: l,
		B: b,
		R: r,
		T: t,
	}
}

// Constructs a BB centered on a point with the given extents (half sizes).
func BBNewForExtents(c Vect, hw, hh float64) BB {
	ret := C.cpBBNewForExtents(
		*(*C.cpVect)(unsafe.Pointer(&c)),
		C.cpFloat(hw),
		C.cpFloat(hh),
	)
	return *(*BB)(unsafe.Pointer(&ret))
}

// Constructs a BB for a circle with the given position and radius.
func BBNewForCircle(p Vect, r float64) BB {
	ret := C.cpBBNewForCircle(
		*(*C.cpVect)(unsafe.Pointer(&p)),
		C.cpFloat(r),
	)
	return *(*BB)(unsafe.Pointer(&ret))
}

// Returns true if a and b intersect.
func (a BB) Intersects(b BB) bool {
	return goBool(C.cpBBIntersects(
		*(*C.cpBB)(unsafe.Pointer(&a)),
		*(*C.cpBB)(unsafe.Pointer(&b)),
	))
}

// Returns true if  other lies completely within bb.
func (bb BB) ContainsBB(other BB) bool {
	return goBool(C.cpBBContainsBB(
		*(*C.cpBB)(unsafe.Pointer(&bb)),
		*(*C.cpBB)(unsafe.Pointer(&other)),
	))
}

// Returns true if bb contains v.
func (bb BB) ContainsVect(v Vect) bool {
	return goBool(C.cpBBContainsVect(
		*(*C.cpBB)(unsafe.Pointer(&bb)),
		*(*C.cpVect)(unsafe.Pointer(&v)),
	))
}

// Returns a bounding box that holds both bounding boxes.
func (a BB) Merge(b BB) BB {
	ret := C.cpBBMerge(
		*(*C.cpBB)(unsafe.Pointer(&a)),
		*(*C.cpBB)(unsafe.Pointer(&b)),
	)
	return *(*BB)(unsafe.Pointer(&ret))
}

// Returns a bounding box that holds both bb and v.
func (bb BB) Expand(v Vect) BB {
	ret := C.cpBBExpand(
		*(*C.cpBB)(unsafe.Pointer(&bb)),
		*(*C.cpVect)(unsafe.Pointer(&v)),
	)
	return *(*BB)(unsafe.Pointer(&ret))
}

// Returns the center of a bounding box.
func (bb BB) Center() Vect {
	ret := C.cpBBCenter(
		*(*C.cpBB)(unsafe.Pointer(&bb)),
	)
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Returns the area of the bounding box.
func (bb BB) Area() float64 {
	return float64(C.cpBBArea(
		*(*C.cpBB)(unsafe.Pointer(&bb)),
	))
}

// Merges a and b and returns the area of the merged bounding box.
func (a BB) MergedArea(b BB) float64 {
	return float64(C.cpBBMergedArea(
		*(*C.cpBB)(unsafe.Pointer(&a)),
		*(*C.cpBB)(unsafe.Pointer(&b)),
	))
}

// Returns the fraction along the segment query the BB is hit. Returns
// INFINITY if it doesn't hit.
func (bb BB) SegmentQuery(a, b Vect) float64 {
	return float64(C.cpBBSegmentQuery(
		*(*C.cpBB)(unsafe.Pointer(&bb)),
		*(*C.cpVect)(unsafe.Pointer(&a)),
		*(*C.cpVect)(unsafe.Pointer(&b)),
	))
}

// Return true if the bounding box intersects the line segment with ends  a and  b.
func (bb BB) IntersectsSegment(a, b Vect) bool {
	return goBool(C.cpBBIntersectsSegment(
		*(*C.cpBB)(unsafe.Pointer(&bb)),
		*(*C.cpVect)(unsafe.Pointer(&a)),
		*(*C.cpVect)(unsafe.Pointer(&b)),
	))
}

// Clamp a vector to a bounding box.
func (bb BB) ClampVect(v Vect) Vect {
	ret := C.cpBBClampVect(
		*(*C.cpBB)(unsafe.Pointer(&bb)),
		*(*C.cpVect)(unsafe.Pointer(&v)),
	)
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Wrap a vector to a bounding box.
func (bb BB) WrapVect(v Vect) Vect {
	ret := C.cpBBWrapVect(
		*(*C.cpBB)(unsafe.Pointer(&bb)),
		*(*C.cpVect)(unsafe.Pointer(&v)),
	)
	return *(*Vect)(unsafe.Pointer(&ret))
}
