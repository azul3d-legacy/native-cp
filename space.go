// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cp

/*
#include "chipmunk/include/chipmunk/chipmunk.h"

extern void pre_go_chipmunk_space_point_query_func(cpShape *shape, cpVect point, cpFloat distance, cpVect gradient, void *data);
extern void pre_go_chipmunk_space_segment_query_func(cpShape *shape, cpVect point, cpVect normal, cpFloat alpha, void *data);
extern void pre_go_chipmunk_space_bb_query_func(cpShape *shape, void *data);
extern void pre_go_chipmunk_space_shape_query_func(cpShape *shape, cpContactPointSet* points, void *data);
extern void pre_go_chipmunk_space_body_iterator_func(cpBody* body, void *data);
extern void pre_go_chipmunk_space_shape_iterator_func(cpShape* shape, void *data);
extern void pre_go_chipmunk_space_constraint_iterator_func(cpConstraint* constraint, void *data);

extern void pre_go_chipmunk_space_debug_draw_circle_impl(cpVect pos, cpFloat angle, cpFloat radius, cpSpaceDebugColor outlineColor, cpSpaceDebugColor fillColor, cpDataPointer *data);
extern void pre_go_chipmunk_space_debug_draw_segment_impl(cpVect a, cpVect b, cpSpaceDebugColor color, cpDataPointer *data);
extern void pre_go_chipmunk_space_debug_draw_fat_segment_impl(cpVect a, cpVect b, cpFloat radius, cpSpaceDebugColor outlineColor, cpSpaceDebugColor fillColor, cpDataPointer *data);
extern void pre_go_chipmunk_space_debug_draw_polygon_impl(int count, const cpVect *verts, cpFloat radius, cpSpaceDebugColor outlineColor, cpSpaceDebugColor fillColor, cpDataPointer *data);
extern void pre_go_chipmunk_space_debug_draw_dot_impl(cpFloat size, cpVect pos, cpSpaceDebugColor color, cpDataPointer *data);
extern cpSpaceDebugColor pre_go_chipmunk_space_debug_draw_color_for_shape_impl(cpShape *shape, cpDataPointer *data);

extern cpBool pre_go_chipmunk_collision_begin_func(cpArbiter *arb, cpSpace *space, cpDataPointer userData);
extern cpBool pre_go_chipmunk_collision_pre_solve_func(cpArbiter *arb, cpSpace *space, cpDataPointer userData);
extern void pre_go_chipmunk_collision_post_solve_func(cpArbiter *arb, cpSpace *space, cpDataPointer userData);
extern void pre_go_chipmunk_collision_separate_func(cpArbiter *arb, cpSpace *space, cpDataPointer userData);
*/
import "C"

import (
	"runtime"
	"reflect"
	"unsafe"
)

type (
	SpaceArbiterApplyImpulseFunc func(arb *Arbiter)

	// Collision begin event function callback type.
	//
	// Returning false from a begin callback causes the collision to be ignored
	// until the the separate callback is called when the objects stop
	// colliding.
	CollisionBeginFunc func(arb *Arbiter, space *Space, userData interface{}) bool

	// Collision pre-solve event function callback type.
	//
	// Returning false from a pre-step callback causes the collision to be
	// ignored until the next step.
	CollisionPreSolveFunc func(arb *Arbiter, space *Space, userData interface{}) bool

	// Collision post-solve event function callback type.
	CollisionPostSolveFunc func(arb *Arbiter, space *Space, userData interface{})

	// Collision separate event function callback type.
	CollisionSeparateFunc func(arb *Arbiter, space *Space, userData interface{})

	CollisionHandler struct {
		TypeA, TypeB  CollisionType
		BeginFunc     CollisionBeginFunc
		PreSolveFunc  CollisionPreSolveFunc
		PostSolveFunc CollisionPostSolveFunc
		SeparateFunc  CollisionSeparateFunc
		UserData      interface{}
	}
)

type Space struct {
	c                 *C.cpSpace
	userData          interface{}
	postStepCallbacks []func()
}

func goSpace(c *C.cpSpace) *Space {
	data := C.cpSpaceGetUserData(c)
	return (*Space)(data)
}

