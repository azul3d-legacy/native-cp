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

// Check if a constraint is a groove joint.
func (c *Constraint) IsGrooveJoint() bool {
	return goBool(C.cpConstraintIsGrooveJoint(c.c))
}

// Allocate and initialize a groove joint.
func GrooveJointNew(a, b *Body, grooveA, grooveB, anchorB Vect) *Constraint {
	c := new(Constraint)
	c.c = C.cpGrooveJointNew(
		(*C.cpBody)(unsafe.Pointer(a)),
		(*C.cpBody)(unsafe.Pointer(b)),
		*(*C.cpVect)(unsafe.Pointer(&grooveA)),
		*(*C.cpVect)(unsafe.Pointer(&grooveB)),
		*(*C.cpVect)(unsafe.Pointer(&anchorB)),
	)
	if c.c == nil {
		return nil
	}
	C.cpConstraintSetUserData(c.c, C.cpDataPointer(unsafe.Pointer(c)))
	return c
}

// Get the first endpoint of the groove relative to the first body.
func (c *Constraint) GrooveJointGrooveA() Vect {
	ret := C.cpGrooveJointGetGrooveA(c.c)
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Set the first endpoint of the groove relative to the first body.
func (c *Constraint) GrooveJointSetGrooveA(grooveA Vect) {
	C.cpGrooveJointSetGrooveA(
		c.c,
		*(*C.cpVect)(unsafe.Pointer(&grooveA)),
	)
}

// Get the first endpoint of the groove relative to the first body.
func (c *Constraint) GrooveJointGrooveB() Vect {
	ret := C.cpGrooveJointGetGrooveB(c.c)
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Set the first endpoint of the groove relative to the first body.
func (c *Constraint) GrooveJointSetGrooveB(grooveB Vect) {
	C.cpGrooveJointSetGrooveB(
		c.c,
		*(*C.cpVect)(unsafe.Pointer(&grooveB)),
	)
}

// Get the location of the second anchor relative to the second body.
func (c *Constraint) GrooveJointAnchorB() Vect {
	ret := C.cpGrooveJointGetAnchorB(c.c)
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Set the location of the second anchor relative to the second body.
func (c *Constraint) GrooveJointSetAnchorB(grooveB Vect) {
	C.cpGrooveJointSetAnchorB(
		c.c,
		*(*C.cpVect)(unsafe.Pointer(&grooveB)),
	)
}
