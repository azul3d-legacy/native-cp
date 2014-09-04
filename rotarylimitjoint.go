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

// Check if a constraint is a rotary limit joint.
func (c *Constraint) IsRotaryLimitJoint() bool {
	return goBool(C.cpConstraintIsRotaryLimitJoint(c.c))
}

// Allocate and initialize a damped rotary limit joint.
func RotaryLimitJointNew(a, b *Body, min, max float64) *Constraint {
	c := new(Constraint)
	c.c = C.cpRotaryLimitJointNew(
		a.c,
		b.c,
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

// Get the minimum distance the joint will maintain between the two anchors.
func (c *Constraint) RotaryLimitJointMin() float64 {
	return float64(C.cpRotaryLimitJointGetMin(c.c))
}

// Set the minimum distance the joint will maintain between the two anchors.
func (c *Constraint) RotaryLimitJointSetMin(min float64) {
	C.cpRotaryLimitJointSetMin(c.c, C.cpFloat(min))
}

// Get the maximum distance the joint will maintain between the two anchors.
func (c *Constraint) RotaryLimitJointMax() float64 {
	return float64(C.cpRotaryLimitJointGetMax(c.c))
}

// Set the maximum distance the joint will maintain between the two anchors.
func (c *Constraint) RotaryLimitJointSetMax(max float64) {
	C.cpRotaryLimitJointSetMax(c.c, C.cpFloat(max))
}