// Allocate and initialize a cpSpace.
func SpaceNew() *Space {
	s := new(Space)
	s.c = C.cpSpaceNew()
	if s.c == nil {
		return nil
	}
	C.cpSpaceSetUserData(s.c, C.cpDataPointer(unsafe.Pointer(s)))
	runtime.SetFinalizer(s, finalizeSpace)
	return s
}

func finalizeSpace(s *Space) {
	if s.c != nil {
		s.c = nil
		C.cpSpaceFree(s.c)
	}
}

// Free is deprecated. Do not use it, it is no-op.
func (s *Space) Free() {
}

// Number of iterations to use in the impulse solver to solve contacts and
// other constraints.
func (s *Space) Iterations() int {
	return int(C.cpSpaceGetIterations(s.c))
}

func (s *Space) SetIterations(iterations int) {
	C.cpSpaceSetIterations(s.c, C.int(iterations))
}

// Gravity to pass to rigid bodies when integrating velocity.
func (s *Space) Gravity() Vect {
	ret := C.cpSpaceGetGravity(s.c)
	return *(*Vect)(unsafe.Pointer(&ret))
}

func (s *Space) SetGravity(gravity Vect) {
	C.cpSpaceSetGravity(
		s.c,
		*(*C.cpVect)(unsafe.Pointer(&gravity)),
	)
}

// Damping rate expressed as the fraction of velocity bodies retain each second.
//
// A value of 0.9 would mean that each body's velocity will drop 10% per second.
//
// The default value is 1.0, meaning no damping is applied.
//
// This damping value is different than those of cpDampedSpring and cpDampedRotarySpring.
func (s *Space) Damping() float64 {
	return float64(C.cpSpaceGetDamping(s.c))
}

func (s *Space) SetDamping(damping float64) {
	C.cpSpaceSetDamping(s.c, C.cpFloat(damping))
}

// Speed threshold for a body to be considered idle.
//
// The default value of 0 means to let the space guess a good threshold based
// on gravity.
func (s *Space) IdleSpeedThreshold() float64 {
	return float64(C.cpSpaceGetIdleSpeedThreshold(s.c))
}

func (s *Space) SetIdleSpeedThreshold(idleSpeedThreshold float64) {
	C.cpSpaceSetIdleSpeedThreshold(s.c, C.cpFloat(idleSpeedThreshold))
}

// Time a group of bodies must remain idle in order to fall asleep.
//
// Enabling sleeping also implicitly enables the the contact graph.
//
// The default value of INFINITY disables the sleeping algorithm.
func (s *Space) SleepTimeThreshold() float64 {
	return float64(C.cpSpaceGetSleepTimeThreshold(s.c))
}

func (s *Space) SetSleepTimeThreshold(sleepTimeThreshold float64) {
	C.cpSpaceSetSleepTimeThreshold(s.c, C.cpFloat(sleepTimeThreshold))
}

// Amount of encouraged penetration between colliding shapes.
//
// Used to reduce oscillating contacts and keep the collision cache warm.
//
// Defaults to 0.1. If you have poor simulation quality,
// increase this number as much as possible without allowing visible amounts of overlap.
func (s *Space) CollisionSlop() float64 {
	return float64(C.cpSpaceGetCollisionSlop(s.c))
}

func (s *Space) SetCollisionSlop(collisionSlop float64) {
	C.cpSpaceSetCollisionSlop(s.c, C.cpFloat(collisionSlop))
}

// Determines how fast overlapping shapes are pushed apart.
// Expressed as a fraction of the error remaining after each second.
// Defaults to pow(1.0 - 0.1, 60.0) meaning that Chipmunk fixes 10% of overlap each frame at 60Hz.
func (s *Space) CollisionBias() float64 {
	return float64(C.cpSpaceGetCollisionBias(s.c))
}

func (s *Space) SetCollisionBias(collisionBias float64) {
	C.cpSpaceSetCollisionBias(s.c, C.cpFloat(collisionBias))
}

// Number of frames that contact information should persist.
//
// Defaults to 3. There is probably never a reason to change this value.
func (s *Space) CollisionPersistence() Timestamp {
	return Timestamp(C.cpSpaceGetCollisionPersistence(s.c))
}

func (s *Space) SetCollisionPersistence(collisionPersistence Timestamp) {
	C.cpSpaceSetCollisionPersistence(
		s.c,
		*(*C.cpTimestamp)(unsafe.Pointer(&collisionPersistence)),
	)
}

