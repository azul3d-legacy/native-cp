// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs ztypes.go

package cp

import "unsafe"

type HashValue uint64

type CollisionID uint32

type DataPointer unsafe.Pointer

type CollisionType uint64

type Group uint64

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

type BB struct {
	L float64
	B float64
	R float64
	T float64
}

type SpaceDebugColor struct {
	R float32
	G float32
	B float32
	A float32
}
