// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package variant // import "go.opentelemetry.io/otel/attribute"

import (
	"reflect"
	"unsafe"

	"github.com/bogdandrutu/lazyproto/variant/internal"
)

const totalBitCount = 8 * unsafe.Sizeof(int64(0))
const capFieldBitCount = totalBitCount / 2
const capFieldMask = (1 << capFieldBitCount) - 1

// MaxSliceLen is the maximum length of a slice-type that can be stored in Variant. The length of Go slices
// can be at most maxint, however Variant is not able to store lengths of maxint. Len field
// in Variant uses lenFieldShiftCount bits less than int, i.e. the maximum length of a slice
// stored in Variant is maxint / (2^lenFieldShiftCount), which we calculate below.
const MaxSliceLen = int((^uint(0))>>1) >> (totalBitCount / 2)

// Type describes the type of the data Value holds.
type Type int

// Value represents the value part in key-value pairs.
type Value struct {
	ptr  unsafe.Pointer
	bits uint64
}

// BoolValue creates a BOOL Value.
func BoolValue(v bool) Value {
	return Value{
		bits: internal.BoolToRaw(v),
	}
}

// Int64Value creates an INT64 Value.
func Int64Value(v int64) Value {
	return Value{
		bits: internal.Int64ToRaw(v),
	}
}

// Float64Value creates a FLOAT64 Value.
func Float64Value(v float64) Value {
	return Value{
		bits: internal.Float64ToRaw(v),
	}
}

// StringValue creates a STRING Value.
func StringValue(v string) Value {
	sdr := (*reflect.StringHeader)(unsafe.Pointer(&v))

	return Value{
		ptr:  unsafe.Pointer(sdr.Data),
		bits: uint64(sdr.Len),
	}
}

// PointerValue creates a POINTER Value.
func PointerValue(v unsafe.Pointer) Value {
	return Value{
		ptr: v,
	}
}

// BytesValue creates a BYTES Value.
func BytesValue(v []byte) Value {
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&v))
	if hdr.Len > MaxSliceLen {
		panic("maximum len exceeded")
	}

	return Value{
		ptr:  unsafe.Pointer(hdr.Data),
		bits: uint64(hdr.Len<<capFieldBitCount) | uint64(hdr.Cap),
	}
}

// AsBool returns the bool value.
func (v Value) AsBool() bool {
	return internal.RawToBool(v.bits)
}

// AsInt64 returns the int64 value.
func (v Value) AsInt64() int64 {
	return internal.RawToInt64(v.bits)
}

// AsFloat64 returns the float64 value.
func (v Value) AsFloat64() float64 {
	return internal.RawToFloat64(v.bits)
}

// AsString returns the string value.
func (v Value) AsString() string {
	var s string
	dest := (*reflect.StringHeader)(unsafe.Pointer(&s))
	dest.Data = uintptr(v.ptr)
	dest.Len = int(v.bits)
	return s
}

// AsPointer returns the pinter value.
func (v Value) AsPointer() unsafe.Pointer {
	return v.ptr
}

// AsBytes returns the []byte value.
func (v Value) AsBytes() []byte {
	var b []byte
	dest := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	dest.Data = uintptr(v.ptr)
	dest.Len = int(v.bits >> capFieldBitCount)
	dest.Cap = int((v.bits >> capFieldBitCount) & capFieldMask)
	return b
}