// User definable data interface.
//
// Generally this points to your game's controller or game state
// class so you can access it when given a cpSpace reference in a callback.
func (s *Space) UserData() interface{} {
	return s.userData
}

func (s *Space) SetUserData(i interface{}) {
	s.userData = i
}

// The Space provided static body for a given cpSpace.
//
// This is merely provided for convenience and you are not required to use it.
func (s *Space) StaticBody() *Body {
	b := new(Body)
	b.c = C.cpSpaceGetStaticBody(s.c)
	if b.c == nil {
		return nil
	}
	C.cpBodySetUserData(b.c, C.cpDataPointer(unsafe.Pointer(b)))
	return b
}

// Returns the current (or most recent) time step used with the given space.
// Useful from callbacks if your time step is not a compile-time global.
func (s *Space) CurrentTimeStep() float64 {
	return float64(C.cpSpaceGetCurrentTimeStep(s.c))
}

// returns true from inside a callback when objects cannot be added/removed.
func (s *Space) IsLocked() bool {
	return goBool(C.cpSpaceIsLocked(s.c))
}

// Add a collision shape to the simulation.
//
// If the shape is attached to a static body, it will be added as a static shape.
func (s *Space) AddShape(shape *Shape) *Shape {
	return goShape(C.cpSpaceAddShape(
		s.c,
		shape.c,
	))
}

// Add a rigid body to the simulation.
func (s *Space) AddBody(body *Body) *Body {
	return goBody(C.cpSpaceAddBody(
		s.c,
		body.c,
	))
}

// Add a constraint to the simulation.
func (s *Space) AddConstraint(constraint *Constraint) *Constraint {
	return goConstraint(C.cpSpaceAddConstraint(
		s.c,
		constraint.c,
	))
}

// Remove a collision shape from the simulation.
func (s *Space) RemoveShape(shape *Shape) {
	C.cpSpaceRemoveShape(s.c, shape.c)
}

// Remove a rigid body from the simulation.
func (s *Space) RemoveBody(body *Body) {
	C.cpSpaceRemoveBody(s.c, body.c)
}

// Remove a constraint from the simulation.
func (s *Space) RemoveConstraint(constraint *Constraint) {
	C.cpSpaceRemoveConstraint(s.c, constraint.c)
}

// Test if a collision shape has been added to the space.
func (s *Space) ContainsShape(shape *Shape) bool {
	return goBool(C.cpSpaceContainsShape(s.c, shape.c))
}

// Test if a rigid body has been added to the space.
func (s *Space) ContainsBody(body *Body) bool {
	return goBool(C.cpSpaceContainsBody(s.c, body.c))
}

// Test if a constraint has been added to the space.
func (s *Space) ContainsConstraint(constraint *Constraint) bool {
	return goBool(C.cpSpaceContainsConstraint(s.c, constraint.c))
}

// Schedule a post-step callback to be called when space.Step() finishes.
func (s *Space) AddPostStepCallback(f func()) {
	s.postStepCallbacks = append(s.postStepCallbacks, f)
}

// Step the space forward in time by dt.
func (s *Space) Step(dt float64) {
	C.cpSpaceStep(s.c, C.cpFloat(dt))
	for _, f := range s.postStepCallbacks {
		f()
	}
	s.postStepCallbacks = s.postStepCallbacks[:0]
}

func (s *Space) AddDefaultCollisionHandler() {
	C.cpSpaceAddDefaultCollisionHandler(s.c)
}

//export pre_go_chipmunk_collision_begin_func
func pre_go_chipmunk_collision_begin_func(arb *C.cpArbiter, space *C.cpSpace, userData C.cpDataPointer) C.cpBool {
	handler := (*CollisionHandler)(unsafe.Pointer(userData))
	if handler != nil && handler.BeginFunc != nil {
		if handler.BeginFunc(
			goArbiter(arb),
			goSpace(space),
			handler.UserData,
		) {
			return C.cpBool(1)
		} else {
			return C.cpBool(0)
		}
	}
	return C.cpBool(1)
}

