// Code generated by execgen; DO NOT EDIT.
// Copyright 2018 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexec

import (
	"strings"
	"unsafe"

	"github.com/cockroachdb/apd/v2"
	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/sql/colexecbase/colexecerror"
	"github.com/cockroachdb/cockroach/pkg/sql/colmem"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/types"
	"github.com/cockroachdb/cockroach/pkg/util/duration"
	"github.com/cockroachdb/errors"
)

func newSumOrderedAggAlloc(
	allocator *colmem.Allocator, t *types.T, allocSize int64,
) (aggregateFuncAlloc, error) {
	allocBase := aggAllocBase{allocator: allocator, allocSize: allocSize}
	switch t.Family() {
	case types.IntFamily:
		switch t.Width() {
		case 16:
			return &sumInt16OrderedAggAlloc{aggAllocBase: allocBase}, nil
		case 32:
			return &sumInt32OrderedAggAlloc{aggAllocBase: allocBase}, nil
		default:
			return &sumInt64OrderedAggAlloc{aggAllocBase: allocBase}, nil
		}
	case types.DecimalFamily:
		return &sumDecimalOrderedAggAlloc{aggAllocBase: allocBase}, nil
	case types.FloatFamily:
		return &sumFloat64OrderedAggAlloc{aggAllocBase: allocBase}, nil
	case types.IntervalFamily:
		return &sumIntervalOrderedAggAlloc{aggAllocBase: allocBase}, nil
	default:
		return nil, errors.Errorf("unsupported sum %s agg type %s", strings.ToLower(""), t.Name())
	}
}

type sumInt16OrderedAgg struct {
	orderedAggregateFuncBase
	scratch struct {
		// curAgg holds the running total, so we can index into the slice once per
		// group, instead of on each iteration.
		curAgg apd.Decimal
		// vec points to the output vector we are updating.
		vec []apd.Decimal
		// foundNonNullForCurrentGroup tracks if we have seen any non-null values
		// for the group that is currently being aggregated.
		foundNonNullForCurrentGroup bool
	}
	overloadHelper overloadHelper
}

var _ aggregateFunc = &sumInt16OrderedAgg{}

const sizeOfSumInt16OrderedAgg = int64(unsafe.Sizeof(sumInt16OrderedAgg{}))

func (a *sumInt16OrderedAgg) Init(groups []bool, vec coldata.Vec) {
	a.orderedAggregateFuncBase.Init(groups, vec)
	a.scratch.vec = vec.Decimal()
	a.Reset()
}

func (a *sumInt16OrderedAgg) Reset() {
	a.orderedAggregateFuncBase.Reset()
	a.scratch.foundNonNullForCurrentGroup = false
}

