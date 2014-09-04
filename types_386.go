// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs ztypes_cgodefs.go

package cp

import "unsafe"

type HashValue uint32

type CollisionID uint32

type DataPointer unsafe.Pointer

type CollisionType uint32

type Group uint32

type Bitmask uint32

type Timestamp uint32

type Transform struct {
	A  float64
	B  float64
	C  float64
	D  float64
	Tx float64
	Ty float64
}

type SpaceDebugColor struct {
	R float32
	G float32
	B float32
	A float32
}