//export pre_go_chipmunk_collision_pre_solve_func
func pre_go_chipmunk_collision_pre_solve_func(arb *C.cpArbiter, space *C.cpSpace, userData C.cpDataPointer) C.cpBool {
	handler := (*CollisionHandler)(unsafe.Pointer(userData))
	if handler != nil && handler.PreSolveFunc != nil {
		if handler.PreSolveFunc(
			goArbiter(arb),
			goSpace(space),
			handler.UserData,
		) {
			return C.cpBool(1)
		} else {
			return C.cpBool(0)
		}
	}
	return C.cpBool(1)
}

//export pre_go_chipmunk_collision_post_solve_func
func pre_go_chipmunk_collision_post_solve_func(arb *C.cpArbiter, space *C.cpSpace, userData C.cpDataPointer) {
	handler := (*CollisionHandler)(unsafe.Pointer(userData))
	if handler != nil && handler.PostSolveFunc != nil {
		handler.PostSolveFunc(
			goArbiter(arb),
			goSpace(space),
			handler.UserData,
		)
	}
}

//export pre_go_chipmunk_collision_separate_func
func pre_go_chipmunk_collision_separate_func(arb *C.cpArbiter, space *C.cpSpace, userData C.cpDataPointer) {
	handler := (*CollisionHandler)(unsafe.Pointer(userData))
	if handler != nil && handler.SeparateFunc != nil {
		handler.SeparateFunc(
			goArbiter(arb),
			goSpace(space),
			handler.UserData,
		)
	}
}

func (s *Space) AddCollisionHandler(a, b CollisionType, handler *CollisionHandler) {
	cHandler := C.cpSpaceAddCollisionHandler(
		s.c,
		C.cpCollisionType(a),
		C.cpCollisionType(b),
	)
	cHandler.beginFunc = (C.cpCollisionBeginFunc)(C.pre_go_chipmunk_collision_begin_func)
	cHandler.preSolveFunc = (C.cpCollisionPreSolveFunc)(C.pre_go_chipmunk_collision_pre_solve_func)
	cHandler.postSolveFunc = (C.cpCollisionPostSolveFunc)(C.pre_go_chipmunk_collision_post_solve_func)
	cHandler.separateFunc = (C.cpCollisionSeparateFunc)(C.pre_go_chipmunk_collision_separate_func)
	if handler != nil {
		cHandler.userData = (C.cpDataPointer)(unsafe.Pointer(&handler))
	}
}

func (s *Space) AddWildcardHandler(t CollisionType, handler *CollisionHandler) {
	cHandler := C.cpSpaceAddWildcardHandler(
		s.c,
		C.cpCollisionType(t),
	)
	cHandler.beginFunc = (C.cpCollisionBeginFunc)(C.pre_go_chipmunk_collision_begin_func)
	cHandler.preSolveFunc = (C.cpCollisionPreSolveFunc)(C.pre_go_chipmunk_collision_pre_solve_func)
	cHandler.postSolveFunc = (C.cpCollisionPostSolveFunc)(C.pre_go_chipmunk_collision_post_solve_func)
	cHandler.separateFunc = (C.cpCollisionSeparateFunc)(C.pre_go_chipmunk_collision_separate_func)
	if handler != nil {
		cHandler.userData = (C.cpDataPointer)(unsafe.Pointer(&handler))
	}
}

// Nearest point query callback function type.
type SpacePointQueryFunc func(shape *Shape, point Vect, distance float64, gradient Vect, data interface{})

type spacePointQueryPair struct {
	f    SpacePointQueryFunc
	data interface{}
}

//export go_chipmunk_space_point_query_func
func go_chipmunk_space_point_query_func(cshape unsafe.Pointer, cpoint C.cpVect, cdist C.cpFloat, cgradient C.cpVect, data unsafe.Pointer) {
	shape := goShape((*C.cpShape)(cshape))
	point := *(*Vect)(unsafe.Pointer(&cpoint))
	dist := *(*float64)(unsafe.Pointer(&cdist))
	gradient := *(*Vect)(unsafe.Pointer(&cgradient))
	pair := (*spacePointQueryPair)(data)
	pair.f(shape, point, dist, gradient, pair.data)
}

