// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cp

/*
#include "chipmunk/include/chipmunk/chipmunk.h"

extern void pre_go_chipmunk_constraint_pre_solve_func(cpConstraint *constraint, cpSpace *space);
extern void pre_go_chipmunk_constraint_post_solve_func(cpConstraint *constraint, cpSpace *space);
*/
import "C"

import (
	"unsafe"
)

type Constraint struct {
	c                            *C.cpConstraint
	aBodyRef, bBodyRef *Body
	userData                     interface{}
	preSolveFunc, postSolveFunc  func(*Constraint, *Space)
	dampedRotarySpringTorqueFunc func(spring *Constraint, relativeAngle float64) float64
	dampedSpringForceFunc        func(spring *Constraint, dist float64) float64
}

func goConstraint(c *C.cpConstraint) *Constraint {
	data := C.cpConstraintGetUserData(c)
	return (*Constraint)(data)
}

func finalizeConstraint(c *Constraint) {
	if c.c != nil {
		c.c = nil
		C.cpConstraintFree(c.c)
	}
}

// Free is deprecated. Do not use it, it is no-op.
func (c *Constraint) Free() {
}

// Get the Space this constraint is added to.
func (c *Constraint) Space() *Space {
	return goSpace(C.cpConstraintGetSpace(c.c))
}

// Get the first body the constraint is attached to.
func (c *Constraint) BodyA() *Body {
	return goBody(C.cpConstraintGetBodyA(c.c), c.Space())
}

// Get the second body the constraint is attached to.
func (c *Constraint) BodyB() *Body {
	return goBody(C.cpConstraintGetBodyB(c.c), c.Space())
}

// Get the maximum force that this constraint is allowed to use.
func (c *Constraint) MaxForce() float64 {
	return float64(C.cpConstraintGetMaxForce(c.c))
}

// Set the maximum force that this constraint is allowed to use. (defaults to INFINITY)
func (c *Constraint) SetMaxForce(maxForce float64) {
	C.cpConstraintSetMaxForce(c.c, C.cpFloat(maxForce))
}

// Get rate at which joint error is corrected.
func (c *Constraint) ErrorBias() float64 {
	return float64(C.cpConstraintGetErrorBias(c.c))
}

// Set rate at which joint error is corrected.
//
// Defaults to pow(1.0 - 0.1, 60.0) meaning that it will
// correct 10% of the error every 1/60th of a second.
func (c *Constraint) SetErrorBias(errorBias float64) {
	C.cpConstraintSetErrorBias(c.c, C.cpFloat(errorBias))
}

// Get the maximum rate at which joint error is corrected.
func (c *Constraint) MaxBias() float64 {
	return float64(C.cpConstraintGetMaxBias(c.c))
}

// Set the maximum rate at which joint error is corrected. (defaults to INFINITY)
func (c *Constraint) SetMaxBias(maxBias float64) {
	C.cpConstraintSetMaxBias(c.c, C.cpFloat(maxBias))
}

// Get if the two bodies connected by the constraint are allowed to collide or not.
func (c *Constraint) CollideBodies() bool {
	return goBool(C.cpConstraintGetCollideBodies(c.c))
}

// Set if the two bodies connected by the constraint are allowed to collide or not. (defaults to cpFalse)
func (c *Constraint) SetCollideBodies(collideBodies bool) {
	var cbool C.cpBool = C.cpTrue
	if !collideBodies {
		cbool = C.cpFalse
	}
	C.cpConstraintSetCollideBodies(c.c, cbool)
}

// Get the user definable data pointer for this constraint
func (c *Constraint) UserData() interface{} {
	return c.userData
}

// Set the user definable data pointer for this constraint
func (c *Constraint) SetUserData(i interface{}) {
	c.userData = i
}

// Get the last impulse applied by this constraint.
func (c *Constraint) Impulse() float64 {
	return float64(C.cpConstraintGetImpulse(c.c))
}

type (
	// Callback function type that gets called before solving a joint.
	ConstraintPreSolveFunc func(c *Constraint, space *Space)

	// Callback function type that gets called after solving a joint.
	ConstraintPostSolveFunc func(c *Constraint, space *Space)
)

//export go_chipmunk_constraint_pre_solve_func
func go_chipmunk_constraint_pre_solve_func(cconstraint, cspace unsafe.Pointer) {
	constraint := goConstraint((*C.cpConstraint)(cconstraint))
	space := goSpace((*C.cpSpace)(cspace))
	constraint.PreSolveFunc()(constraint, space)
}

//export go_chipmunk_constraint_post_solve_func
func go_chipmunk_constraint_post_solve_func(cconstraint, cspace unsafe.Pointer) {
	constraint := goConstraint((*C.cpConstraint)(cconstraint))
	space := goSpace((*C.cpSpace)(cspace))
	constraint.PostSolveFunc()(constraint, space)
}

// Get the pre-solve function that is called before the solver runs.
func (c *Constraint) PreSolveFunc() func(*Constraint, *Space) {
	return c.preSolveFunc
}

// Set the pre-solve function that is called before the solver runs.
func (c *Constraint) SetPreSolveFunc(f func(*Constraint, *Space)) {
	c.preSolveFunc = f
	C.cpConstraintSetPreSolveFunc(
		c.c,
		(*[0]byte)(unsafe.Pointer(C.pre_go_chipmunk_constraint_pre_solve_func)),
	)
}

// Get the post-solve function that is called before the solver runs.
func (c *Constraint) PostSolveFunc() func(*Constraint, *Space) {
	return c.postSolveFunc
}

// Set the post-solve function that is called before the solver runs.
func (c *Constraint) SetPostSolveFunc(f func(*Constraint, *Space)) {
	c.postSolveFunc = f
	C.cpConstraintSetPostSolveFunc(
		c.c,
		(*[0]byte)(unsafe.Pointer(C.pre_go_chipmunk_constraint_post_solve_func)),
	)
}
