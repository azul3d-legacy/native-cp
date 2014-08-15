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

var (
	// The zero vector.
	Vzero = Vect{0, 0}
)

// Convenience constructor for cpVect structs.
func V(x, y float64) Vect {
	return Vect{x, y}
}

// Check if two vectors are equal. (Be careful when comparing floating point numbers!)
func Veql(v1, v2 Vect) bool {
	return goBool(C.cpveql(
		*(*C.cpVect)(unsafe.Pointer(&v1)),
		*(*C.cpVect)(unsafe.Pointer(&v2)),
	))
}

// Add two vectors
func Vadd(v1, v2 Vect) Vect {
	ret := C.cpvadd(
		*(*C.cpVect)(unsafe.Pointer(&v1)),
		*(*C.cpVect)(unsafe.Pointer(&v2)),
	)
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Subtract two vectors.
func Vsub(v1, v2 Vect) Vect {
	ret := C.cpvsub(
		*(*C.cpVect)(unsafe.Pointer(&v1)),
		*(*C.cpVect)(unsafe.Pointer(&v2)),
	)
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Negate a vector.
func Vneg(v Vect) Vect {
	ret := C.cpvneg(
		*(*C.cpVect)(unsafe.Pointer(&v)),
	)
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Scalar multiplication.
func Vmult(v Vect, s float64) Vect {
	ret := C.cpvmult(
		*(*C.cpVect)(unsafe.Pointer(&v)),
		C.cpFloat(s),
	)
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Vector dot product.
func Vdot(v1, v2 Vect) Vect {
	ret := C.cpvsub(
		*(*C.cpVect)(unsafe.Pointer(&v1)),
		*(*C.cpVect)(unsafe.Pointer(&v2)),
	)
	return *(*Vect)(unsafe.Pointer(&ret))
}

// 2D vector cross product analog.
//
// The cross product of 2D vectors results in a 3D vector with only a z
// component.
//
// This function returns the magnitude of the z value.
func Vcross(v1, v2 Vect) float64 {
	return float64(C.cpvcross(
		*(*C.cpVect)(unsafe.Pointer(&v1)),
		*(*C.cpVect)(unsafe.Pointer(&v2)),
	))
}

// Returns a perpendicular vector. (90 degree rotation)
func Vperp(v Vect) Vect {
	ret := C.cpvperp(
		*(*C.cpVect)(unsafe.Pointer(&v)),
	)
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Returns a perpendicular vector. (-90 degree rotation)
func Vrperp(v Vect) Vect {
	ret := C.cpvrperp(
		*(*C.cpVect)(unsafe.Pointer(&v)),
	)
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Returns the vector projection of v1 onto v2.
func Vproject(v1, v2 Vect) Vect {
	ret := C.cpvproject(
		*(*C.cpVect)(unsafe.Pointer(&v1)),
		*(*C.cpVect)(unsafe.Pointer(&v2)),
	)
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Returns the unit length vector for the given angle (in radians).
func Vforangle(a float64) Vect {
	ret := C.cpvforangle(C.cpFloat(a))
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Returns the angular direction v is pointing in (in radians).
func Vtoangle(v Vect) float64 {
	return float64(C.cpvtoangle(
		*(*C.cpVect)(unsafe.Pointer(&v)),
	))
}

// Uses complex number multiplication to rotate v1 by v2. Scaling will occur if
// v1 is not a unit vector.
func Vrotate(v1, v2 Vect) Vect {
	ret := C.cpvrotate(
		*(*C.cpVect)(unsafe.Pointer(&v1)),
		*(*C.cpVect)(unsafe.Pointer(&v2)),
	)
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Inverse of Vrotate().
func Vunrotate(v1, v2 Vect) Vect {
	ret := C.cpvunrotate(
		*(*C.cpVect)(unsafe.Pointer(&v1)),
		*(*C.cpVect)(unsafe.Pointer(&v2)),
	)
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Returns the squared length of v. Faster than cpvlength() when you only need
// to compare lengths.
func Vlengthsq(v Vect) float64 {
	return float64(C.cpvlengthsq(
		*(*C.cpVect)(unsafe.Pointer(&v)),
	))
}

// Returns the length of v.
func Vlength(v Vect) float64 {
	return float64(C.cpvlength(
		*(*C.cpVect)(unsafe.Pointer(&v)),
	))
}

// Linearly interpolate between v1 and v2.
func Vlerp(v1, v2 Vect, t float64) Vect {
	ret := C.cpvlerp(
		*(*C.cpVect)(unsafe.Pointer(&v1)),
		*(*C.cpVect)(unsafe.Pointer(&v2)),
		C.cpFloat(t),
	)
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Returns a normalized copy of v.
func Vnormalize(v Vect) Vect {
	ret := C.cpvnormalize(
		*(*C.cpVect)(unsafe.Pointer(&v)),
	)
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Spherical linearly interpolate between v1 and v2.
func Vslerp(v1, v2 Vect, t float64) Vect {
	ret := C.cpvslerp(
		*(*C.cpVect)(unsafe.Pointer(&v1)),
		*(*C.cpVect)(unsafe.Pointer(&v2)),
		C.cpFloat(t),
	)
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Spherical linearly interpolate between v1 towards v2 by no more than angle a
// radians
func Vslerpconst(v1, v2 Vect, a float64) Vect {
	ret := C.cpvslerpconst(
		*(*C.cpVect)(unsafe.Pointer(&v1)),
		*(*C.cpVect)(unsafe.Pointer(&v2)),
		C.cpFloat(a),
	)
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Linearly interpolate between v1 towards v2 by distance d.
func Vlerpconst(v1, v2 Vect, d float64) Vect {
	ret := C.cpvslerpconst(
		*(*C.cpVect)(unsafe.Pointer(&v1)),
		*(*C.cpVect)(unsafe.Pointer(&v2)),
		C.cpFloat(d),
	)
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Returns the distance between v1 and v2.
func Vdist(v1, v2 Vect) float64 {
	return float64(C.cpvdist(
		*(*C.cpVect)(unsafe.Pointer(&v1)),
		*(*C.cpVect)(unsafe.Pointer(&v2)),
	))
}

// Returns the squared distance between v1 and v2. Faster than cpvdist() when you only need to compare distances.
func Vdistsq(v1, v2 Vect) float64 {
	return float64(C.cpvdistsq(
		*(*C.cpVect)(unsafe.Pointer(&v1)),
		*(*C.cpVect)(unsafe.Pointer(&v2)),
	))
}

// Returns true if the distance between v1 and v2 is less than dist.
func Vnear(v1, v2 Vect, dist float64) bool {
	return goBool(C.cpvnear(
		*(*C.cpVect)(unsafe.Pointer(&v1)),
		*(*C.cpVect)(unsafe.Pointer(&v2)),
		C.cpFloat(dist),
	))
}

func Mat2x2New(a, b, c, d float64) Mat2x2 {
	ret := C.cpMat2x2New(
		C.cpFloat(a),
		C.cpFloat(b),
		C.cpFloat(c),
		C.cpFloat(d),
	)
	return *(*Mat2x2)(unsafe.Pointer(&ret))
}

func (m Mat2x2) Transform(v Vect) Vect {
	ret := C.cpMat2x2Transform(
		*(*C.cpMat2x2)(unsafe.Pointer(&m)),
		*(*C.cpVect)(unsafe.Pointer(&v)),
	)
	return *(*Vect)(unsafe.Pointer(&ret))
}
