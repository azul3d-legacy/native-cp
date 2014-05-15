// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "_cgo_export.h"

void pre_go_chipmunk_body_velocity_func(cpBody *body, cpVect gravity, cpFloat damping, cpFloat dt) {
	go_chipmunk_body_velocity_func((void*)body, gravity, damping, dt);
}

void pre_go_chipmunk_body_position_func(cpBody *body, cpFloat dt) {
	go_chipmunk_body_position_func((void*)body, dt);
}

void pre_go_chipmunk_body_each_shape(cpBody *body, cpShape *shape, void *data) {
	go_chipmunk_body_each_shape((void*)body, (void*)shape, data);
}

void pre_go_chipmunk_body_each_constraint(cpBody *body, cpConstraint *constraint, void *data) {
	go_chipmunk_body_each_constraint((void*)body, (void*)constraint, data);
}

void pre_go_chipmunk_body_each_arbiter(cpBody *body, cpArbiter *arbiter, void *data) {
	go_chipmunk_body_each_arbiter((void*)body, (void*)arbiter, data);
}

