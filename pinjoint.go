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

// Check if a constraint is a pin joint.
func (c *Constraint) IsPinJoint() bool {
	return goBool(C.cpConstraintIsPinJoint(c.c))
}

// Allocate and initialize a pin joint.
func PinJointNew(a, b *Body, anchorA, anchorB Vect) *Constraint {
	c := new(Constraint)
	c.c = C.cpPinJointNew(
		(*C.cpBody)(unsafe.Pointer(a)),
		(*C.cpBody)(unsafe.Pointer(b)),
		*(*C.cpVect)(unsafe.Pointer(&anchorA)),
		*(*C.cpVect)(unsafe.Pointer(&anchorB)),
	)
	if c.c == nil {
		return nil
	}
	C.cpConstraintSetUserData(c.c, C.cpDataPointer(unsafe.Pointer(c)))
	runtime.SetFinalizer(c, finalizeConstraint)
	return c
}

// Get the location of the first anchor relative to the first body.
func (c *Constraint) PinJointAnchorA() Vect {
	ret := C.cpPinJointGetAnchorA(c.c)
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Set the location of the first anchor relative to the first body.
func (c *Constraint) PinJointSetAnchorA(anchorA Vect) {
	C.cpPinJointSetAnchorA(
		c.c,
		*(*C.cpVect)(unsafe.Pointer(&anchorA)),
	)
}

// Get the location of the second anchor relative to the second body.
func (c *Constraint) PinJointAnchorB() Vect {
	ret := C.cpPinJointGetAnchorB(c.c)
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Set the location of the second anchor relative to the second body.
func (c *Constraint) PinJointSetAnchorB(anchorB Vect) {
	C.cpPinJointSetAnchorB(
		c.c,
		*(*C.cpVect)(unsafe.Pointer(&anchorB)),
	)
}

// Get the distance the joint will maintain between the two anchors.
func (c *Constraint) PinJointDist() float64 {
	return float64(C.cpPinJointGetDist(c.c))
}

// Set the distance the joint will maintain between the two anchors.
func (c *Constraint) PinJointSetDist(dist float64) {
	C.cpPinJointSetDist(c.c, C.cpFloat(dist))
}
