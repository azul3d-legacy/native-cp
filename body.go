// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cp

/*
#include "chipmunk/include/chipmunk/chipmunk.h"

extern void pre_go_chipmunk_body_velocity_func(cpBody *body, cpVect gravity, cpFloat damping, cpFloat dt);
extern void pre_go_chipmunk_body_position_func(cpBody *body, cpFloat dt);
extern void pre_go_chipmunk_body_each_shape(cpBody *body, cpShape *shape, void *data);
extern void pre_go_chipmunk_body_each_constraint(cpBody *body, cpConstraint *constraint, void *data);
extern void pre_go_chipmunk_body_each_arbiter(cpBody *body, cpArbiter *arbiter, void *data);
*/
import "C"

import (
	"runtime"
	"unsafe"
)

// Chipmunk's rigid body type. Rigid bodies hold the physical properties of an object like
// it's mass, and position and velocity of it's center of gravity. They don't have an shape on their own.
// They are given a shape by creating collision shapes (cpShape) that point to the body.
type BodyType int

const (
	BODY_TYPE_DYNAMIC   BodyType = C.CP_BODY_TYPE_DYNAMIC
	BODY_TYPE_KINEMATIC BodyType = C.CP_BODY_TYPE_KINEMATIC
	BODY_TYPE_STATIC    BodyType = C.CP_BODY_TYPE_STATIC
)

type (
	// Rigid body velocity update function type.
	BodyVelocityFunc func(body *Body, gravity Vect, damping, dt float64)

	// Rigid body position update function type.
	BodyPositionFunc func(body *Body, dt float64)
)

type Body struct {
	c                *C.cpBody
	userData         interface{}
	bodyVelocityFunc BodyVelocityFunc
	bodyPositionFunc BodyPositionFunc
	spaceRef *Space
}

func goBody(c *C.cpBody, optional *Space) *Body {
	data := C.cpBodyGetUserData(c)
	g := (*Body)(data)
	g.spaceRef = optional
	return g
}

//export go_chipmunk_body_velocity_func
func go_chipmunk_body_velocity_func(cbody unsafe.Pointer, gravity C.cpVect, damping C.cpFloat, dt C.cpFloat) {
	b := goBody((*C.cpBody)(unsafe.Pointer(cbody)), nil)
	b.bodyVelocityFunc(b, goVect(gravity), float64(damping), float64(dt))
}

//export go_chipmunk_body_position_func
func go_chipmunk_body_position_func(cbody unsafe.Pointer, dt C.cpFloat) {
	b := goBody((*C.cpBody)(unsafe.Pointer(cbody)), nil)
	b.bodyPositionFunc(b, float64(dt))
}

// Allocate and initialize a Body.
func BodyNew(mass, moment float64) *Body {
	b := new(Body)
	b.c = C.cpBodyNew(C.cpFloat(mass), C.cpFloat(moment))
	if b.c == nil {
		return nil
	}
	C.cpBodySetUserData(b.c, C.cpDataPointer(unsafe.Pointer(b)))
	runtime.SetFinalizer(b, finalizeBody)
	return b
}

// Allocate and initialize a Body, and set it as a kinematic body.
func BodyNewKinematic() *Body {
	b := new(Body)
	b.c = C.cpBodyNewKinematic()
	if b.c == nil {
		return nil
	}
	C.cpBodySetUserData(b.c, C.cpDataPointer(unsafe.Pointer(b)))
	runtime.SetFinalizer(b, finalizeBody)
	return b
}

// Allocate and initialize a cpBody, and set it as a static body.
func BodyNewStatic() *Body {
	b := new(Body)
	b.c = C.cpBodyNewStatic()
	if b.c == nil {
		return nil
	}
	C.cpBodySetUserData(b.c, C.cpDataPointer(unsafe.Pointer(b)))
	runtime.SetFinalizer(b, finalizeBody)
	return b
}

func finalizeBody(b *Body) {
	if b.c != nil {
		b.c = nil
		C.cpBodyFree(b.c)
	}
}

// Free is deprecated. Do not use it, it is no-op.
func (b *Body) Free() {
}

// Wake up a sleeping or idle body.
func (b *Body) Activate() {
	C.cpBodyActivate(b.c)
}

// Wake up any sleeping or idle bodies touching a static body.
func (b *Body) ActivateStatic(filter *Shape) {
	C.cpBodyActivateStatic(
		b.c,
		(*C.cpShape)(unsafe.Pointer(filter)),
	)
}

// Force a body to fall asleep immediately.
func (b *Body) Sleep() {
	C.cpBodySleep(b.c)
}

// Force a body to fall asleep immediately along with other bodies in a group.
func (b *Body) SleepWithGroup(group *Body) {
	C.cpBodySleepWithGroup(
		b.c,
		group.c,
	)
}

