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

// Identity transform matrix.
var (
	TransformIdentity = Transform{1, 0, 0, 1, 0, 0}
)

// Construct a new transform matrix.
//
// (a, b) is the x basis vector.
//
// (c, d) is the y basis vector.
//
// (tx, ty) is the translation.
//
func TransformNew(a, b, c, d, tx, ty float64) Transform {
	ret := C.cpTransformNew(
		C.cpFloat(a),
		C.cpFloat(b),
		C.cpFloat(c),
		C.cpFloat(d),
		C.cpFloat(tx),
		C.cpFloat(ty),
	)
	return *(*Transform)(unsafe.Pointer(&ret))
}

// Construct a new transform matrix in transposed order.
func TransformNewTranspose(a, c, tx, b, d, ty float64) Transform {
	ret := C.cpTransformNewTranspose(
		C.cpFloat(a),
		C.cpFloat(c),
		C.cpFloat(tx),
		C.cpFloat(b),
		C.cpFloat(d),
		C.cpFloat(ty),
	)
	return *(*Transform)(unsafe.Pointer(&ret))
}

// Get the inverse of a transform matrix.
func (t Transform) Inverse() Transform {
	ret := C.cpTransformInverse(
		*(*C.cpTransform)(unsafe.Pointer(&t)),
	)
	return *(*Transform)(unsafe.Pointer(&ret))
}

// Multiply two transformation matrices.
func (t1 Transform) Mult(t2 Transform) Transform {
	ret := C.cpTransformMult(
		*(*C.cpTransform)(unsafe.Pointer(&t1)),
		*(*C.cpTransform)(unsafe.Pointer(&t2)),
	)
	return *(*Transform)(unsafe.Pointer(&ret))
}

// Transform an absolute point. (i.e. a vertex)
func (t Transform) Point(p Vect) Vect {
	ret := C.cpTransformPoint(
		*(*C.cpTransform)(unsafe.Pointer(&t)),
		*(*C.cpVect)(unsafe.Pointer(&p)),
	)
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Transform a vector (i.e. a normal)
func (t Transform) Vect(v Vect) Vect {
	ret := C.cpTransformVect(
		*(*C.cpTransform)(unsafe.Pointer(&t)),
		*(*C.cpVect)(unsafe.Pointer(&v)),
	)
	return *(*Vect)(unsafe.Pointer(&ret))
}

// Transform a BB.
func (t Transform) BB(bb BB) BB {
	ret := C.cpTransformbBB(
		*(*C.cpTransform)(unsafe.Pointer(&t)),
		*(*C.cpBB)(unsafe.Pointer(&bb)),
	)
	return *(*BB)(unsafe.Pointer(&ret))
}

// Create a translation matrix.
func TransformTranslate(translate Vect) Transform {
	ret := C.cpTransformTranslate(
		*(*C.cpVect)(unsafe.Pointer(&translate)),
	)
	return *(*Transform)(unsafe.Pointer(&ret))
}

// Create a scale matrix.
func TransformScale(scaleX, scaleY float64) Transform {
	ret := C.cpTransformScale(
		C.cpFloat(scaleX),
		C.cpFloat(scaleY),
	)
	return *(*Transform)(unsafe.Pointer(&ret))
}

// Create a rotation matrix.
func TransformRotate(radians float64) Transform {
	ret := C.cpTransformRotate(
		C.cpFloat(radians),
	)
	return *(*Transform)(unsafe.Pointer(&ret))
}

// Create a rigid transformation matrix. (translation + rotation)
func TransformRigid(translate Vect, radians float64) Transform {
	ret := C.cpTransformRigid(
		*(*C.cpVect)(unsafe.Pointer(&translate)),
		C.cpFloat(radians),
	)
	return *(*Transform)(unsafe.Pointer(&ret))
}

// Fast inverse of a rigid transformation matrix.
func (t Transform) RigidInverse() Transform {
	ret := C.cpTransformRigidInverse(
		*(*C.cpTransform)(unsafe.Pointer(&t)),
	)
	return *(*Transform)(unsafe.Pointer(&ret))
}

func (outer Transform) Wrap(inner Transform) Transform {
	ret := C.cpTransformWrap(
		*(*C.cpTransform)(unsafe.Pointer(&outer)),
		*(*C.cpTransform)(unsafe.Pointer(&inner)),
	)
	return *(*Transform)(unsafe.Pointer(&ret))
}

func (outer Transform) WrapInverse(inner Transform) Transform {
	ret := C.cpTransformWrapInverse(
		*(*C.cpTransform)(unsafe.Pointer(&outer)),
		*(*C.cpTransform)(unsafe.Pointer(&inner)),
	)
	return *(*Transform)(unsafe.Pointer(&ret))
}

func TransformOrtho(bb BB) Transform {
	ret := C.cpTransformOrtho(
		*(*C.cpBB)(unsafe.Pointer(&bb)),
	)
	return *(*Transform)(unsafe.Pointer(&ret))
}

func TransformBoneScale(v0, v1 Vect) Transform {
	ret := C.cpTransformBoneScale(
		*(*C.cpVect)(unsafe.Pointer(&v0)),
		*(*C.cpVect)(unsafe.Pointer(&v1)),
	)
	return *(*Transform)(unsafe.Pointer(&ret))
}

func TransformAxialScale(axis, pivot Vect, scale float64) Transform {
	ret := C.cpTransformAxialScale(
		*(*C.cpVect)(unsafe.Pointer(&axis)),
		*(*C.cpVect)(unsafe.Pointer(&pivot)),
		C.float(scale),
	)
	return *(*Transform)(unsafe.Pointer(&ret))
}
