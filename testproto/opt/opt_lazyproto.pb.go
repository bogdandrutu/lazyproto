// Code generated by protoc-gen-go_lazyproto. DO NOT EDIT.
// protoc-gen-go_lazyproto version: (devel)
// source: testproto/opt/opt.proto

package opt

import (
	binary "encoding/binary"
	fmt "fmt"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	math "math"
	bits "math/bits"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OptionalEnum int32

const (
	OptionalEnumUnspecified OptionalEnum = 0
	OptionalEnumValue1      OptionalEnum = 1
	OptionalEnumValue2      OptionalEnum = 2
	OptionalEnumValue3      OptionalEnum = 3
)

type OptionalMessage struct {
	optionalValue *OptionalValue
}

type OptionalValue struct {
	optionalInt32    *int32
	optionalInt64    *int64
	optionalUint32   *uint32
	optionalUint64   *uint64
	optionalSint32   *int32
	optionalSint64   *int64
	optionalFixed32  *uint32
	optionalFixed64  *uint64
	optionalSfixed32 *int32
	optionalSfixed64 *int64
	optionalFloat    *float32
	optionalDouble   *float64
	optionalBool     *bool
	optionalString   *string
	optionalBytes    []byte
	optionalEnum     *OptionalEnum
}

func (m *OptionalMessage) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OptionalMessage) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *OptionalMessage) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	if m.optionalValue != nil {
		size, err := m.optionalValue.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OptionalValue) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OptionalValue) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *OptionalValue) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	if m.optionalEnum != nil {
		i = encodeVarint(dAtA, i, uint64(*m.optionalEnum))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.optionalBytes != nil {
		if len(m.optionalBytes) > 0 {
			i -= len(m.optionalBytes)
			copy(dAtA[i:], m.optionalBytes)
			i = encodeVarint(dAtA, i, uint64(len(m.optionalBytes)))
			i--
			dAtA[i] = 0x7a
		}
	}
	if m.optionalString != nil {
		i -= len(*m.optionalString)
		copy(dAtA[i:], *m.optionalString)
		i = encodeVarint(dAtA, i, uint64(len(*m.optionalString)))
		i--
		dAtA[i] = 0x72
	}
	if m.optionalBool != nil {
		i--
		if *m.optionalBool {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x68
	}
	if m.optionalDouble != nil {
		i -= 8
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(*m.optionalDouble))))
		i--
		dAtA[i] = 0x61
	}
	if m.optionalFloat != nil {
		i -= 4
		binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(*m.optionalFloat))))
		i--
		dAtA[i] = 0x5d
	}
	if m.optionalSfixed64 != nil {
		i -= 8
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(*m.optionalSfixed64))
		i--
		dAtA[i] = 0x51
	}
	if m.optionalSfixed32 != nil {
		i -= 4
		binary.LittleEndian.PutUint32(dAtA[i:], uint32(*m.optionalSfixed32))
		i--
		dAtA[i] = 0x4d
	}
	if m.optionalFixed64 != nil {
		i -= 8
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(*m.optionalFixed64))
		i--
		dAtA[i] = 0x41
	}
	if m.optionalFixed32 != nil {
		i -= 4
		binary.LittleEndian.PutUint32(dAtA[i:], uint32(*m.optionalFixed32))
		i--
		dAtA[i] = 0x3d
	}
	if m.optionalSint64 != nil {
		i = encodeVarint(dAtA, i, uint64((uint64(*m.optionalSint64)<<1)^uint64((*m.optionalSint64>>63))))
		i--
		dAtA[i] = 0x30
	}
	if m.optionalSint32 != nil {
		i = encodeVarint(dAtA, i, uint64((uint32(*m.optionalSint32)<<1)^uint32((*m.optionalSint32>>31))))
		i--
		dAtA[i] = 0x28
	}
	if m.optionalUint64 != nil {
		i = encodeVarint(dAtA, i, uint64(*m.optionalUint64))
		i--
		dAtA[i] = 0x20
	}
	if m.optionalUint32 != nil {
		i = encodeVarint(dAtA, i, uint64(*m.optionalUint32))
		i--
		dAtA[i] = 0x18
	}
	if m.optionalInt64 != nil {
		i = encodeVarint(dAtA, i, uint64(*m.optionalInt64))
		i--
		dAtA[i] = 0x10
	}
	if m.optionalInt32 != nil {
		i = encodeVarint(dAtA, i, uint64(*m.optionalInt32))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarint(dAtA []byte, offset int, v uint64) int {
	offset -= sov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OptionalMessage) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.optionalValue != nil {
		l = m.optionalValue.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}

func (m *OptionalValue) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.optionalInt32 != nil {
		n += 1 + sov(uint64(*m.optionalInt32))
	}
	if m.optionalInt64 != nil {
		n += 1 + sov(uint64(*m.optionalInt64))
	}
	if m.optionalUint32 != nil {
		n += 1 + sov(uint64(*m.optionalUint32))
	}
	if m.optionalUint64 != nil {
		n += 1 + sov(uint64(*m.optionalUint64))
	}
	if m.optionalSint32 != nil {
		n += 1 + soz(uint64(*m.optionalSint32))
	}
	if m.optionalSint64 != nil {
		n += 1 + soz(uint64(*m.optionalSint64))
	}
	if m.optionalFixed32 != nil {
		n += 5
	}
	if m.optionalFixed64 != nil {
		n += 9
	}
	if m.optionalSfixed32 != nil {
		n += 5
	}
	if m.optionalSfixed64 != nil {
		n += 9
	}
	if m.optionalFloat != nil {
		n += 5
	}
	if m.optionalDouble != nil {
		n += 9
	}
	if m.optionalBool != nil {
		n += 2
	}
	if m.optionalString != nil {
		l = len(*m.optionalString)
		n += 1 + l + sov(uint64(l))
	}
	if m.optionalBytes != nil {
		l = len(m.optionalBytes)
		if l > 0 {
			n += 1 + l + sov(uint64(l))
		}
	}
	if m.optionalEnum != nil {
		n += 2 + sov(uint64(*m.optionalEnum))
	}
	return n
}

