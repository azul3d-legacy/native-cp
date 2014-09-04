// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cp

/*
#include "chipmunk/include/chipmunk/chipmunk.h"
extern void pre_go_chipmunk_damped_spring_force_func(struct cpConstraint *spring, cpFloat dist);
*/
import "C"

import (
	"runtime"
	"unsafe"
)

// Check if a constraint is a damped spring.
func (c *Constraint) IsDampedSpring() bool {
	return goBool(C.cpConstraintIsDampedSpring(c.c))
}

// Allocate and initialize a damped spring.
func DampedSpringNew(a, b *Body, anchorA, anchorB Vect, restLength, stiffness, damping float64) *Constraint {
	c := &Constraint{
		aBodyRef: a,
		bBodyRef: b,
		c: C.cpDampedSpringNew(
			a.c,
			b.c,
			anchorA.c(),
			anchorB.c(),
			C.cpFloat(restLength),
			C.cpFloat(stiffness),
			C.cpFloat(damping),
		),
	}
	if c.c == nil {
		return nil
	}
	C.cpConstraintSetUserData(c.c, C.cpDataPointer(unsafe.Pointer(c)))
	runtime.SetFinalizer(c, finalizeConstraint)
	return c
}

// Get the rest length of the spring.
func (c *Constraint) DampedSpringRestLength() float64 {
	return float64(C.cpDampedSpringGetRestLength(c.c))
}

// Set the rest length of the spring.
func (c *Constraint) DampedSpringSetRestLength(restLength float64) {
	C.cpDampedSpringSetRestLength(c.c, C.cpFloat(restLength))
}

// Get the stiffness of the spring in force/distance.
func (c *Constraint) DampedSpringStiffness() float64 {
	return float64(C.cpDampedSpringGetStiffness(c.c))
}

// Set the stiffness of the spring in force/distance.
func (c *Constraint) DampedSpringSetStiffness(stiffness float64) {
	C.cpDampedSpringSetStiffness(c.c, C.cpFloat(stiffness))
}

// Get the damping of the spring.
func (c *Constraint) DampedSpringDamping() float64 {
	return float64(C.cpDampedSpringGetDamping(c.c))
}

// Set the damping of the spring.
func (c *Constraint) DampedSpringSetDamping(damping float64) {
	C.cpDampedSpringSetDamping(c.c, C.cpFloat(damping))
}

// Get the location of the first anchor relative to the first body.
func (c *Constraint) DampedSpringAnchorA() Vect {
	return goVect(C.cpDampedSpringGetAnchorA(c.c))
}

// Set the location of the first anchor relative to the first body.
func (c *Constraint) DampedSpringSetAnchorA(anchorA Vect) {
	C.cpDampedSpringSetAnchorA(c.c, anchorA.c())
}

// Get the location of the second anchor relative to the second body.
func (c *Constraint) DampedSpringAnchorB() Vect {
	return goVect(C.cpDampedSpringGetAnchorB(c.c))
}

// Set the location of the second anchor relative to the second body.
func (c *Constraint) DampedSpringSetAnchorB(anchorA Vect) {
	C.cpDampedSpringSetAnchorB(c.c, anchorA.c())
}

// Set the damping spring force callback function.
func (c *Constraint) DampedSpringSetForceFunc(f func(spring *Constraint, dist float64) float64) {
	c.dampedSpringForceFunc = f
	C.cpDampedSpringSetSpringForceFunc(
		c.c,
		(*[0]byte)(unsafe.Pointer(C.pre_go_chipmunk_damped_spring_force_func)),
	)
}

// Get the damping rotary spring torque callback function.
func (c *Constraint) DampedSpringForceFunc() func(spring *Constraint, dist float64) float64 {
	return c.dampedSpringForceFunc
}

//export go_chipmunk_damped_spring_force_func
func go_chipmunk_damped_spring_force_func(cspring unsafe.Pointer, cdist C.cpFloat) float64 {
	spring := goConstraint((*C.cpConstraint)(cspring))
	return spring.dampedRotarySpringTorqueFunc(spring, float64(cdist))
}
