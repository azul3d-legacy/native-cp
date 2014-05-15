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

const (
	MAX_CONTACTS_PER_ARBITER = C.CP_MAX_CONTACTS_PER_ARBITER
)

// The Arbiter struct controls pairs of colliding shapes.
//
// They are also used in conjuction with collision handler callbacks allowing
// you to retrieve information on the collision and control it.
type Arbiter struct {
	c        *C.cpArbiter
	userData interface{}
}

func goArbiter(c *C.cpArbiter) *Arbiter {
	ptr := C.cpArbiterGetUserData(c)
	return (*Arbiter)(ptr)
}

func (a *Arbiter) Restitution() float64 {
	return float64(C.cpArbiterGetRestitution(a.c))
}

func (a *Arbiter) SetRestitution(restitution float64) {
	C.cpArbiterSetRestitution(
		a.c,
		C.cpFloat(restitution),
	)
}

func (a *Arbiter) Friction() float64 {
	return float64(C.cpArbiterGetFriction(a.c))
}

func (a *Arbiter) SetFriction(friction float64) {
	C.cpArbiterSetFriction(
		a.c,
		C.cpFloat(friction),
	)
}

// Get the relative surface velocity of the two shapes in contact.
func (a *Arbiter) SurfaceVelocity() Vect {
	ret := C.cpArbiterGetSurfaceVelocity(a.c)
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Override the relative surface velocity of the two shapes in contact.
//
// By default this is calculated to be the difference of the two surface
// velocities clamped to the tangent plane.
func (a *Arbiter) SetSurfaceVelocity(vr Vect) {
	C.cpArbiterSetSurfaceVelocity(
		a.c,
		*(*C.cpVect)(unsafe.Pointer(&vr)),
	)
}

// Calculate the total impulse including the friction that was applied by this arbiter.
// This function should only be called from a post-solve, post-step or cpBodyEachArbiter callback.
func (a *Arbiter) TotalImpulse() Vect {
	ret := C.cpArbiterTotalImpulse(a.c)
	return *(*Vect)(unsafe.Pointer(&ret))
}

func (a *Arbiter) UserData() interface{} {
	return a.userData
}

func (a *Arbiter) SetUserData(i interface{}) {
	a.userData = i
}

// Calculate the amount of energy lost in a collision including static, but not dynamic friction.
// This function should only be called from a post-solve, post-step or cpBodyEachArbiter callback.
func (a *Arbiter) TotalKE() float64 {
	return float64(C.cpArbiterTotalKE(a.c))
}

func (a *Arbiter) Ignore() bool {
	return goBool(C.cpArbiterIgnore(a.c))
}

// Return the colliding shapes involved for this arbiter.
//
// The order of their cpSpace.collision_type values will match
// the order set when the collision handler was registered.
func (arb *Arbiter) Shapes() (a, b *Shape) {
	var ca, cb *C.cpShape
	C.cpArbiterGetShapes(
		arb.c,
		&ca,
		&cb,
	)
	return goShape(ca), goShape(cb)
}

type Int C.int

type ContactPoint struct {
	// The position of the contact on the surface of each shape.
	Point1, Point2 Vect

	// Penetration distance of the two shapes. Overlapping means it will be
	// negative.
	//
	// This value is calculated as Vdot(Vsub(point2, point1), normal) and is
	// ignored by arbiter.SetContactPointSet().
	Distance float64
}

// A struct that wraps up the important collision data for an arbiter.
type ContactPointSet struct {
	// The number of contact points in the set.
	Count Int

	// The normal of the collision.
	Normal Vect

	// The array of contact points.
	Points [MAX_CONTACTS_PER_ARBITER]ContactPoint
}

// Return a contact set from an arbiter.
func (a *Arbiter) ContactPointSet() ContactPointSet {
	ret := C.cpArbiterGetContactPointSet(a.c)
	return *(*ContactPointSet)(unsafe.Pointer(&ret))
}

// Replace the contact point set for an arbiter.
//
// This can be a very powerful feature, but use it with caution!
func (a *Arbiter) SetContactPointSet(set *ContactPointSet) {
	C.cpArbiterSetContactPointSet(
		a.c,
		(*C.cpContactPointSet)(unsafe.Pointer(set)),
	)
}

// Returns true if this is the first step a pair of objects started colliding.
func (a *Arbiter) IsFirstContact() bool {
	return goBool(C.cpArbiterIsFirstContact(a.c))
}

// Returns true if in separate callback due to a shape being removed from the space.
func (a *Arbiter) IsRemoval() bool {
	return goBool(C.cpArbiterIsRemoval(a.c))
}

// Get the number of contact points for this arbiter.
func (a *Arbiter) Count() int {
	return int(C.cpArbiterGetCount(a.c))
}

// Get the normal of the collision.
func (a *Arbiter) GetNormal() Vect {
	ret := C.cpArbiterGetNormal(a.c)
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Get the position of the  ith contact point on the surface of the first shape.
func (a *Arbiter) Point1(i int) Vect {
	ret := C.cpArbiterGetPoint1(a.c, C.int(i))
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Get the position of the  ith contact point on the surface of the second shape.
func (a *Arbiter) Point2(i int) Vect {
	ret := C.cpArbiterGetPoint2(a.c, C.int(i))
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Get the depth of the  ith contact point.
func (a *Arbiter) Depth(i int) float64 {
	return float64(C.cpArbiterGetDepth(a.c, C.int(i)))
}

func (a *Arbiter) CallWildcardBeginA(space *Space) bool {
	return goBool(C.cpArbiterCallWildcardBeginA(a.c, space.c))
}

func (a *Arbiter) CallWildcardBeginB(space *Space) bool {
	return goBool(C.cpArbiterCallWildcardBeginB(a.c, space.c))
}

func (a *Arbiter) CallWildcardPreSolveA(space *Space) bool {
	return goBool(C.cpArbiterCallWildcardPreSolveA(a.c, space.c))
}

func (a *Arbiter) CallWildcardPreSolveB(space *Space) bool {
	return goBool(C.cpArbiterCallWildcardPreSolveB(a.c, space.c))
}

func (a *Arbiter) CallWildcardPostSolveA(space *Space) {
	C.cpArbiterCallWildcardPostSolveA(a.c, space.c)
}

func (a *Arbiter) CallWildcardPostSolveB(space *Space) {
	C.cpArbiterCallWildcardPostSolveB(a.c, space.c)
}

func (a *Arbiter) CallWildcardSeparateA(space *Space) {
	C.cpArbiterCallWildcardSeparateA(a.c, space.c)
}

func (a *Arbiter) CallWildcardSeparateB(space *Space) {
	C.cpArbiterCallWildcardSeparateB(a.c, space.c)
}
