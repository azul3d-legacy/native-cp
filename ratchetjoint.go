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
func (c *Constraint) IsRatchetJoint() bool {
	return goBool(C.cpConstraintIsRatchetJoint(c.c))
}

// Allocate and initialize a ratchet joint.
func RatchetJointNew(a, b *Body, phase, ratchet float64) *Constraint {
	c := new(Constraint)
	c.c = C.cpRatchetJointNew(
		a.c,
		b.c,
		C.cpFloat(phase),
		C.cpFloat(ratchet),
	)
	if c.c == nil {
		return nil
	}
	C.cpConstraintSetUserData(c.c, C.cpDataPointer(unsafe.Pointer(c)))
	runtime.SetFinalizer(c, finalizeConstraint)
	return c
}

// Get the angle of the current ratchet tooth.
func (c *Constraint) RatchetJointAngle() float64 {
	return float64(C.cpRatchetJointGetAngle(c.c))
}

// Set the angle of the current ratchet tooth.
func (c *Constraint) RatchetJointSetAngle(angle float64) {
	C.cpRatchetJointSetAngle(c.c, C.cpFloat(angle))
}

// Get the phase offset of the ratchet.
func (c *Constraint) RatchetJointPhase() float64 {
	return float64(C.cpRatchetJointGetPhase(c.c))
}

// Get the phase offset of the ratchet.
func (c *Constraint) RatchetJointSetPhase(phase float64) {
	C.cpRatchetJointSetPhase(c.c, C.cpFloat(phase))
}

// Get the angular distance of each ratchet.
func (c *Constraint) RatchetJointRatchet() float64 {
	return float64(C.cpRatchetJointGetRatchet(c.c))
}

// Set the angular distance of each ratchet.
func (c *Constraint) RatchetJointSetRatchet(ratchet float64) {
	C.cpRatchetJointSetRatchet(c.c, C.cpFloat(ratchet))
}