func (a *sumInt16OrderedAgg) Compute(
	vecs []coldata.Vec, inputIdxs []uint32, inputLen int, sel []int,
) {
	// In order to inline the templated code of overloads, we need to have a
	// "_overloadHelper" local variable of type "overloadHelper".
	_overloadHelper := a.overloadHelper
	vec := vecs[inputIdxs[0]]
	col, nulls := vec.Int16(), vec.Nulls()
	groups := a.groups
	if sel == nil {
		_ = groups[inputLen-1]
		col = col[:inputLen]
		if nulls.MaybeHasNulls() {
			for i := range col {

				if groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.nulls.SetNull(a.curIdx)
					} else {
						a.scratch.vec[a.curIdx] = a.scratch.curAgg
					}
					a.curIdx++
					a.scratch.curAgg = zeroDecimalValue

					a.scratch.foundNonNullForCurrentGroup = false
				}

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{

						tmpDec := &_overloadHelper.tmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curAgg, &a.scratch.curAgg, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			for i := range col {

				if groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.nulls.SetNull(a.curIdx)
					} else {
						a.scratch.vec[a.curIdx] = a.scratch.curAgg
					}
					a.curIdx++
					a.scratch.curAgg = zeroDecimalValue

				}

				var isNull bool
				isNull = false
				if !isNull {

					{

						tmpDec := &_overloadHelper.tmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curAgg, &a.scratch.curAgg, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	} else {
		sel = sel[:inputLen]
		if nulls.MaybeHasNulls() {
			for _, i := range sel {

				if groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.nulls.SetNull(a.curIdx)
					} else {
						a.scratch.vec[a.curIdx] = a.scratch.curAgg
					}
					a.curIdx++
					a.scratch.curAgg = zeroDecimalValue

					a.scratch.foundNonNullForCurrentGroup = false
				}

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{

						tmpDec := &_overloadHelper.tmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curAgg, &a.scratch.curAgg, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			for _, i := range sel {

				if groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.nulls.SetNull(a.curIdx)
					} else {
						a.scratch.vec[a.curIdx] = a.scratch.curAgg
					}
					a.curIdx++
					a.scratch.curAgg = zeroDecimalValue

				}

				var isNull bool
				isNull = false
				if !isNull {

					{

						tmpDec := &_overloadHelper.tmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curAgg, &a.scratch.curAgg, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	}
}

func (a *sumInt16OrderedAgg) Flush(outputIdx int) {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// null.
	// Go around "argument overwritten before first use" linter error.
	_ = outputIdx
	outputIdx = a.curIdx
	a.curIdx++
	if !a.scratch.foundNonNullForCurrentGroup {
		a.nulls.SetNull(outputIdx)
	} else {
		a.scratch.vec[outputIdx] = a.scratch.curAgg
	}
}

type sumInt16OrderedAggAlloc struct {
	aggAllocBase
	aggFuncs []sumInt16OrderedAgg
}

var _ aggregateFuncAlloc = &sumInt16OrderedAggAlloc{}

func (a *sumInt16OrderedAggAlloc) newAggFunc() aggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(sizeOfSumInt16OrderedAgg * a.allocSize)
		a.aggFuncs = make([]sumInt16OrderedAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	a.aggFuncs = a.aggFuncs[1:]
	return f
}

type sumInt32OrderedAgg struct {
	orderedAggregateFuncBase
	scratch struct {
		// curAgg holds the running total, so we can index into the slice once per
		// group, instead of on each iteration.
		curAgg apd.Decimal
		// vec points to the output vector we are updating.
		vec []apd.Decimal
		// foundNonNullForCurrentGroup tracks if we have seen any non-null values
		// for the group that is currently being aggregated.
		foundNonNullForCurrentGroup bool
	}
	overloadHelper overloadHelper
}

var _ aggregateFunc = &sumInt32OrderedAgg{}

const sizeOfSumInt32OrderedAgg = int64(unsafe.Sizeof(sumInt32OrderedAgg{}))

func (a *sumInt32OrderedAgg) Init(groups []bool, vec coldata.Vec) {
	a.orderedAggregateFuncBase.Init(groups, vec)
	a.scratch.vec = vec.Decimal()
	a.Reset()
}

func (a *sumInt32OrderedAgg) Reset() {
	a.orderedAggregateFuncBase.Reset()
	a.scratch.foundNonNullForCurrentGroup = false
}

func (a *sumInt32OrderedAgg) Compute(
	vecs []coldata.Vec, inputIdxs []uint32, inputLen int, sel []int,
) {
	// In order to inline the templated code of overloads, we need to have a
	// "_overloadHelper" local variable of type "overloadHelper".
	_overloadHelper := a.overloadHelper
	vec := vecs[inputIdxs[0]]
	col, nulls := vec.Int32(), vec.Nulls()
	groups := a.groups
	if sel == nil {
		_ = groups[inputLen-1]
		col = col[:inputLen]
		if nulls.MaybeHasNulls() {
			for i := range col {

				if groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.nulls.SetNull(a.curIdx)
					} else {
						a.scratch.vec[a.curIdx] = a.scratch.curAgg
					}
					a.curIdx++
					a.scratch.curAgg = zeroDecimalValue

					a.scratch.foundNonNullForCurrentGroup = false
				}

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{

						tmpDec := &_overloadHelper.tmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curAgg, &a.scratch.curAgg, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			for i := range col {

				if groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.nulls.SetNull(a.curIdx)
					} else {
						a.scratch.vec[a.curIdx] = a.scratch.curAgg
					}
					a.curIdx++
					a.scratch.curAgg = zeroDecimalValue

				}

				var isNull bool
				isNull = false
				if !isNull {

					{

						tmpDec := &_overloadHelper.tmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curAgg, &a.scratch.curAgg, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	} else {
		sel = sel[:inputLen]
		if nulls.MaybeHasNulls() {
			for _, i := range sel {

				if groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.nulls.SetNull(a.curIdx)
					} else {
						a.scratch.vec[a.curIdx] = a.scratch.curAgg
					}
					a.curIdx++
					a.scratch.curAgg = zeroDecimalValue

					a.scratch.foundNonNullForCurrentGroup = false
				}

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{

						tmpDec := &_overloadHelper.tmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curAgg, &a.scratch.curAgg, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			for _, i := range sel {

				if groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.nulls.SetNull(a.curIdx)
					} else {
						a.scratch.vec[a.curIdx] = a.scratch.curAgg
					}
					a.curIdx++
					a.scratch.curAgg = zeroDecimalValue

				}

				var isNull bool
				isNull = false
				if !isNull {

					{

						tmpDec := &_overloadHelper.tmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curAgg, &a.scratch.curAgg, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	}
}

func (a *sumInt32OrderedAgg) Flush(outputIdx int) {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// null.
	// Go around "argument overwritten before first use" linter error.
	_ = outputIdx
	outputIdx = a.curIdx
	a.curIdx++
	if !a.scratch.foundNonNullForCurrentGroup {
		a.nulls.SetNull(outputIdx)
	} else {
		a.scratch.vec[outputIdx] = a.scratch.curAgg
	}
}

type sumInt32OrderedAggAlloc struct {
	aggAllocBase
	aggFuncs []sumInt32OrderedAgg
}

var _ aggregateFuncAlloc = &sumInt32OrderedAggAlloc{}

func (a *sumInt32OrderedAggAlloc) newAggFunc() aggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(sizeOfSumInt32OrderedAgg * a.allocSize)
		a.aggFuncs = make([]sumInt32OrderedAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	a.aggFuncs = a.aggFuncs[1:]
	return f
}

type sumInt64OrderedAgg struct {
	orderedAggregateFuncBase
	scratch struct {
		// curAgg holds the running total, so we can index into the slice once per
		// group, instead of on each iteration.
		curAgg apd.Decimal
		// vec points to the output vector we are updating.
		vec []apd.Decimal
		// foundNonNullForCurrentGroup tracks if we have seen any non-null values
		// for the group that is currently being aggregated.
		foundNonNullForCurrentGroup bool
	}
	overloadHelper overloadHelper
}

var _ aggregateFunc = &sumInt64OrderedAgg{}

const sizeOfSumInt64OrderedAgg = int64(unsafe.Sizeof(sumInt64OrderedAgg{}))

func (a *sumInt64OrderedAgg) Init(groups []bool, vec coldata.Vec) {
	a.orderedAggregateFuncBase.Init(groups, vec)
	a.scratch.vec = vec.Decimal()
	a.Reset()
}

func (a *sumInt64OrderedAgg) Reset() {
	a.orderedAggregateFuncBase.Reset()
	a.scratch.foundNonNullForCurrentGroup = false
}

func (a *sumInt64OrderedAgg) Compute(
	vecs []coldata.Vec, inputIdxs []uint32, inputLen int, sel []int,
) {
	// In order to inline the templated code of overloads, we need to have a
	// "_overloadHelper" local variable of type "overloadHelper".
	_overloadHelper := a.overloadHelper
	vec := vecs[inputIdxs[0]]
	col, nulls := vec.Int64(), vec.Nulls()
	groups := a.groups
	if sel == nil {
		_ = groups[inputLen-1]
		col = col[:inputLen]
		if nulls.MaybeHasNulls() {
			for i := range col {

				if groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.nulls.SetNull(a.curIdx)
					} else {
						a.scratch.vec[a.curIdx] = a.scratch.curAgg
					}
					a.curIdx++
					a.scratch.curAgg = zeroDecimalValue

					a.scratch.foundNonNullForCurrentGroup = false
				}

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{

						tmpDec := &_overloadHelper.tmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curAgg, &a.scratch.curAgg, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			for i := range col {

				if groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.nulls.SetNull(a.curIdx)
					} else {
						a.scratch.vec[a.curIdx] = a.scratch.curAgg
					}
					a.curIdx++
					a.scratch.curAgg = zeroDecimalValue

				}

				var isNull bool
				isNull = false
				if !isNull {

					{

						tmpDec := &_overloadHelper.tmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curAgg, &a.scratch.curAgg, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	} else {
		sel = sel[:inputLen]
		if nulls.MaybeHasNulls() {
			for _, i := range sel {

				if groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.nulls.SetNull(a.curIdx)
					} else {
						a.scratch.vec[a.curIdx] = a.scratch.curAgg
					}
					a.curIdx++
					a.scratch.curAgg = zeroDecimalValue

					a.scratch.foundNonNullForCurrentGroup = false
				}

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{

						tmpDec := &_overloadHelper.tmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curAgg, &a.scratch.curAgg, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			for _, i := range sel {

				if groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.nulls.SetNull(a.curIdx)
					} else {
						a.scratch.vec[a.curIdx] = a.scratch.curAgg
					}
					a.curIdx++
					a.scratch.curAgg = zeroDecimalValue

				}

				var isNull bool
				isNull = false
				if !isNull {

					{

						tmpDec := &_overloadHelper.tmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curAgg, &a.scratch.curAgg, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	}
}

func (a *sumInt64OrderedAgg) Flush(outputIdx int) {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// null.
	// Go around "argument overwritten before first use" linter error.
	_ = outputIdx
	outputIdx = a.curIdx
	a.curIdx++
	if !a.scratch.foundNonNullForCurrentGroup {
		a.nulls.SetNull(outputIdx)
	} else {
		a.scratch.vec[outputIdx] = a.scratch.curAgg
	}
}

type sumInt64OrderedAggAlloc struct {
	aggAllocBase
	aggFuncs []sumInt64OrderedAgg
}

var _ aggregateFuncAlloc = &sumInt64OrderedAggAlloc{}

func (a *sumInt64OrderedAggAlloc) newAggFunc() aggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(sizeOfSumInt64OrderedAgg * a.allocSize)
		a.aggFuncs = make([]sumInt64OrderedAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	a.aggFuncs = a.aggFuncs[1:]
	return f
}

type sumDecimalOrderedAgg struct {
	orderedAggregateFuncBase
	scratch struct {
		// curAgg holds the running total, so we can index into the slice once per
		// group, instead of on each iteration.
		curAgg apd.Decimal
		// vec points to the output vector we are updating.
		vec []apd.Decimal
		// foundNonNullForCurrentGroup tracks if we have seen any non-null values
		// for the group that is currently being aggregated.
		foundNonNullForCurrentGroup bool
	}
}

var _ aggregateFunc = &sumDecimalOrderedAgg{}

const sizeOfSumDecimalOrderedAgg = int64(unsafe.Sizeof(sumDecimalOrderedAgg{}))

func (a *sumDecimalOrderedAgg) Init(groups []bool, vec coldata.Vec) {
	a.orderedAggregateFuncBase.Init(groups, vec)
	a.scratch.vec = vec.Decimal()
	a.Reset()
}

func (a *sumDecimalOrderedAgg) Reset() {
	a.orderedAggregateFuncBase.Reset()
	a.scratch.foundNonNullForCurrentGroup = false
}

func (a *sumDecimalOrderedAgg) Compute(
	vecs []coldata.Vec, inputIdxs []uint32, inputLen int, sel []int,
) {
	vec := vecs[inputIdxs[0]]
	col, nulls := vec.Decimal(), vec.Nulls()
	groups := a.groups
	if sel == nil {
		_ = groups[inputLen-1]
		col = col[:inputLen]
		if nulls.MaybeHasNulls() {
			for i := range col {

				if groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.nulls.SetNull(a.curIdx)
					} else {
						a.scratch.vec[a.curIdx] = a.scratch.curAgg
					}
					a.curIdx++
					a.scratch.curAgg = zeroDecimalValue

					a.scratch.foundNonNullForCurrentGroup = false
				}

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{

						_, err := tree.ExactCtx.Add(&a.scratch.curAgg, &a.scratch.curAgg, &col[i])
						if err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			for i := range col {

				if groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.nulls.SetNull(a.curIdx)
					} else {
						a.scratch.vec[a.curIdx] = a.scratch.curAgg
					}
					a.curIdx++
					a.scratch.curAgg = zeroDecimalValue

				}

				var isNull bool
				isNull = false
				if !isNull {

					{

						_, err := tree.ExactCtx.Add(&a.scratch.curAgg, &a.scratch.curAgg, &col[i])
						if err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	} else {
		sel = sel[:inputLen]
		if nulls.MaybeHasNulls() {
			for _, i := range sel {

				if groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.nulls.SetNull(a.curIdx)
					} else {
						a.scratch.vec[a.curIdx] = a.scratch.curAgg
					}
					a.curIdx++
					a.scratch.curAgg = zeroDecimalValue

					a.scratch.foundNonNullForCurrentGroup = false
				}

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{

						_, err := tree.ExactCtx.Add(&a.scratch.curAgg, &a.scratch.curAgg, &col[i])
						if err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			for _, i := range sel {

				if groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.nulls.SetNull(a.curIdx)
					} else {
						a.scratch.vec[a.curIdx] = a.scratch.curAgg
					}
					a.curIdx++
					a.scratch.curAgg = zeroDecimalValue

				}

				var isNull bool
				isNull = false
				if !isNull {

					{

						_, err := tree.ExactCtx.Add(&a.scratch.curAgg, &a.scratch.curAgg, &col[i])
						if err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	}
}

func (a *sumDecimalOrderedAgg) Flush(outputIdx int) {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// null.
	// Go around "argument overwritten before first use" linter error.
	_ = outputIdx
	outputIdx = a.curIdx
	a.curIdx++
	if !a.scratch.foundNonNullForCurrentGroup {
		a.nulls.SetNull(outputIdx)
	} else {
		a.scratch.vec[outputIdx] = a.scratch.curAgg
	}
}

type sumDecimalOrderedAggAlloc struct {
	aggAllocBase
	aggFuncs []sumDecimalOrderedAgg
}

var _ aggregateFuncAlloc = &sumDecimalOrderedAggAlloc{}

func (a *sumDecimalOrderedAggAlloc) newAggFunc() aggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(sizeOfSumDecimalOrderedAgg * a.allocSize)
		a.aggFuncs = make([]sumDecimalOrderedAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	a.aggFuncs = a.aggFuncs[1:]
	return f
}

type sumFloat64OrderedAgg struct {
	orderedAggregateFuncBase
	scratch struct {
		// curAgg holds the running total, so we can index into the slice once per
		// group, instead of on each iteration.
		curAgg float64
		// vec points to the output vector we are updating.
		vec []float64
		// foundNonNullForCurrentGroup tracks if we have seen any non-null values
		// for the group that is currently being aggregated.
		foundNonNullForCurrentGroup bool
	}
}

var _ aggregateFunc = &sumFloat64OrderedAgg{}

const sizeOfSumFloat64OrderedAgg = int64(unsafe.Sizeof(sumFloat64OrderedAgg{}))

func (a *sumFloat64OrderedAgg) Init(groups []bool, vec coldata.Vec) {
	a.orderedAggregateFuncBase.Init(groups, vec)
	a.scratch.vec = vec.Float64()
	a.Reset()
}

func (a *sumFloat64OrderedAgg) Reset() {
	a.orderedAggregateFuncBase.Reset()
	a.scratch.foundNonNullForCurrentGroup = false
}

func (a *sumFloat64OrderedAgg) Compute(
	vecs []coldata.Vec, inputIdxs []uint32, inputLen int, sel []int,
) {
	vec := vecs[inputIdxs[0]]
	col, nulls := vec.Float64(), vec.Nulls()
	groups := a.groups
	if sel == nil {
		_ = groups[inputLen-1]
		col = col[:inputLen]
		if nulls.MaybeHasNulls() {
			for i := range col {

				if groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.nulls.SetNull(a.curIdx)
					} else {
						a.scratch.vec[a.curIdx] = a.scratch.curAgg
					}
					a.curIdx++
					a.scratch.curAgg = zeroFloat64Value

					a.scratch.foundNonNullForCurrentGroup = false
				}

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{

						a.scratch.curAgg = float64(a.scratch.curAgg) + float64(col[i])
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			for i := range col {

				if groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.nulls.SetNull(a.curIdx)
					} else {
						a.scratch.vec[a.curIdx] = a.scratch.curAgg
					}
					a.curIdx++
					a.scratch.curAgg = zeroFloat64Value

				}

				var isNull bool
				isNull = false
				if !isNull {

					{

						a.scratch.curAgg = float64(a.scratch.curAgg) + float64(col[i])
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	} else {
		sel = sel[:inputLen]
		if nulls.MaybeHasNulls() {
			for _, i := range sel {

				if groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.nulls.SetNull(a.curIdx)
					} else {
						a.scratch.vec[a.curIdx] = a.scratch.curAgg
					}
					a.curIdx++
					a.scratch.curAgg = zeroFloat64Value

					a.scratch.foundNonNullForCurrentGroup = false
				}

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{

						a.scratch.curAgg = float64(a.scratch.curAgg) + float64(col[i])
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			for _, i := range sel {

				if groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.nulls.SetNull(a.curIdx)
					} else {
						a.scratch.vec[a.curIdx] = a.scratch.curAgg
					}
					a.curIdx++
					a.scratch.curAgg = zeroFloat64Value

				}

				var isNull bool
				isNull = false
				if !isNull {

					{

						a.scratch.curAgg = float64(a.scratch.curAgg) + float64(col[i])
					}

					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	}
}

func (a *sumFloat64OrderedAgg) Flush(outputIdx int) {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// null.
	// Go around "argument overwritten before first use" linter error.
	_ = outputIdx
	outputIdx = a.curIdx
	a.curIdx++
	if !a.scratch.foundNonNullForCurrentGroup {
		a.nulls.SetNull(outputIdx)
	} else {
		a.scratch.vec[outputIdx] = a.scratch.curAgg
	}
}

type sumFloat64OrderedAggAlloc struct {
	aggAllocBase
	aggFuncs []sumFloat64OrderedAgg
}

var _ aggregateFuncAlloc = &sumFloat64OrderedAggAlloc{}

func (a *sumFloat64OrderedAggAlloc) newAggFunc() aggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(sizeOfSumFloat64OrderedAgg * a.allocSize)
		a.aggFuncs = make([]sumFloat64OrderedAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	a.aggFuncs = a.aggFuncs[1:]
	return f
}

type sumIntervalOrderedAgg struct {
	orderedAggregateFuncBase
	scratch struct {
		// curAgg holds the running total, so we can index into the slice once per
		// group, instead of on each iteration.
		curAgg duration.Duration
		// vec points to the output vector we are updating.
		vec []duration.Duration
		// foundNonNullForCurrentGroup tracks if we have seen any non-null values
		// for the group that is currently being aggregated.
		foundNonNullForCurrentGroup bool
	}
}

var _ aggregateFunc = &sumIntervalOrderedAgg{}

const sizeOfSumIntervalOrderedAgg = int64(unsafe.Sizeof(sumIntervalOrderedAgg{}))

func (a *sumIntervalOrderedAgg) Init(groups []bool, vec coldata.Vec) {
	a.orderedAggregateFuncBase.Init(groups, vec)
	a.scratch.vec = vec.Interval()
	a.Reset()
}

func (a *sumIntervalOrderedAgg) Reset() {
	a.orderedAggregateFuncBase.Reset()
	a.scratch.foundNonNullForCurrentGroup = false
}

func (a *sumIntervalOrderedAgg) Compute(
	vecs []coldata.Vec, inputIdxs []uint32, inputLen int, sel []int,
) {
	vec := vecs[inputIdxs[0]]
	col, nulls := vec.Interval(), vec.Nulls()
	groups := a.groups
	if sel == nil {
		_ = groups[inputLen-1]
		col = col[:inputLen]
		if nulls.MaybeHasNulls() {
			for i := range col {

				if groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.nulls.SetNull(a.curIdx)
					} else {
						a.scratch.vec[a.curIdx] = a.scratch.curAgg
					}
					a.curIdx++
					a.scratch.curAgg = zeroIntervalValue

					a.scratch.foundNonNullForCurrentGroup = false
				}

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {
					a.scratch.curAgg = a.scratch.curAgg.Add(col[i])
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			for i := range col {

				if groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.nulls.SetNull(a.curIdx)
					} else {
						a.scratch.vec[a.curIdx] = a.scratch.curAgg
					}
					a.curIdx++
					a.scratch.curAgg = zeroIntervalValue

				}

				var isNull bool
				isNull = false
				if !isNull {
					a.scratch.curAgg = a.scratch.curAgg.Add(col[i])
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	} else {
		sel = sel[:inputLen]
		if nulls.MaybeHasNulls() {
			for _, i := range sel {

				if groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.nulls.SetNull(a.curIdx)
					} else {
						a.scratch.vec[a.curIdx] = a.scratch.curAgg
					}
					a.curIdx++
					a.scratch.curAgg = zeroIntervalValue

					a.scratch.foundNonNullForCurrentGroup = false
				}

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {
					a.scratch.curAgg = a.scratch.curAgg.Add(col[i])
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			for _, i := range sel {

				if groups[i] {
					// If we encounter a new group, and we haven't found any non-nulls for the
					// current group, the output for this group should be null.
					if !a.scratch.foundNonNullForCurrentGroup {
						a.nulls.SetNull(a.curIdx)
					} else {
						a.scratch.vec[a.curIdx] = a.scratch.curAgg
					}
					a.curIdx++
					a.scratch.curAgg = zeroIntervalValue

				}

				var isNull bool
				isNull = false
				if !isNull {
					a.scratch.curAgg = a.scratch.curAgg.Add(col[i])
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	}
}

func (a *sumIntervalOrderedAgg) Flush(outputIdx int) {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// null.
	// Go around "argument overwritten before first use" linter error.
	_ = outputIdx
	outputIdx = a.curIdx
	a.curIdx++
	if !a.scratch.foundNonNullForCurrentGroup {
		a.nulls.SetNull(outputIdx)
	} else {
		a.scratch.vec[outputIdx] = a.scratch.curAgg
	}
}

type sumIntervalOrderedAggAlloc struct {
	aggAllocBase
	aggFuncs []sumIntervalOrderedAgg
}

var _ aggregateFuncAlloc = &sumIntervalOrderedAggAlloc{}

func (a *sumIntervalOrderedAggAlloc) newAggFunc() aggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(sizeOfSumIntervalOrderedAgg * a.allocSize)
		a.aggFuncs = make([]sumIntervalOrderedAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	a.aggFuncs = a.aggFuncs[1:]
	return f
}
