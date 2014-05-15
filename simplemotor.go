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

// Check if a constraint is a simple motor.
func (c *Constraint) IsSimpleMotor() bool {
	return goBool(C.cpConstraintIsSimpleMotor(c.c))
}

// Allocate and initialize a simple motor.
func SimpleMotorNew(a, b *Body, rate float64) *Constraint {
	c := new(Constraint)
	c.c = C.cpSimpleMotorNew(
		(*C.cpBody)(unsafe.Pointer(a)),
		(*C.cpBody)(unsafe.Pointer(b)),
		C.cpFloat(rate),
	)
	if c.c == nil {
		return nil
	}
	C.cpConstraintSetUserData(c.c, C.cpDataPointer(unsafe.Pointer(c)))
	return c
}

// Get the rate of the motor.
func (c *Constraint) SimpleMotorRate() float64 {
	return float64(C.cpSimpleMotorGetRate(c.c))
}

// Set the rate of the motor.
func (c *Constraint) SimpleMotorSetRate(rate float64) {
	C.cpSimpleMotorSetRate(c.c, C.cpFloat(rate))
}
