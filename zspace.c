// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "chipmunk/chipmunk.h"
#include "_cgo_export.h"

void pre_go_chipmunk_space_point_query_func(cpShape *shape, cpVect point, cpFloat distance, cpVect gradient, void *data) {
	go_chipmunk_space_point_query_func((void*)shape, point, distance, gradient, data);
}

void pre_go_chipmunk_space_segment_query_func(cpShape *shape, cpVect point, cpVect normal, cpFloat alpha, void *data) {
	go_chipmunk_space_segment_query_func((void*)shape, point, normal, alpha, data);
}

void pre_go_chipmunk_space_bb_query_func(cpShape *shape, void *data) {
	go_chipmunk_space_bb_query_func((void*)shape, data);
}

void pre_go_chipmunk_space_shape_query_func(cpShape *shape, cpContactPointSet* points, void *data) {
	go_chipmunk_space_shape_query_func((void*)shape, (void*)points, data);
}

void pre_go_chipmunk_space_body_iterator_func(cpBody* body, void *data) {
	go_chipmunk_space_body_iterator_func((void*)body, data);
}

void pre_go_chipmunk_space_shape_iterator_func(cpShape* shape, void *data) {
	go_chipmunk_space_shape_iterator_func((void*)shape, data);
}

void pre_go_chipmunk_space_constraint_iterator_func(cpConstraint* constraint, void *data) {
	go_chipmunk_space_constraint_iterator_func((void*)constraint, data);
}

void pre_go_chipmunk_space_debug_draw_circle_impl(cpVect pos, cpFloat angle, cpFloat radius, cpSpaceDebugColor outlineColor, cpSpaceDebugColor fillColor, cpDataPointer *data) {
	go_chipmunk_space_debug_draw_circle_impl(pos, angle, radius, outlineColor, fillColor, data);
}

void pre_go_chipmunk_space_debug_draw_segment_impl(cpVect a, cpVect b, cpSpaceDebugColor color, cpDataPointer *data) {
	go_chipmunk_space_debug_draw_segment_impl(a, b, color, data);
}

void pre_go_chipmunk_space_debug_draw_fat_segment_impl(cpVect a, cpVect b, cpFloat radius, cpSpaceDebugColor outlineColor, cpSpaceDebugColor fillColor, cpDataPointer *data) {
	go_chipmunk_space_debug_draw_fat_segment_impl(a, b, radius, outlineColor, fillColor, data);
}

void pre_go_chipmunk_space_debug_draw_polygon_impl(int count, const cpVect *verts, cpFloat radius, cpSpaceDebugColor outlineColor, cpSpaceDebugColor fillColor, cpDataPointer *data) {
	go_chipmunk_space_debug_draw_polygon_impl(count, (cpVect*)verts, radius, outlineColor, fillColor, data);
}

void pre_go_chipmunk_space_debug_draw_dot_impl(cpFloat size, cpVect pos, cpSpaceDebugColor color, cpDataPointer *data) {
	go_chipmunk_space_debug_draw_dot_impl(size, pos, color, data);
}

cpSpaceDebugColor pre_go_chipmunk_space_debug_draw_color_for_shape_impl(cpShape *shape, cpDataPointer *data) {
	go_chipmunk_space_debug_draw_color_for_shape_impl(shape, data);
}