// Query the space at a point and call f for each shape found.
func (s *Space) PointQuery(point Vect, maxDistance float64, filter ShapeFilter, f SpacePointQueryFunc, data interface{}) {
	pair := &spacePointQueryPair{f, data}
	C.cpSpacePointQuery(
		s.c,
		*(*C.cpVect)(unsafe.Pointer(&point)),
		C.cpFloat(maxDistance),
		*(*C.cpShapeFilter)(unsafe.Pointer(&filter)),
		(*[0]byte)(unsafe.Pointer(C.pre_go_chipmunk_space_point_query_func)),
		unsafe.Pointer(pair),
	)
}

// Query the space at a point and return the nearest shape found. Returns NULL if no shapes were found.
func (s *Space) PointQueryNearest(point Vect, maxDistance float64, filter ShapeFilter) (shape *Shape, out *PointQueryInfo) {
	out = new(PointQueryInfo)
	shape = goShape(C.cpSpacePointQueryNearest(
		s.c,
		*(*C.cpVect)(unsafe.Pointer(&point)),
		C.cpFloat(maxDistance),
		*(*C.cpShapeFilter)(unsafe.Pointer(&filter)),
		(*C.cpPointQueryInfo)(unsafe.Pointer(out)),
	))
	return
}

// Nearest point query callback function type.
type SpaceSegmentQueryFunc func(shape *Shape, point, normal Vect, alpha float64, data interface{})

type spaceSegmentQueryPair struct {
	f    SpaceSegmentQueryFunc
	data interface{}
}

//export go_chipmunk_space_segment_query_func
func go_chipmunk_space_segment_query_func(cshape unsafe.Pointer, cpoint, cnormal C.cpVect, calpha C.cpFloat, data unsafe.Pointer) {
	shape := goShape((*C.cpShape)(cshape))
	point := *(*Vect)(unsafe.Pointer(&cpoint))
	normal := *(*Vect)(unsafe.Pointer(&cnormal))
	alpha := *(*float64)(unsafe.Pointer(&calpha))
	pair := (*spaceSegmentQueryPair)(data)
	pair.f(shape, point, normal, alpha, pair.data)
}

// Perform a directed line segment query (like a raycast) against the space
// calling f for each shape intersected.
func (s *Space) SegmentQuery(start, end Vect, radius float64, filter ShapeFilter, f SpaceSegmentQueryFunc, data interface{}) {
	pair := &spaceSegmentQueryPair{f, data}
	C.cpSpaceSegmentQuery(
		s.c,
		*(*C.cpVect)(unsafe.Pointer(&start)),
		*(*C.cpVect)(unsafe.Pointer(&end)),
		C.cpFloat(radius),
		*(*C.cpShapeFilter)(unsafe.Pointer(&filter)),
		(*[0]byte)(unsafe.Pointer(C.pre_go_chipmunk_space_segment_query_func)),
		unsafe.Pointer(pair),
	)
}

// Perform a directed line segment query (like a raycast) against the space and return the first shape hit. Returns NULL if no shapes were hit.
func (s *Space) SegmentQueryFirst(start, end Vect, radius float64, filter ShapeFilter) (shape *Shape, out *SegmentQueryInfo) {
	out = new(SegmentQueryInfo)
	shape = goShape(C.cpSpaceSegmentQueryFirst(
		s.c,
		*(*C.cpVect)(unsafe.Pointer(&start)),
		*(*C.cpVect)(unsafe.Pointer(&end)),
		C.cpFloat(radius),
		*(*C.cpShapeFilter)(unsafe.Pointer(&filter)),
		(*C.cpSegmentQueryInfo)(unsafe.Pointer(out)),
	))
	return
}

// Rectangle Query callback function type.
type SpaceBBQueryFunc func(shape *Shape, data interface{})

type spaceBBQueryPair struct {
	f    SpaceBBQueryFunc
	data interface{}
}

//export go_chipmunk_space_bb_query_func
func go_chipmunk_space_bb_query_func(cshape unsafe.Pointer, data unsafe.Pointer) {
	shape := goShape((*C.cpShape)(cshape))
	pair := (*spaceBBQueryPair)(data)
	pair.f(shape, pair.data)
}

// Perform a fast rectangle query on the space calling  func for each shape found.
//
// Only the shape's bounding boxes are checked for overlap, not their full shape.
func (s *Space) BBQuery(bb BB, filter ShapeFilter, f SpaceBBQueryFunc, data interface{}) {
	pair := &spaceBBQueryPair{f, data}
	C.cpSpaceBBQuery(
		s.c,
		*(*C.cpBB)(unsafe.Pointer(&bb)),
		*(*C.cpShapeFilter)(unsafe.Pointer(&filter)),
		(*[0]byte)(unsafe.Pointer(C.pre_go_chipmunk_space_bb_query_func)),
		unsafe.Pointer(pair),
	)
}

