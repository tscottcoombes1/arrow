// Code generated by tensor/numeric.gen.go.tmpl. DO NOT EDIT.

// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tensor

import (
	"github.com/apache/arrow/go/v18/arrow"
)

// Int8 is an n-dim array of int8s.
type Int8 struct {
	tensorBase
	values []int8
}

// NewInt8 returns a new n-dimensional array of int8s.
// If strides is nil, row-major strides will be inferred.
// If names is nil, a slice of empty strings will be created.
func NewInt8(data arrow.ArrayData, shape, strides []int64, names []string) *Int8 {
	tsr := &Int8{tensorBase: *newTensor(arrow.PrimitiveTypes.Int8, data, shape, strides, names)}
	vals := tsr.data.Buffers()[1]
	if vals != nil {
		tsr.values = arrow.Int8Traits.CastFromBytes(vals.Bytes())
		beg := tsr.data.Offset()
		end := beg + tsr.data.Len()
		tsr.values = tsr.values[beg:end]
	}
	return tsr
}

func (tsr *Int8) Value(i []int64) int8 { j := int(tsr.offset(i)); return tsr.values[j] }
func (tsr *Int8) Int8Values() []int8   { return tsr.values }

// Int16 is an n-dim array of int16s.
type Int16 struct {
	tensorBase
	values []int16
}

// NewInt16 returns a new n-dimensional array of int16s.
// If strides is nil, row-major strides will be inferred.
// If names is nil, a slice of empty strings will be created.
func NewInt16(data arrow.ArrayData, shape, strides []int64, names []string) *Int16 {
	tsr := &Int16{tensorBase: *newTensor(arrow.PrimitiveTypes.Int16, data, shape, strides, names)}
	vals := tsr.data.Buffers()[1]
	if vals != nil {
		tsr.values = arrow.Int16Traits.CastFromBytes(vals.Bytes())
		beg := tsr.data.Offset()
		end := beg + tsr.data.Len()
		tsr.values = tsr.values[beg:end]
	}
	return tsr
}

func (tsr *Int16) Value(i []int64) int16 { j := int(tsr.offset(i)); return tsr.values[j] }
func (tsr *Int16) Int16Values() []int16  { return tsr.values }

// Int32 is an n-dim array of int32s.
type Int32 struct {
	tensorBase
	values []int32
}

// NewInt32 returns a new n-dimensional array of int32s.
// If strides is nil, row-major strides will be inferred.
// If names is nil, a slice of empty strings will be created.
func NewInt32(data arrow.ArrayData, shape, strides []int64, names []string) *Int32 {
	tsr := &Int32{tensorBase: *newTensor(arrow.PrimitiveTypes.Int32, data, shape, strides, names)}
	vals := tsr.data.Buffers()[1]
	if vals != nil {
		tsr.values = arrow.Int32Traits.CastFromBytes(vals.Bytes())
		beg := tsr.data.Offset()
		end := beg + tsr.data.Len()
		tsr.values = tsr.values[beg:end]
	}
	return tsr
}

func (tsr *Int32) Value(i []int64) int32 { j := int(tsr.offset(i)); return tsr.values[j] }
func (tsr *Int32) Int32Values() []int32  { return tsr.values }

// Int64 is an n-dim array of int64s.
type Int64 struct {
	tensorBase
	values []int64
}

// NewInt64 returns a new n-dimensional array of int64s.
// If strides is nil, row-major strides will be inferred.
// If names is nil, a slice of empty strings will be created.
func NewInt64(data arrow.ArrayData, shape, strides []int64, names []string) *Int64 {
	tsr := &Int64{tensorBase: *newTensor(arrow.PrimitiveTypes.Int64, data, shape, strides, names)}
	vals := tsr.data.Buffers()[1]
	if vals != nil {
		tsr.values = arrow.Int64Traits.CastFromBytes(vals.Bytes())
		beg := tsr.data.Offset()
		end := beg + tsr.data.Len()
		tsr.values = tsr.values[beg:end]
	}
	return tsr
}

func (tsr *Int64) Value(i []int64) int64 { j := int(tsr.offset(i)); return tsr.values[j] }
func (tsr *Int64) Int64Values() []int64  { return tsr.values }

// Uint8 is an n-dim array of uint8s.
type Uint8 struct {
	tensorBase
	values []uint8
}

