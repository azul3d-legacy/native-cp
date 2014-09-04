// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cp

/*
#include "chipmunk/include/chipmunk/chipmunk.h"
*/
import "C"

// Chipmunk's axis-aligned 2D bounding box type. (left, bottom, right, top)
type BB struct {
	L float64
	B float64
	R float64
	T float64
}

// c converts a BB to a C.cpBB.
func (b BB) c() C.cpBB {
	var cp C.cpBB
	cp.l = C.cpFloat(b.L)
	cp.b = C.cpFloat(b.B)
	cp.r = C.cpFloat(b.R)
	cp.t = C.cpFloat(b.T)
	return cp
}

// goBB converts C.cpBB to a Go BB.
func goBB(b C.cpBB) BB {
	return BB{
		L: float64(b.l),
		B: float64(b.b),
		R: float64(b.r),
		T: float64(b.t),
	}
}

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
	return goBB(C.cpBBNewForExtents(c.c(), C.cpFloat(hw), C.cpFloat(hh)))
}

// Constructs a BB for a circle with the given position and radius.
func BBNewForCircle(p Vect, r float64) BB {
	return goBB(C.cpBBNewForCircle(p.c(), C.cpFloat(r)))
}

// Returns true if a and b intersect.
func (a BB) Intersects(b BB) bool {
	return goBool(C.cpBBIntersects(a.c(), b.c()))
}

// Returns true if  other lies completely within bb.
func (bb BB) ContainsBB(other BB) bool {
	return goBool(C.cpBBContainsBB(bb.c(), other.c()))
}

// Returns true if bb contains v.
func (bb BB) ContainsVect(v Vect) bool {
	return goBool(C.cpBBContainsVect(bb.c(), v.c()))
}

// Returns a bounding box that holds both bounding boxes.
func (a BB) Merge(b BB) BB {
	return goBB(C.cpBBMerge(a.c(), b.c()))
}

// Returns a bounding box that holds both bb and v.
func (bb BB) Expand(v Vect) BB {
	return goBB(C.cpBBExpand(bb.c(), v.c()))
}

// Returns the center of a bounding box.
func (bb BB) Center() Vect {
	return goVect(C.cpBBCenter(bb.c()))
}

// Returns the area of the bounding box.
func (bb BB) Area() float64 {
	return float64(C.cpBBArea(bb.c()))
}

// Merges a and b and returns the area of the merged bounding box.
func (a BB) MergedArea(b BB) float64 {
	return float64(C.cpBBMergedArea(a.c(), b.c()))
}

// Returns the fraction along the segment query the BB is hit. Returns
// INFINITY if it doesn't hit.
func (bb BB) SegmentQuery(a, b Vect) float64 {
	return float64(C.cpBBSegmentQuery(bb.c(), a.c(), b.c()))
}

// Return true if the bounding box intersects the line segment with ends  a and  b.
func (bb BB) IntersectsSegment(a, b Vect) bool {
	return goBool(C.cpBBIntersectsSegment(bb.c(), a.c(), b.c()))
}

// Clamp a vector to a bounding box.
func (bb BB) ClampVect(v Vect) Vect {
	return goVect(C.cpBBClampVect(bb.c(), v.c()))
}

// Wrap a vector to a bounding box.
func (bb BB) WrapVect(v Vect) Vect {
	return goVect(C.cpBBWrapVect(bb.c(), v.c()))
}