// Returns true if the body is sleeping.
func (b *Body) IsSleeping() bool {
	return goBool(C.cpBodyIsSleeping(b.c))
}

// Get the type of the body.
func (b *Body) Type() BodyType {
	return BodyType(C.cpBodyGetType(b.c))
}

// Set the type of the body.
func (b *Body) SetType(t BodyType) {
	C.cpBodySetType(
		b.c,
		C.cpBodyType(t),
	)
}

// Get the space this body is added to.
func (b *Body) Space() *Space {
	return goSpace(C.cpBodyGetSpace(b.c))
}

// Get the mass of the body.
func (b *Body) Mass() float64 {
	return float64(C.cpBodyGetMass(b.c))
}

// Set the mass of the body.
func (b *Body) SetMass(m float64) {
	C.cpBodySetMass(
		b.c,
		C.cpFloat(m),
	)
}

// Get the moment of inertia of the body.
func (b *Body) Moment() float64 {
	return float64(C.cpBodyGetMoment(b.c))
}

// Set the moment of inertia of the body.
func (b *Body) SetMoment(i float64) {
	C.cpBodySetMoment(
		b.c,
		C.cpFloat(i),
	)
}

// Get the position of a body.
func (b *Body) Position() Vect {
	return goVect(C.cpBodyGetPosition(b.c))
}

// Set the position of the body.
func (b *Body) SetPosition(pos Vect) {
	C.cpBodySetPosition(b.c, pos.c())
}

// Get the offset of the center of gravity in body local coordinates.
func (b *Body) CenterOfGravity() Vect {
	return goVect(C.cpBodyGetCenterOfGravity(b.c))
}

// Set the offset of the center of gravity in body local coordinates.
func (b *Body) SetCenterOfGravity(cog Vect) {
	C.cpBodySetCenterOfGravity(b.c, cog.c())
}

// Get the velocity of the body.
func (b *Body) Velocity() Vect {
	return goVect(C.cpBodyGetVelocity(b.c))
}

// Set the velocity of the body.
func (b *Body) SetVelocity(velocity Vect) {
	C.cpBodySetVelocity(b.c, velocity.c())
}

// Get the force applied to the body for the next time step.
func (b *Body) Force() Vect {
	return goVect(C.cpBodyGetForce(b.c))
}

// Set the force applied to the body for the next time step.
func (b *Body) SetForce(force Vect) {
	C.cpBodySetForce(b.c, force.c())
}

// Get the angle of the body.
func (b *Body) Angle() float64 {
	return float64(C.cpBodyGetAngle(b.c))
}

// Set the angle of a body.
func (b *Body) SetAngle(a float64) {
	C.cpBodySetAngle(
		b.c,
		C.cpFloat(a),
	)
}

// Get the angular velocity of the body.
func (b *Body) AngularVelocity() float64 {
	return float64(C.cpBodyGetAngularVelocity(b.c))
}

// Set the angular velocity of the body.
func (b *Body) SetAngularVelocity(angularVelocity float64) {
	C.cpBodySetAngularVelocity(
		b.c,
		C.cpFloat(angularVelocity),
	)
}

// Get the torque applied to the body for the next time step.
func (b *Body) Torque() float64 {
	return float64(C.cpBodyGetTorque(
		b.c,
	))
}

// Set the torque applied to the body for the next time step.
func (b *Body) SetTorque(torque float64) {
	C.cpBodySetTorque(
		b.c,
		C.cpFloat(torque),
	)
}

// Get the rotation vector of the body. (The x basis vector of it's transform.)
func (b *Body) Rotation() Vect {
	return goVect(C.cpBodyGetRotation(b.c))
}

// Get the user data interface assigned to the body.
func (b *Body) UserData() interface{} {
	return b.userData
}

// Set the user data interface assigned to the body.
func (b *Body) SetUserData(userData interface{}) {
	b.userData = userData
}

// Set the callback used to update a body's velocity.
func (b *Body) SetVelocityUpdateFunc(f BodyVelocityFunc) {
	old := b.bodyVelocityFunc
	b.bodyVelocityFunc = f
	if old == nil {
		C.cpBodySetVelocityUpdateFunc(
			b.c,
			(*[0]byte)(unsafe.Pointer(C.pre_go_chipmunk_body_velocity_func)),
		)
	}
}

// Set the callback used to update a body's position.
//
// NOTE: It's not generally recommended to override this.
func (b *Body) SetPositionUpdateFunc(f BodyPositionFunc) {
	old := b.bodyPositionFunc
	b.bodyPositionFunc = f
	if old == nil {
		C.cpBodySetPositionUpdateFunc(
			b.c,
			(*[0]byte)(unsafe.Pointer(C.pre_go_chipmunk_body_position_func)),
		)
	}
}