// NewUint8 returns a new n-dimensional array of uint8s.
// If strides is nil, row-major strides will be inferred.
// If names is nil, a slice of empty strings will be created.
func NewUint8(data arrow.ArrayData, shape, strides []int64, names []string) *Uint8 {
	tsr := &Uint8{tensorBase: *newTensor(arrow.PrimitiveTypes.Uint8, data, shape, strides, names)}
	vals := tsr.data.Buffers()[1]
	if vals != nil {
		tsr.values = arrow.Uint8Traits.CastFromBytes(vals.Bytes())
		beg := tsr.data.Offset()
		end := beg + tsr.data.Len()
		tsr.values = tsr.values[beg:end]
	}
	return tsr
}

func (tsr *Uint8) Value(i []int64) uint8 { j := int(tsr.offset(i)); return tsr.values[j] }
func (tsr *Uint8) Uint8Values() []uint8  { return tsr.values }

// Uint16 is an n-dim array of uint16s.
type Uint16 struct {
	tensorBase
	values []uint16
}

// NewUint16 returns a new n-dimensional array of uint16s.
// If strides is nil, row-major strides will be inferred.
// If names is nil, a slice of empty strings will be created.
func NewUint16(data arrow.ArrayData, shape, strides []int64, names []string) *Uint16 {
	tsr := &Uint16{tensorBase: *newTensor(arrow.PrimitiveTypes.Uint16, data, shape, strides, names)}
	vals := tsr.data.Buffers()[1]
	if vals != nil {
		tsr.values = arrow.Uint16Traits.CastFromBytes(vals.Bytes())
		beg := tsr.data.Offset()
		end := beg + tsr.data.Len()
		tsr.values = tsr.values[beg:end]
	}
	return tsr
}

func (tsr *Uint16) Value(i []int64) uint16 { j := int(tsr.offset(i)); return tsr.values[j] }
func (tsr *Uint16) Uint16Values() []uint16 { return tsr.values }

// Uint32 is an n-dim array of uint32s.
type Uint32 struct {
	tensorBase
	values []uint32
}

// NewUint32 returns a new n-dimensional array of uint32s.
// If strides is nil, row-major strides will be inferred.
// If names is nil, a slice of empty strings will be created.
func NewUint32(data arrow.ArrayData, shape, strides []int64, names []string) *Uint32 {
	tsr := &Uint32{tensorBase: *newTensor(arrow.PrimitiveTypes.Uint32, data, shape, strides, names)}
	vals := tsr.data.Buffers()[1]
	if vals != nil {
		tsr.values = arrow.Uint32Traits.CastFromBytes(vals.Bytes())
		beg := tsr.data.Offset()
		end := beg + tsr.data.Len()
		tsr.values = tsr.values[beg:end]
	}
	return tsr
}

func (tsr *Uint32) Value(i []int64) uint32 { j := int(tsr.offset(i)); return tsr.values[j] }
func (tsr *Uint32) Uint32Values() []uint32 { return tsr.values }

// Uint64 is an n-dim array of uint64s.
type Uint64 struct {
	tensorBase
	values []uint64
}

// NewUint64 returns a new n-dimensional array of uint64s.
// If strides is nil, row-major strides will be inferred.
// If names is nil, a slice of empty strings will be created.
func NewUint64(data arrow.ArrayData, shape, strides []int64, names []string) *Uint64 {
	tsr := &Uint64{tensorBase: *newTensor(arrow.PrimitiveTypes.Uint64, data, shape, strides, names)}
	vals := tsr.data.Buffers()[1]
	if vals != nil {
		tsr.values = arrow.Uint64Traits.CastFromBytes(vals.Bytes())
		beg := tsr.data.Offset()
		end := beg + tsr.data.Len()
		tsr.values = tsr.values[beg:end]
	}
	return tsr
}

func (tsr *Uint64) Value(i []int64) uint64 { j := int(tsr.offset(i)); return tsr.values[j] }
func (tsr *Uint64) Uint64Values() []uint64 { return tsr.values }

// Float32 is an n-dim array of float32s.
type Float32 struct {
	tensorBase
	values []float32
}