// Shape query callback function type.
type SpaceShapeQueryFunc func(shape *Shape, points *ContactPointSet, data interface{})

type spaceShapeQueryPair struct {
	f    SpaceShapeQueryFunc
	data interface{}
}

//export go_chipmunk_space_shape_query_func
func go_chipmunk_space_shape_query_func(cshape, cpoints unsafe.Pointer, data unsafe.Pointer) {
	shape := goShape((*C.cpShape)(cshape))
	points := (*ContactPointSet)(cpoints)
	pair := (*spaceShapeQueryPair)(data)
	pair.f(shape, points, pair.data)
}

// Query a space for any shapes overlapping the given shape and call  func for each shape found.
func (s *Space) ShapeQuery(shape *Shape, f SpaceShapeQueryFunc, data interface{}) bool {
	pair := &spaceShapeQueryPair{f, data}
	return goBool(C.cpSpaceShapeQuery(
		s.c,
		shape.c,
		(*[0]byte)(unsafe.Pointer(C.pre_go_chipmunk_space_shape_query_func)),
		unsafe.Pointer(pair),
	))
}

//export go_chipmunk_space_body_iterator_func
func go_chipmunk_space_body_iterator_func(cbody, data unsafe.Pointer) {
	body := goBody((*C.cpBody)(cbody))
	f := *(*func(*Body))(data)
	f(body)
}

// Space/body iterator callback function type.
func (s *Space) EachBody(space *Space, f func(b *Body)) {
	C.cpSpaceEachBody(
		s.c,
		(*[0]byte)(unsafe.Pointer(C.pre_go_chipmunk_space_body_iterator_func)),
		unsafe.Pointer(&f),
	)
}

//export go_chipmunk_space_shape_iterator_func
func go_chipmunk_space_shape_iterator_func(cshape, data unsafe.Pointer) {
	shape := goShape((*C.cpShape)(cshape))
	f := *(*func(*Shape))(data)
	f(shape)
}

// Call f for each shape in the space.
func (s *Space) EachShape(space *Space, f func(s *Shape)) {
	C.cpSpaceEachShape(
		s.c,
		(*[0]byte)(unsafe.Pointer(C.pre_go_chipmunk_space_shape_iterator_func)),
		unsafe.Pointer(&f),
	)
}

//export go_chipmunk_space_constraint_iterator_func
func go_chipmunk_space_constraint_iterator_func(cconstraint, data unsafe.Pointer) {
	constraint := goConstraint((*C.cpConstraint)(cconstraint))
	f := *(*func(*Constraint))(data)
	f(constraint)
}

// Call f for each shape in the space.
func (s *Space) EachConstraint(space *Space, f func(c *Constraint)) {
	C.cpSpaceEachConstraint(
		s.c,
		(*[0]byte)(unsafe.Pointer(C.pre_go_chipmunk_space_constraint_iterator_func)),
		unsafe.Pointer(&f),
	)
}

// Update the collision detection info for the static shapes in the space.
func (s *Space) ReindexStatic() {
	C.cpSpaceReindexStatic(s.c)
}

// Update the collision detection data for a specific shape in the space.
func (s *Space) ReindexShape(shape *Shape) {
	C.cpSpaceReindexShape(s.c, shape.c)
}

// Update the collision detection data for all shapes attached to a body.
func (s *Space) ReindexShapesForBody(body *Body) {
	C.cpSpaceReindexShapesForBody(s.c, body.c)
}

// Switch the space to use a spatial has as it's spatial index.
func (s *Space) UseSpatialHash(dim float64, count int) {
	C.cpSpaceUseSpatialHash(
		s.c,
		C.cpFloat(dim),
		C.int(count),
	)
}

type SpaceDebugDrawFlags int

