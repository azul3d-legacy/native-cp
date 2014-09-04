// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cp

/*
#include "chipmunk/include/chipmunk/chipmunk.h"
*/
import "C"

// Column major affine transform.
type Transform struct {
	A  float64
	B  float64
	C  float64
	D  float64
	Tx float64
	Ty float64
}

// c converts a Transform to a C.cpTransform
func (m Transform) c() C.cpTransform {
	var cp C.cpTransform
	cp.a = C.cpFloat(m.A)
	cp.b = C.cpFloat(m.B)
	cp.c = C.cpFloat(m.C)
	cp.d = C.cpFloat(m.D)
	cp.tx = C.cpFloat(m.Tx)
	cp.ty = C.cpFloat(m.Ty)
	return cp
}

// goTransform converts C.cpTransform to a Go Transform.
func goTransform(t C.cpTransform) Transform {
	return Transform{
		A:  float64(t.a),
		B:  float64(t.b),
		C:  float64(t.c),
		D:  float64(t.d),
		Tx: float64(t.tx),
		Ty: float64(t.ty),
	}
}

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
	return Transform{
		A:  a,
		B:  b,
		C:  c,
		D:  d,
		Tx: tx,
		Ty: ty,
	}
}

// Construct a new transform matrix in transposed order.
func TransformNewTranspose(a, c, tx, b, d, ty float64) Transform {
	return goTransform(C.cpTransformNewTranspose(
		C.cpFloat(a),
		C.cpFloat(c),
		C.cpFloat(tx),
		C.cpFloat(b),
		C.cpFloat(d),
		C.cpFloat(ty),
	))
}

// Get the inverse of a transform matrix.
func (t Transform) Inverse() Transform {
	return goTransform(C.cpTransformInverse(t.c()))
}

// Multiply two transformation matrices.
func (t1 Transform) Mult(t2 Transform) Transform {
	return goTransform(C.cpTransformMult(t1.c(), t2.c()))
}

// Transform an absolute point. (i.e. a vertex)
func (t Transform) Point(p Vect) Vect {
	return goVect(C.cpTransformPoint(t.c(), p.c()))
}

// Transform a vector (i.e. a normal)
func (t Transform) Vect(v Vect) Vect {
	return goVect(C.cpTransformVect(t.c(), v.c()))
}

// Transform a BB.
func (t Transform) BB(bb BB) BB {
	return goBB(C.cpTransformbBB(t.c(), bb.c()))
}

// Create a translation matrix.
func TransformTranslate(translate Vect) Transform {
	return goTransform(C.cpTransformTranslate(translate.c()))
}

// Create a scale matrix.
func TransformScale(scaleX, scaleY float64) Transform {
	return goTransform(C.cpTransformScale(
		C.cpFloat(scaleX),
		C.cpFloat(scaleY),
	))
}

// Create a rotation matrix.
func TransformRotate(radians float64) Transform {
	return goTransform(C.cpTransformRotate(C.cpFloat(radians)))
}

// Create a rigid transformation matrix. (translation + rotation)
func TransformRigid(translate Vect, radians float64) Transform {
	return goTransform(C.cpTransformRigid(translate.c(), C.cpFloat(radians)))
}

// Fast inverse of a rigid transformation matrix.
func (t Transform) RigidInverse() Transform {
	return goTransform(C.cpTransformRigidInverse(t.c()))
}

func (outer Transform) Wrap(inner Transform) Transform {
	return goTransform(C.cpTransformWrap(outer.c(), inner.c()))
}

func (outer Transform) WrapInverse(inner Transform) Transform {
	return goTransform(C.cpTransformWrapInverse(outer.c(), inner.c()))
}

func TransformOrtho(bb BB) Transform {
	return goTransform(C.cpTransformOrtho(bb.c()))
}

func TransformBoneScale(v0, v1 Vect) Transform {
	return goTransform(C.cpTransformBoneScale(v0.c(), v1.c()))
}

func TransformAxialScale(axis, pivot Vect, scale float64) Transform {
	return goTransform(C.cpTransformAxialScale(
		axis.c(),
		pivot.c(),
		C.cpFloat(scale),
	))
}