// NewFloat32 returns a new n-dimensional array of float32s.
// If strides is nil, row-major strides will be inferred.
// If names is nil, a slice of empty strings will be created.
func NewFloat32(data arrow.ArrayData, shape, strides []int64, names []string) *Float32 {
	tsr := &Float32{tensorBase: *newTensor(arrow.PrimitiveTypes.Float32, data, shape, strides, names)}
	vals := tsr.data.Buffers()[1]
	if vals != nil {
		tsr.values = arrow.Float32Traits.CastFromBytes(vals.Bytes())
		beg := tsr.data.Offset()
		end := beg + tsr.data.Len()
		tsr.values = tsr.values[beg:end]
	}
	return tsr
}

func (tsr *Float32) Value(i []int64) float32  { j := int(tsr.offset(i)); return tsr.values[j] }
func (tsr *Float32) Float32Values() []float32 { return tsr.values }

// Float64 is an n-dim array of float64s.
type Float64 struct {
	tensorBase
	values []float64
}

// NewFloat64 returns a new n-dimensional array of float64s.
// If strides is nil, row-major strides will be inferred.
// If names is nil, a slice of empty strings will be created.
func NewFloat64(data arrow.ArrayData, shape, strides []int64, names []string) *Float64 {
	tsr := &Float64{tensorBase: *newTensor(arrow.PrimitiveTypes.Float64, data, shape, strides, names)}
	vals := tsr.data.Buffers()[1]
	if vals != nil {
		tsr.values = arrow.Float64Traits.CastFromBytes(vals.Bytes())
		beg := tsr.data.Offset()
		end := beg + tsr.data.Len()
		tsr.values = tsr.values[beg:end]
	}
	return tsr
}

func (tsr *Float64) Value(i []int64) float64  { j := int(tsr.offset(i)); return tsr.values[j] }
func (tsr *Float64) Float64Values() []float64 { return tsr.values }

// Date32 is an n-dim array of date32s.
type Date32 struct {
	tensorBase
	values []arrow.Date32
}

// NewDate32 returns a new n-dimensional array of date32s.
// If strides is nil, row-major strides will be inferred.
// If names is nil, a slice of empty strings will be created.
func NewDate32(data arrow.ArrayData, shape, strides []int64, names []string) *Date32 {
	tsr := &Date32{tensorBase: *newTensor(arrow.PrimitiveTypes.Date32, data, shape, strides, names)}
	vals := tsr.data.Buffers()[1]
	if vals != nil {
		tsr.values = arrow.Date32Traits.CastFromBytes(vals.Bytes())
		beg := tsr.data.Offset()
		end := beg + tsr.data.Len()
		tsr.values = tsr.values[beg:end]
	}
	return tsr
}

func (tsr *Date32) Value(i []int64) arrow.Date32 { j := int(tsr.offset(i)); return tsr.values[j] }
func (tsr *Date32) Date32Values() []arrow.Date32 { return tsr.values }

// Date64 is an n-dim array of date64s.
type Date64 struct {
	tensorBase
	values []arrow.Date64
}

// NewDate64 returns a new n-dimensional array of date64s.
// If strides is nil, row-major strides will be inferred.
// If names is nil, a slice of empty strings will be created.
func NewDate64(data arrow.ArrayData, shape, strides []int64, names []string) *Date64 {
	tsr := &Date64{tensorBase: *newTensor(arrow.PrimitiveTypes.Date64, data, shape, strides, names)}
	vals := tsr.data.Buffers()[1]
	if vals != nil {
		tsr.values = arrow.Date64Traits.CastFromBytes(vals.Bytes())
		beg := tsr.data.Offset()
		end := beg + tsr.data.Len()
		tsr.values = tsr.values[beg:end]
	}
	return tsr
}

func (tsr *Date64) Value(i []int64) arrow.Date64 { j := int(tsr.offset(i)); return tsr.values[j] }
func (tsr *Date64) Date64Values() []arrow.Date64 { return tsr.values }

var (
	_ Interface = (*Int8)(nil)
	_ Interface = (*Int16)(nil)
	_ Interface = (*Int32)(nil)
	_ Interface = (*Int64)(nil)
	_ Interface = (*Uint8)(nil)
	_ Interface = (*Uint16)(nil)
	_ Interface = (*Uint32)(nil)
	_ Interface = (*Uint64)(nil)
	_ Interface = (*Float32)(nil)
	_ Interface = (*Float64)(nil)
	_ Interface = (*Date32)(nil)
	_ Interface = (*Date64)(nil)
)
