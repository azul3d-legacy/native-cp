// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cp

// indexes are data structures that are used to accelerate collision detection
// and spatial queries. Chipmunk provides a number of spatial index algorithms
// to pick from and they are programmed in a generic way so that you can use
// them for holding more than just cpShape structs.
//
// It works by using void pointers to the objects you add and using a callback
// to ask your code for bounding boxes when it needs them. Several types of
// queries can be performed an index as well as reindexing and full collision
// information. All communication to the spatial indexes is performed through
// callback functions.

/*
// Spatial index bounding box callback function type.
// The spatial index calls this function and passes you a pointer to an object you added
// when it needs to get the bounding box associated with that object.
typedef cpBB (*cpSpatialIndexBBFunc)(void *obj);

// Spatial index/object iterator callback function type.
typedef void (*cpSpatialIndexIteratorFunc)(void *obj, void *data);

// Spatial query callback function type.
typedef cpCollisionID (*cpSpatialIndexQueryFunc)(void *obj1, void *obj2, cpCollisionID id, void *data);

// Spatial segment query callback function type.
typedef cpFloat (*cpSpatialIndexSegmentQueryFunc)(void *obj1, void *obj2, void *data);

typedef struct cpSpatialIndexClass cpSpatialIndexClass;
typedef struct cpSpatialIndex cpSpatialIndex;
typedef struct cpSpaceHash cpSpaceHash;
typedef struct cpBBTree cpBBTree;
typedef struct cpSweep1D cpSweep1D;

// Allocate and initialize a spatial hash.
cpSpatialIndex* cpSpaceHashNew(cpFloat celldim, int cells, cpSpatialIndexBBFunc bbfunc, cpSpatialIndex *staticIndex);

// Change the cell dimensions and table size of the spatial hash to tune it.
//
// The cell dimensions should roughly match the average size of your objects
// and the table size should be ~10 larger than the number of objects inserted.
//
// Some trial and error is required to find the optimum numbers for efficiency.
void cpSpaceHashResize(cpSpaceHash *hash, cpFloat celldim, int numcells);

// Allocate a bounding box tree.
cpBBTree* cpBBTreeAlloc(void);

// Initialize a bounding box tree.
cpSpatialIndex* cpBBTreeInit(cpBBTree *tree, cpSpatialIndexBBFunc bbfunc, cpSpatialIndex *staticIndex);

// Allocate and initialize a bounding box tree.
cpSpatialIndex* cpBBTreeNew(cpSpatialIndexBBFunc bbfunc, cpSpatialIndex *staticIndex);

// Perform a static top down optimization of the tree.
void cpBBTreeOptimize(cpSpatialIndex *index);

// Bounding box tree velocity callback function.
// This function should return an estimate for the object's velocity.
typedef cpVect (*cpBBTreeVelocityFunc)(void *obj);

// Set the velocity function for the bounding box tree to enable temporal coherence.
void cpBBTreeSetVelocityFunc(cpSpatialIndex *index, cpBBTreeVelocityFunc func);

// Allocate a 1D sort and sweep broadphase.
cpSweep1D* cpSweep1DAlloc(void);

// Initialize a 1D sort and sweep broadphase.
cpSpatialIndex* cpSweep1DInit(cpSweep1D *sweep, cpSpatialIndexBBFunc bbfunc, cpSpatialIndex *staticIndex);

// Allocate and initialize a 1D sort and sweep broadphase.
cpSpatialIndex* cpSweep1DNew(cpSpatialIndexBBFunc bbfunc, cpSpatialIndex *staticIndex);

//MARK: Spatial Index Implementation
typedef void (*cpSpatialIndexDestroyImpl)(cpSpatialIndex *index);
typedef int (*cpSpatialIndexCountImpl)(cpSpatialIndex *index);
typedef void (*cpSpatialIndexEachImpl)(cpSpatialIndex *index, cpSpatialIndexIteratorFunc func, void *data);
typedef cpBool (*cpSpatialIndexContainsImpl)(cpSpatialIndex *index, void *obj, cpHashValue hashid);
typedef void (*cpSpatialIndexInsertImpl)(cpSpatialIndex *index, void *obj, cpHashValue hashid);
typedef void (*cpSpatialIndexRemoveImpl)(cpSpatialIndex *index, void *obj, cpHashValue hashid);
typedef void (*cpSpatialIndexReindexImpl)(cpSpatialIndex *index);
typedef void (*cpSpatialIndexReindexObjectImpl)(cpSpatialIndex *index, void *obj, cpHashValue hashid);
typedef void (*cpSpatialIndexReindexQueryImpl)(cpSpatialIndex *index, cpSpatialIndexQueryFunc func, void *data);
typedef void (*cpSpatialIndexQueryImpl)(cpSpatialIndex *index, void *obj, cpBB bb, cpSpatialIndexQueryFunc func, void *data);
typedef void (*cpSpatialIndexSegmentQueryImpl)(cpSpatialIndex *index, void *obj, cpVect a, cpVect b, cpFloat t_exit, cpSpatialIndexSegmentQueryFunc func, void *data);

// Destroy and free a spatial index.
void cpSpatialIndexFree(cpSpatialIndex *index);

// Collide the objects in @c dynamicIndex against the objects in @c staticIndex using the query callback function.
void cpSpatialIndexCollideStatic(cpSpatialIndex *dynamicIndex, cpSpatialIndex *staticIndex, cpSpatialIndexQueryFunc func, void *data);

// Get the number of objects in the spatial index.
static inline int cpSpatialIndexCount(cpSpatialIndex *index)

// Iterate the objects in the spatial index. @c func will be called once for each object.
static inline void cpSpatialIndexEach(cpSpatialIndex *index, cpSpatialIndexIteratorFunc func, void *data)

// Returns true if the spatial index contains the given object.
// Most spatial indexes use hashed storage, so you must provide a hash value too.
static inline cpBool cpSpatialIndexContains(cpSpatialIndex *index, void *obj, cpHashValue hashid)

// Add an object to a spatial index.
// Most spatial indexes use hashed storage, so you must provide a hash value too.
static inline void cpSpatialIndexInsert(cpSpatialIndex *index, void *obj, cpHashValue hashid)

// Remove an object from a spatial index.
// Most spatial indexes use hashed storage, so you must provide a hash value too.
static inline void cpSpatialIndexRemove(cpSpatialIndex *index, void *obj, cpHashValue hashid)

// Perform a full reindex of a spatial index.
static inline void cpSpatialIndexReindex(cpSpatialIndex *index)

// Reindex a single object in the spatial index.
static inline void cpSpatialIndexReindexObject(cpSpatialIndex *index, void *obj, cpHashValue hashid)

// Perform a rectangle query against the spatial index, calling @c func for each potential match.
static inline void cpSpatialIndexQuery(cpSpatialIndex *index, void *obj, cpBB bb, cpSpatialIndexQueryFunc func, void *data)

// Perform a segment query against the spatial index, calling @c func for each potential match.
static inline void cpSpatialIndexSegmentQuery(cpSpatialIndex *index, void *obj, cpVect a, cpVect b, cpFloat t_exit, cpSpatialIndexSegmentQueryFunc func, void *data)

// Simultaneously reindex and find all colliding objects.
// @c func will be called once for each potentially overlapping pair of objects found.
// If the spatial index was initialized with a static index, it will collide it's objects against that as well.
static inline void cpSpatialIndexReindexQuery(cpSpatialIndex *index, cpSpatialIndexQueryFunc func, void *data)
*/
