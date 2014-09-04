// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// +build ignore

package cp

/*
#include "chipmunk/include/chipmunk/chipmunk.h"
*/
import "C"

// Hash value type.
type HashValue C.cpHashValue

// Type used internally to cache colliding object info for cpCollideShapes().
// Should be at least 32 bits.
type CollisionID C.cpCollisionID

// Type used for user data pointers.
type DataPointer C.cpDataPointer

// Type used for cpSpace.collision_type.
type CollisionType C.cpCollisionType

// Type used for cpShape.group.
type Group C.cpGroup

// Type used for cpShapeFilter category and mask.
type Bitmask C.cpBitmask

// Type used for various timestamps in Chipmunk.
type Timestamp C.cpTimestamp

// Column major affine transform.
type Transform C.cpTransform

// Chipmunk's axis-aligned 2D bounding box type. (left, bottom, right, top)
type BB C.cpBB

type SpaceDebugColor C.cpSpaceDebugColor
