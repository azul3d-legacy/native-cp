// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cp

/*
#include "chipmunk/include/chipmunk/chipmunk.h"
*/
import "C"

const (
	// Value for cpShape.group signifying that a shape is in no group.
	NO_GROUP = C.CP_NO_GROUP

	// Value for cpShape.layers signifying that a shape is in every layer.
	ALL_CATEGORIES = C.CP_ALL_CATEGORIES

	// cpCollisionType value internally reserved for hashing wildcard handlers.
	WILDCARD_COLLISION_TYPE = C.CP_WILDCARD_COLLISION_TYPE
)

func goBool(c C.cpBool) bool {
	if c == C.cpTrue {
		return true
	}
	return false
}

// Return the max of two cpFloats.
func Fmax(a, b float64) float64 {
	return float64(C.cpfmax(
		C.cpFloat(a),
		C.cpFloat(b),
	))
}

// Return the min of two cpFloats.
func Fmin(a, b float64) float64 {
	return float64(C.cpfmin(
		C.cpFloat(a),
		C.cpFloat(b),
	))
}

// Return the absolute value of a cpFloat.
func Fabs(f float64) float64 {
	return float64(C.cpfabs(
		C.cpFloat(f),
	))
}

// Clamp f to be between min and max.
func Fclamp(f, min, max float64) float64 {
	return float64(C.cpfclamp(
		C.cpFloat(f),
		C.cpFloat(min),
		C.cpFloat(max),
	))
}

// Clamp f to be between 0 and 1.
func Fclamp01(f float64) float64 {
	return float64(C.cpfclamp01(
		C.cpFloat(f),
	))
}

// Linearly interpolate (or extrapolate) between f1 and f2 by t percent.
func Flerp(f1, f2, t float64) float64 {
	return float64(C.cpflerp(
		C.cpFloat(f1),
		C.cpFloat(f2),
		C.cpFloat(t),
	))
}

// Linearly interpolate from f1 to f2 by no more than d.
func Flerpconst(f1, f2, d float64) float64 {
	return float64(C.cpflerpconst(
		C.cpFloat(f1),
		C.cpFloat(f2),
		C.cpFloat(d),
	))
}
