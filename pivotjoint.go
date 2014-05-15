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

// Check if a constraint is a pivot joint.
func (c *Constraint) IsPivotJoint() bool {
	return goBool(C.cpConstraintIsPivotJoint(c.c))
}

// Allocate and initialize a pivot joint.
func PivotJointNew(a, b *Body, pivot Vect) *Constraint {
	c := new(Constraint)
	c.c = C.cpPivotJointNew(
		(*C.cpBody)(unsafe.Pointer(a)),
		(*C.cpBody)(unsafe.Pointer(b)),
		*(*C.cpVect)(unsafe.Pointer(&pivot)),
	)
	if c.c == nil {
		return nil
	}
	C.cpConstraintSetUserData(c.c, C.cpDataPointer(unsafe.Pointer(c)))
	return c
}

// Allocate and initialize a pivot joint with specific anchors.
func PivotJointNew2(a, b *Body, anchorA, anchorB Vect) *Constraint {
	c := new(Constraint)
	c.c = C.cpPivotJointNew2(
		(*C.cpBody)(unsafe.Pointer(a)),
		(*C.cpBody)(unsafe.Pointer(b)),
		*(*C.cpVect)(unsafe.Pointer(&anchorA)),
		*(*C.cpVect)(unsafe.Pointer(&anchorB)),
	)
	if c.c == nil {
		return nil
	}
	C.cpConstraintSetUserData(c.c, C.cpDataPointer(unsafe.Pointer(c)))
	return c
}

// Get the location of the first anchor relative to the first body.
func (c *Constraint) PivotJointAnchorA() Vect {
	ret := C.cpPivotJointGetAnchorA(c.c)
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Set the location of the first anchor relative to the first body.
func (c *Constraint) PivotJointSetAnchorA(anchorA Vect) {
	C.cpPivotJointSetAnchorA(
		c.c,
		*(*C.cpVect)(unsafe.Pointer(&anchorA)),
	)
}

// Get the location of the second anchor relative to the second body.
func (c *Constraint) PivotJointAnchorB() Vect {
	ret := C.cpPivotJointGetAnchorB(c.c)
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Set the location of the second anchor relative to the second body.
func (c *Constraint) PivotJointSetAnchorB(anchorB Vect) {
	C.cpPivotJointSetAnchorB(
		c.c,
		*(*C.cpVect)(unsafe.Pointer(&anchorB)),
	)
}