func sov(x uint64) (n int) {
	return (bits.Len64(x|1) + 6) / 7
}
func soz(x uint64) (n int) {
	return sov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OptionalMessage) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OptionalMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OptionalMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field optionalValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.optionalValue.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OptionalValue) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OptionalValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OptionalValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field optionalInt32", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.optionalInt32 = &v
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field optionalInt64", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.optionalInt64 = &v
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field optionalUint32", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.optionalUint32 = &v
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field optionalUint64", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.optionalUint64 = &v
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field optionalSint32", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31))
			m.optionalSint32 = &v
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field optionalSint64", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
			v2 := int64(v)
			m.optionalSint64 = &v2
		case 7:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field optionalFixed32", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.optionalFixed32 = &v
		case 8:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field optionalFixed64", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.optionalFixed64 = &v
		case 9:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field optionalSfixed32", wireType)
			}
			var v int32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = int32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.optionalSfixed32 = &v
		case 10:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field optionalSfixed64", wireType)
			}
			var v int64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = int64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.optionalSfixed64 = &v
		case 11:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field optionalFloat", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			v2 := float32(math.Float32frombits(v))
			m.optionalFloat = &v2
		case 12:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field optionalDouble", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			v2 := float64(math.Float64frombits(v))
			m.optionalDouble = &v2
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field optionalBool", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.optionalBool = &b
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field optionalString", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.optionalString = &s
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field optionalBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.optionalBytes = append(m.optionalBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.optionalBytes == nil {
				m.optionalBytes = []byte{}
			}
			iNdEx = postIndex
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field optionalEnum", wireType)
			}
			var v OptionalEnum
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= OptionalEnum(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.optionalEnum = &v
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skip(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflow
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflow
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflow
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLength
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroup
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLength
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLength        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflow          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroup = fmt.Errorf("proto: unexpected end of group")
)
