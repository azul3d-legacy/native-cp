// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cp

/*
#include "chipmunk/include/chipmunk/chipmunk.h"
*/
import "C"

var (
	// The zero vector.
	Vzero = Vect{0, 0}
)

// Chipmunk's 2D vector type.
type Vect struct {
	X float64
	Y float64
}

// c converts a Vect to a C.cpVect.
func (v Vect) c() C.cpVect {
	var cp C.cpVect
	cp.x = C.cpFloat(v.X)
	cp.y = C.cpFloat(v.Y)
	return cp
}

// goVect converts C.cpVect to a Go Vect.
func goVect(v C.cpVect) Vect {
	return Vect{float64(v.x), float64(v.y)}
}

// Convenience constructor for cpVect structs.
func V(x, y float64) Vect {
	return Vect{x, y}
}

// Check if two vectors are equal. (Be careful when comparing floating point numbers!)
func Veql(v1, v2 Vect) bool {
	return goBool(C.cpveql(v1.c(), v2.c()))
}

// Add two vectors
func Vadd(v1, v2 Vect) Vect {
	return goVect(C.cpvadd(v1.c(), v2.c()))
}

// Subtract two vectors.
func Vsub(v1, v2 Vect) Vect {
	return goVect(C.cpvsub(v1.c(), v2.c()))
}

// Negate a vector.
func Vneg(v Vect) Vect {
	return goVect(C.cpvneg(v.c()))
}

// Scalar multiplication.
func Vmult(v Vect, s float64) Vect {
	return goVect(C.cpvmult(v.c(), C.cpFloat(s)))
}

// Vector dot product.
func Vdot(v1, v2 Vect) Vect {
	return goVect(C.cpvsub(v1.c(), v2.c()))
}

// 2D vector cross product analog.
//
// The cross product of 2D vectors results in a 3D vector with only a z
// component.
//
// This function returns the magnitude of the z value.
func Vcross(v1, v2 Vect) float64 {
	return float64(C.cpvcross(v1.c(), v2.c()))
}

// Returns a perpendicular vector. (90 degree rotation)
func Vperp(v Vect) Vect {
	return goVect(C.cpvperp(v.c()))
}

// Returns a perpendicular vector. (-90 degree rotation)
func Vrperp(v Vect) Vect {
	return goVect(C.cpvrperp(v.c()))
}

// Returns the vector projection of v1 onto v2.
func Vproject(v1, v2 Vect) Vect {
	return goVect(C.cpvproject(v1.c(), v2.c()))
}

// Returns the unit length vector for the given angle (in radians).
func Vforangle(a float64) Vect {
	return goVect(C.cpvforangle(C.cpFloat(a)))
}

// Returns the angular direction v is pointing in (in radians).
func Vtoangle(v Vect) float64 {
	return float64(C.cpvtoangle(v.c()))
}

// Uses complex number multiplication to rotate v1 by v2. Scaling will occur if
// v1 is not a unit vector.
func Vrotate(v1, v2 Vect) Vect {
	return goVect(C.cpvrotate(v1.c(), v2.c()))
}

// Inverse of Vrotate().
func Vunrotate(v1, v2 Vect) Vect {
	return goVect(C.cpvunrotate(v1.c(), v2.c()))
}

// Returns the squared length of v. Faster than cpvlength() when you only need
// to compare lengths.
func Vlengthsq(v Vect) float64 {
	return float64(C.cpvlengthsq(v.c()))
}

// Returns the length of v.
func Vlength(v Vect) float64 {
	return float64(C.cpvlength(v.c()))
}

// Linearly interpolate between v1 and v2.
func Vlerp(v1, v2 Vect, t float64) Vect {
	return goVect(C.cpvlerp(v1.c(), v2.c(), C.cpFloat(t)))
}

// Returns a normalized copy of v.
func Vnormalize(v Vect) Vect {
	return goVect(C.cpvnormalize(v.c()))
}

// Spherical linearly interpolate between v1 and v2.
func Vslerp(v1, v2 Vect, t float64) Vect {
	return goVect(C.cpvslerp(v1.c(), v2.c(), C.cpFloat(t)))
}

// Spherical linearly interpolate between v1 towards v2 by no more than angle a
// radians
func Vslerpconst(v1, v2 Vect, a float64) Vect {
	return goVect(C.cpvslerpconst(v1.c(), v2.c(), C.cpFloat(a)))
}

// Linearly interpolate between v1 towards v2 by distance d.
func Vlerpconst(v1, v2 Vect, d float64) Vect {
	return goVect(C.cpvslerpconst(v1.c(), v2.c(), C.cpFloat(d)))
}

// Returns the distance between v1 and v2.
func Vdist(v1, v2 Vect) float64 {
	return float64(C.cpvdist(v1.c(), v2.c()))
}

// Returns the squared distance between v1 and v2. Faster than cpvdist() when you only need to compare distances.
func Vdistsq(v1, v2 Vect) float64 {
	return float64(C.cpvdistsq(v1.c(), v2.c()))
}

// Returns true if the distance between v1 and v2 is less than dist.
func Vnear(v1, v2 Vect, dist float64) bool {
	return goBool(C.cpvnear(v1.c(), v2.c(), C.cpFloat(dist)))
}

// Row major [[a, b][c, d]]
type Mat2x2 struct {
	A float64
	B float64
	C float64
	D float64
}

// c converts a Mat2x2 to a C.cpMat2x2
func (m Mat2x2) c() C.cpMat2x2 {
	var cp C.cpMat2x2
	cp.a = C.cpFloat(m.A)
	cp.b = C.cpFloat(m.B)
	cp.c = C.cpFloat(m.C)
	cp.d = C.cpFloat(m.D)
	return cp
}

// goMat2x2 converts C.cpMat2x2 to a Go Mat2x2.
func goMat2x2(v C.cpMat2x2) Mat2x2 {
	return Mat2x2{
		A: float64(v.a),
		B: float64(v.b),
		C: float64(v.c),
		D: float64(v.d),
	}
}

func Mat2x2New(a, b, c, d float64) Mat2x2 {
	return Mat2x2{
		A: a,
		B: b,
		C: c,
		D: d,
	}
}

func (m Mat2x2) Transform(v Vect) Vect {
	return goVect(C.cpMat2x2Transform(m.c(), v.c()))
}