// Default velocity integration function..
func BodyUpdateVelocity(b *Body, gravity Vect, damping, dt float64) {
	C.cpBodyUpdateVelocity(b.c, gravity.c(), C.cpFloat(damping), C.cpFloat(dt))
}

// Default position integration function.
func BodyUpdatePosition(b *Body, dt float64) {
	C.cpBodyUpdatePosition(b.c, C.cpFloat(dt))
}

// Convert body relative/local coordinates to absolute/world coordinates.
func (b *Body) LocalToWorld(point Vect) Vect {
	return goVect(C.cpBodyLocalToWorld(b.c, point.c()))
}

// Convert body absolute/world coordinates to  relative/local coordinates.
func (b *Body) WorldToLocal(point Vect) Vect {
	return goVect(C.cpBodyWorldToLocal(b.c, point.c()))
}

// Apply a force to a body. Both the force and point are expressed in world coordinates.
func (b *Body) ApplyForceAtWorldPoint(force, point Vect) {
	C.cpBodyApplyForceAtWorldPoint(b.c, force.c(), point.c())
}

// Apply a force to a body. Both the force and point are expressed in body local coordinates.
func (b *Body) ApplyForceAtLocalPoint(force, point Vect) {
	C.cpBodyApplyForceAtLocalPoint(b.c, force.c(), point.c())
}

// Apply an impulse to a body. Both the impulse and point are expressed in world coordinates.
func (b *Body) ApplyImpulseAtWorldPoint(impulse, point Vect) {
	C.cpBodyApplyImpulseAtWorldPoint(b.c, impulse.c(), point.c())
}

// Apply an impulse to a body. Both the impulse and point are expressed in body local coordinates.
func (b *Body) ApplyImpulseAtLocalPoint(impulse, point Vect) {
	C.cpBodyApplyImpulseAtLocalPoint(b.c, impulse.c(), point.c())
}

// Get the velocity on a body (in world units) at a point on the body in world coordinates.
func (b *Body) VelocityAtWorldPoint(point Vect) Vect {
	return goVect(C.cpBodyGetVelocityAtWorldPoint(b.c, point.c()))
}

// Get the velocity on a body (in world units) at a point on the body in local coordinates.
func (b *Body) VelocityAtLocalPoint(point Vect) Vect {
	return goVect(C.cpBodyGetVelocityAtLocalPoint(b.c, point.c()))
}

// Get the amount of kinetic energy contained by the body.
func (b *Body) KineticEnergy() float64 {
	return float64(C.cpBodyKineticEnergy(b.c))
}

//export go_chipmunk_body_each_shape
func go_chipmunk_body_each_shape(cbody, cshape, data unsafe.Pointer) {
	body := goBody((*C.cpBody)(cbody), nil)
	shape := goShape((*C.cpShape)(cshape))
	f := *(*func(b *Body, s *Shape))(data)
	f(body, shape)
}

// Returns a slice of all shapes attached to the body and added to the space.
func (b *Body) EachShape(f func(b *Body, s *Shape)) {
	C.cpBodyEachShape(
		b.c,
		(*[0]byte)(C.pre_go_chipmunk_body_each_shape),
		unsafe.Pointer(&f),
	)
}

type bodyEachConstraintData struct {
	cb func(b *Body, c *Constraint)
	*Space
}

//export go_chipmunk_body_each_constraint
func go_chipmunk_body_each_constraint(cbody, cconstraint, cdata unsafe.Pointer) {
	data := (*bodyEachConstraintData)(cdata)
	body := goBody((*C.cpBody)(cbody), data.Space)
	constraint := goConstraint((*C.cpConstraint)(cconstraint), data.Space)
	data.cb(body, constraint)
}

// Returns a slice of all contraints attached to the body and added to the space.
func (b *Body) EachConstraint(f func(b *Body, c *Constraint)) {
	C.cpBodyEachConstraint(
		b.c,
		(*[0]byte)(C.pre_go_chipmunk_body_each_constraint),
		unsafe.Pointer(&bodyEachConstraintData{
			cb: f,
			Space: b.Space(),
		}),
	)
}

//export go_chipmunk_body_each_arbiter
func go_chipmunk_body_each_arbiter(cbody, carbiter, data unsafe.Pointer) {
	body := goBody((*C.cpBody)(cbody), nil)
	arbiter := goArbiter((*C.cpArbiter)(carbiter))
	f := *(*func(b *Body, c *Arbiter))(data)
	f(body, arbiter)
}

// Returns a slice of all arbiters that are currently active on the body.
func (b *Body) EachArbiter(f func(b *Body, a *Arbiter)) {
	C.cpBodyEachArbiter(
		b.c,
		(*[0]byte)(C.pre_go_chipmunk_body_each_arbiter),
		unsafe.Pointer(&f),
	)
}
