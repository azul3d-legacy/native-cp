// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cp

/*
#include "chipmunk/chipmunk.h"
extern void pre_go_chipmunk_rotary_spring_torque_func(struct cpConstraint *spring, cpFloat relativeAngle);
*/
import "C"

import (
	"unsafe"
)

// Check if a constraint is a damped rotary spring.
func (c *Constraint) IsDampedRotarySpring() bool {
	return goBool(C.cpConstraintIsDampedRotarySpring(c.c))
}

// Allocate and initialize a damped rotary spring.
func DampedRotarySpringNew(a, b *Body, restAngle, stiffness, damping float64) *Constraint {
	c := new(Constraint)
	c.c = C.cpDampedRotarySpringNew(
		(*C.cpBody)(unsafe.Pointer(a)),
		(*C.cpBody)(unsafe.Pointer(b)),
		C.cpFloat(restAngle),
		C.cpFloat(stiffness),
		C.cpFloat(damping),
	)
	if c.c == nil {
		return nil
	}
	C.cpConstraintSetUserData(c.c, C.cpDataPointer(unsafe.Pointer(c)))
	return c
}

// Get the rest length of the spring.
func (c *Constraint) DampedRotarySpringRestAngle() float64 {
	return float64(C.cpDampedRotarySpringGetRestAngle(c.c))
}

// Set the rest length of the spring.
func (c *Constraint) RotarySpringSetRestAngle(restAngle float64) {
	C.cpDampedRotarySpringSetRestAngle(c.c, C.cpFloat(restAngle))
}

// Get the stiffness of the spring in force/distance.
func (c *Constraint) DampedRotarySpringStiffness() float64 {
	return float64(C.cpDampedRotarySpringGetStiffness(c.c))
}

// Set the stiffness of the spring in force/distance.
func (c *Constraint) DampedRotarySpringSetStiffness(stiffness float64) {
	C.cpDampedRotarySpringSetStiffness(c.c, C.cpFloat(stiffness))
}

// Get the damping of the spring.
func (c *Constraint) DampedRotarySpringDamping() float64 {
	return float64(C.cpDampedRotarySpringGetDamping(c.c))
}

// Set the damping of the spring.
func (c *Constraint) DampedRotarySpringSetDamping(damping float64) {
	C.cpDampedRotarySpringSetDamping(c.c, C.cpFloat(damping))
}

// Set the damping rotary spring torque callback function.
func (c *Constraint) DampedRotarySpringSetTorqueFunc(f func(spring *Constraint, relativeAngle float64) float64) {
	c.dampedRotarySpringTorqueFunc = f
	C.cpDampedRotarySpringSetSpringTorqueFunc(
		c.c,
		(*[0]byte)(unsafe.Pointer(C.pre_go_chipmunk_rotary_spring_torque_func)),
	)
}

// Get the damping rotary spring torque callback function.
func (c *Constraint) DampedRotarySpringTorqueFunc() func(spring *Constraint, relativeAngle float64) float64 {
	return c.dampedRotarySpringTorqueFunc
}

//export go_chipmunk_rotary_spring_torque_func
func go_chipmunk_rotary_spring_torque_func(cspring unsafe.Pointer, crelativeAngle C.cpFloat) float64 {
	spring := goConstraint((*C.cpConstraint)(cspring))
	relativeAngle := *(*float64)(unsafe.Pointer(&crelativeAngle))
	return spring.dampedRotarySpringTorqueFunc(spring, relativeAngle)
}
