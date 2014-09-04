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

// Check if a constraint is a groove joint.
func (c *Constraint) IsGrooveJoint() bool {
	return goBool(C.cpConstraintIsGrooveJoint(c.c))
}

// Allocate and initialize a groove joint.
func GrooveJointNew(a, b *Body, grooveA, grooveB, anchorB Vect) *Constraint {
	c := new(Constraint)
	c.c = C.cpGrooveJointNew(a.c, b.c, grooveA.c(), grooveB.c(), anchorB.c())
	if c.c == nil {
		return nil
	}
	C.cpConstraintSetUserData(c.c, C.cpDataPointer(unsafe.Pointer(c)))
	runtime.SetFinalizer(c, finalizeConstraint)
	return c
}

// Get the first endpoint of the groove relative to the first body.
func (c *Constraint) GrooveJointGrooveA() Vect {
	return goVect(C.cpGrooveJointGetGrooveA(c.c))
}

// Set the first endpoint of the groove relative to the first body.
func (c *Constraint) GrooveJointSetGrooveA(grooveA Vect) {
	C.cpGrooveJointSetGrooveA(c.c, grooveA.c())
}

// Get the first endpoint of the groove relative to the first body.
func (c *Constraint) GrooveJointGrooveB() Vect {
	return goVect(C.cpGrooveJointGetGrooveB(c.c))
}

// Set the first endpoint of the groove relative to the first body.
func (c *Constraint) GrooveJointSetGrooveB(grooveB Vect) {
	C.cpGrooveJointSetGrooveB(c.c, grooveB.c())
}

// Get the location of the second anchor relative to the second body.
func (c *Constraint) GrooveJointAnchorB() Vect {
	return goVect(C.cpGrooveJointGetAnchorB(c.c))
}

// Set the location of the second anchor relative to the second body.
func (c *Constraint) GrooveJointSetAnchorB(grooveB Vect) {
	C.cpGrooveJointSetAnchorB(c.c, grooveB.c())
}
