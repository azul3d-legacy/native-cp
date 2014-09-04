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

// Check if a constraint is a ratchet joint.
func (c *Constraint) IsGearJoint() bool {
	return goBool(C.cpConstraintIsGearJoint(c.c))
}

// Allocate and initialize a gear joint.
func GearJointNew(a, b *Body, phase, ratio float64) *Constraint {
	c := new(Constraint)
	c.c = C.cpGearJointNew(
		a.c,
		b.c,
		C.cpFloat(phase),
		C.cpFloat(ratio),
	)
	if c.c == nil {
		return nil
	}
	C.cpConstraintSetUserData(c.c, C.cpDataPointer(unsafe.Pointer(c)))
	runtime.SetFinalizer(c, finalizeConstraint)
	return c
}

// Get the phase offset of the gears.
func (c *Constraint) GearJointPhase() float64 {
	return float64(C.cpGearJointGetPhase(c.c))
}

// Set the phase offset of the gears.
func (c *Constraint) GearJointSetPhase(phase float64) {
	C.cpGearJointSetPhase(c.c, C.cpFloat(phase))
}

// Get the angular distance of each ratchet.
func (c *Constraint) GearJointRatio() float64 {
	return float64(C.cpGearJointGetRatio(c.c))
}

// Set the ratio of a gear joint.
func (c *Constraint) GearJointSetRatio(ratio float64) {
	C.cpGearJointSetRatio(c.c, C.cpFloat(ratio))
}