const (
	SPACE_DEBUG_DRAW_SHAPES           SpaceDebugDrawFlags = C.CP_SPACE_DEBUG_DRAW_SHAPES
	SPACE_DEBUG_DRAW_CONSTRAINTS      SpaceDebugDrawFlags = C.CP_SPACE_DEBUG_DRAW_CONSTRAINTS
	SPACE_DEBUG_DRAW_COLLISION_POINTS SpaceDebugDrawFlags = C.CP_SPACE_DEBUG_DRAW_COLLISION_POINTS
)

type (
	SpaceDebugDrawCircleImpl        func(pos Vect, angle, radius float64, outlineColor, fillColor SpaceDebugColor, data interface{})
	SpaceDebugDrawSegmentImpl       func(a, b Vect, color SpaceDebugColor, data interface{})
	SpaceDebugDrawFatSegmentImpl    func(a, b Vect, radius float64, outlineColor, fillColor SpaceDebugColor, data interface{})
	SpaceDebugDrawPolygonImpl       func(verts []Vect, radius float64, outlineColor, fillColor SpaceDebugColor, data interface{})
	SpaceDebugDrawDotImpl           func(size float64, pos Vect, color SpaceDebugColor, data interface{})
	SpaceDebugDrawColorForShapeImpl func(shape *Shape, data interface{}) SpaceDebugColor
)

type SpaceDebugDrawOptions struct {
	DrawCircle     SpaceDebugDrawCircleImpl
	DrawSegment    SpaceDebugDrawSegmentImpl
	DrawFatSegment SpaceDebugDrawFatSegmentImpl
	DrawPolygon    SpaceDebugDrawPolygonImpl
	DrawDot        SpaceDebugDrawDotImpl

	ColorForShape SpaceDebugDrawColorForShapeImpl

	Flags               SpaceDebugDrawFlags
	ShapeOutlineColor   SpaceDebugColor
	ConstraintColor     SpaceDebugColor
	CollisionPointColor SpaceDebugColor

	Data interface{}
}

//export go_chipmunk_space_debug_draw_circle_impl
func go_chipmunk_space_debug_draw_circle_impl(cpos C.cpVect, cangle, cradius C.cpFloat, coutlineColor, cfillColor C.cpSpaceDebugColor, data unsafe.Pointer) {
	options := (*SpaceDebugDrawOptions)(data)
	if options.DrawCircle != nil {
		pos := *(*Vect)(unsafe.Pointer(&cpos))
		angle := *(*float64)(unsafe.Pointer(&cangle))
		radius := *(*float64)(unsafe.Pointer(&cradius))
		outlineColor := *(*SpaceDebugColor)(unsafe.Pointer(&coutlineColor))
		fillColor := *(*SpaceDebugColor)(unsafe.Pointer(&cfillColor))
		options.DrawCircle(pos, angle, radius, outlineColor, fillColor, options.Data)
	}
}

//export go_chipmunk_space_debug_draw_segment_impl
func go_chipmunk_space_debug_draw_segment_impl(ca, cb C.cpVect, ccolor C.cpSpaceDebugColor, data unsafe.Pointer) {
	options := (*SpaceDebugDrawOptions)(data)
	if options.DrawSegment != nil {
		a := *(*Vect)(unsafe.Pointer(&ca))
		b := *(*Vect)(unsafe.Pointer(&cb))
		color := *(*SpaceDebugColor)(unsafe.Pointer(&ccolor))
		options.DrawSegment(a, b, color, options.Data)
	}
}

//export go_chipmunk_space_debug_draw_fat_segment_impl
func go_chipmunk_space_debug_draw_fat_segment_impl(ca, cb C.cpVect, cradius C.cpFloat, coutlineColor, cfillColor C.cpSpaceDebugColor, data unsafe.Pointer) {
	options := (*SpaceDebugDrawOptions)(data)
	if options.DrawFatSegment != nil {
		a := *(*Vect)(unsafe.Pointer(&ca))
		b := *(*Vect)(unsafe.Pointer(&cb))
		radius := *(*float64)(unsafe.Pointer(&cradius))
		outlineColor := *(*SpaceDebugColor)(unsafe.Pointer(&coutlineColor))
		fillColor := *(*SpaceDebugColor)(unsafe.Pointer(&cfillColor))
		options.DrawFatSegment(a, b, radius, outlineColor, fillColor, options.Data)
	}
}

