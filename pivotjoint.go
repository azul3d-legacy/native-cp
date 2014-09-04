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

// Check if a constraint is a pivot joint.
func (c *Constraint) IsPivotJoint() bool {
	return goBool(C.cpConstraintIsPivotJoint(c.c))
}

// Allocate and initialize a pivot joint.
func PivotJointNew(a, b *Body, pivot Vect) *Constraint {
	c := &Constraint{
		aBodyRef: a,
		bBodyRef: b,
		c:        C.cpPivotJointNew(a.c, b.c, pivot.c()),
	}
	if c.c == nil {
		return nil
	}
	C.cpConstraintSetUserData(c.c, C.cpDataPointer(unsafe.Pointer(c)))
	runtime.SetFinalizer(c, finalizeConstraint)
	return c
}

// Allocate and initialize a pivot joint with specific anchors.
func PivotJointNew2(a, b *Body, anchorA, anchorB Vect) *Constraint {
	c := &Constraint{
		aBodyRef: a,
		bBodyRef: b,
		c:        C.cpPivotJointNew2(a.c, b.c, anchorA.c(), anchorB.c()),
	}
	if c.c == nil {
		return nil
	}
	C.cpConstraintSetUserData(c.c, C.cpDataPointer(unsafe.Pointer(c)))
	runtime.SetFinalizer(c, finalizeConstraint)
	return c
}

// Get the location of the first anchor relative to the first body.
func (c *Constraint) PivotJointAnchorA() Vect {
	return goVect(C.cpPivotJointGetAnchorA(c.c))
}

// Set the location of the first anchor relative to the first body.
func (c *Constraint) PivotJointSetAnchorA(anchorA Vect) {
	C.cpPivotJointSetAnchorA(c.c, anchorA.c())
}

// Get the location of the second anchor relative to the second body.
func (c *Constraint) PivotJointAnchorB() Vect {
	return goVect(C.cpPivotJointGetAnchorB(c.c))
}

// Set the location of the second anchor relative to the second body.
func (c *Constraint) PivotJointSetAnchorB(anchorB Vect) {
	C.cpPivotJointSetAnchorB(c.c, anchorB.c())
}
