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

// Check if a constraint is a slide joint.
func (c *Constraint) IsSlideJoint() bool {
	return goBool(C.cpConstraintIsSlideJoint(c.c))
}

// Allocate and initialize a slide joint.
func SlideJointNew(a, b *Body, anchorA, anchorB Vect, min, max float64) *Constraint {
	c := new(Constraint)
	c.c = C.cpSlideJointNew(
		(*C.cpBody)(unsafe.Pointer(a)),
		(*C.cpBody)(unsafe.Pointer(b)),
		*(*C.cpVect)(unsafe.Pointer(&anchorA)),
		*(*C.cpVect)(unsafe.Pointer(&anchorB)),
		C.cpFloat(min),
		C.cpFloat(max),
	)
	if c.c == nil {
		return nil
	}
	C.cpConstraintSetUserData(c.c, C.cpDataPointer(unsafe.Pointer(c)))
	runtime.SetFinalizer(c, finalizeConstraint)
	return c
}

// Get the location of the first anchor relative to the first body.
func (c *Constraint) SlideJointAnchorA() Vect {
	ret := C.cpSlideJointGetAnchorA(c.c)
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Set the location of the first anchor relative to the first body.
func (c *Constraint) SlideJointSetAnchorA(anchorA Vect) {
	C.cpSlideJointSetAnchorA(
		c.c,
		*(*C.cpVect)(unsafe.Pointer(&anchorA)),
	)
}

// Get the location of the second anchor relative to the second body.
func (c *Constraint) SlideJointAnchorB() Vect {
	ret := C.cpSlideJointGetAnchorB(c.c)
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Set the location of the second anchor relative to the second body.
func (c *Constraint) SlideJointSetAnchorB(anchorB Vect) {
	C.cpSlideJointSetAnchorB(
		c.c,
		*(*C.cpVect)(unsafe.Pointer(&anchorB)),
	)
}

// Get the minimum distance the joint will maintain between the two anchors.
func (c *Constraint) SlideJointMin() float64 {
	return float64(C.cpSlideJointGetMin(c.c))
}

// Set the minimum distance the joint will maintain between the two anchors.
func (c *Constraint) SlideJointSetMin(min float64) {
	C.cpSlideJointSetMin(c.c, C.cpFloat(min))
}

// Get the maximum distance the joint will maintain between the two anchors.
func (c *Constraint) SlideJointMax() float64 {
	return float64(C.cpSlideJointGetMax(c.c))
}

// Set the maximum distance the joint will maintain between the two anchors.
func (c *Constraint) SlideJointSetMax(max float64) {
	C.cpSlideJointSetMax(c.c, C.cpFloat(max))
}