//export go_chipmunk_space_debug_draw_polygon_impl
func go_chipmunk_space_debug_draw_polygon_impl(ccount C.int, cverts *C.cpVect, cradius C.cpFloat, coutlineColor, cfillColor C.cpSpaceDebugColor, data unsafe.Pointer) {
	options := (*SpaceDebugDrawOptions)(data)
	if options.DrawPolygon != nil {
		var verts []Vect
		sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&verts))
		sliceHeader.Len = int(ccount)
		sliceHeader.Cap = int(ccount)
		sliceHeader.Data = uintptr(unsafe.Pointer(cverts))

		radius := *(*float64)(unsafe.Pointer(&cradius))
		outlineColor := *(*SpaceDebugColor)(unsafe.Pointer(&coutlineColor))
		fillColor := *(*SpaceDebugColor)(unsafe.Pointer(&cfillColor))

		options.DrawPolygon(verts, radius, outlineColor, fillColor, options.Data)
	}
}

//export go_chipmunk_space_debug_draw_dot_impl
func go_chipmunk_space_debug_draw_dot_impl(csize C.cpFloat, cpos C.cpVect, ccolor C.cpSpaceDebugColor, data unsafe.Pointer) {
	options := (*SpaceDebugDrawOptions)(data)
	if options.DrawDot != nil {
		size := *(*float64)(unsafe.Pointer(&csize))
		pos := *(*Vect)(unsafe.Pointer(&cpos))
		color := *(*SpaceDebugColor)(unsafe.Pointer(&ccolor))
		options.DrawDot(size, pos, color, options.Data)
	}
}

//export go_chipmunk_space_debug_draw_color_for_shape_impl
func go_chipmunk_space_debug_draw_color_for_shape_impl(cshape unsafe.Pointer, data unsafe.Pointer) C.cpSpaceDebugColor {
	options := (*SpaceDebugDrawOptions)(data)
	var color SpaceDebugColor
	if options.ColorForShape != nil {
		shape := (*Shape)(cshape)
		color = options.ColorForShape(shape, options.Data)
	}
	return *(*C.cpSpaceDebugColor)(unsafe.Pointer(&color))
}

// DebugDraw draws the space using the debug draw options (which includes
// callbacks for performing the actual drawing). All options fields are
// entirely optional (callbacks may be nil, etc).
//
// options.Data is arbitrary user data fed into the callbacks (this is just for
// a 1:1 mapping of chipmunks API, in Go you can just use a closure and access
// the data itself not storing it inside options.Data).
func (s *Space) DebugDraw(options *SpaceDebugDrawOptions) {
	var dd C.cpSpaceDebugDrawOptions
	dd.drawCircle = (C.cpSpaceDebugDrawCircleImpl)(unsafe.Pointer(C.pre_go_chipmunk_space_debug_draw_circle_impl))
	dd.drawSegment = (C.cpSpaceDebugDrawSegmentImpl)(unsafe.Pointer(C.pre_go_chipmunk_space_debug_draw_segment_impl))
	dd.drawFatSegment = (C.cpSpaceDebugDrawFatSegmentImpl)(unsafe.Pointer(C.pre_go_chipmunk_space_debug_draw_fat_segment_impl))
	dd.drawPolygon = (C.cpSpaceDebugDrawPolygonImpl)(unsafe.Pointer(C.pre_go_chipmunk_space_debug_draw_polygon_impl))
	dd.drawDot = (C.cpSpaceDebugDrawDotImpl)(unsafe.Pointer(C.pre_go_chipmunk_space_debug_draw_dot_impl))
	dd.colorForShape = (C.cpSpaceDebugDrawColorForShapeImpl)(unsafe.Pointer(C.pre_go_chipmunk_space_debug_draw_color_for_shape_impl))

	dd.flags = C.cpSpaceDebugDrawFlags(options.Flags)
	dd.shapeOutlineColor = *(*C.cpSpaceDebugColor)(unsafe.Pointer(&options.ShapeOutlineColor))
	dd.constraintColor = *(*C.cpSpaceDebugColor)(unsafe.Pointer(&options.ConstraintColor))
	dd.collisionPointColor = *(*C.cpSpaceDebugColor)(unsafe.Pointer(&options.CollisionPointColor))

	dd.data = C.cpDataPointer(unsafe.Pointer(options))

	C.cpSpaceDebugDraw(s.c, &dd)
}
