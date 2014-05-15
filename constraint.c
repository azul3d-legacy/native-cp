// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "_cgo_export.h"

void pre_go_chipmunk_constraint_pre_solve_func(cpConstraint *constraint, cpSpace *space) {
	go_chipmunk_constraint_pre_solve_func((void*)constraint, (void*)space);
}

void pre_go_chipmunk_constraint_post_solve_func(cpConstraint *constraint, cpSpace *space) {
	go_chipmunk_constraint_post_solve_func((void*)constraint, (void*)space);
}

