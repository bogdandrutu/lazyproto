// Copyright (c) 2021 PlanetScale Inc. All rights reserved.
// Copyright (c) 2013, The GoGo Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package size

import (
	"strconv"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/bogdandrutu/lazyproto/generator"
)

func init() {
	generator.RegisterFeature("size", func(gen *generator.GeneratedFile) generator.FeatureGenerator {
		return &size{GeneratedFile: gen}
	})
}

type size struct {
	*generator.GeneratedFile
	once bool
}

var _ generator.FeatureGenerator = (*size)(nil)

func (p *size) Name() string {
	return "size"
}

func (p *size) GenerateFile(file *protogen.File) bool {
	for _, message := range file.Messages {
		p.message(message)
	}

	return p.once
}

func (p *size) GenerateHelpers() {
	p.P(`
	func sov(x uint64) (n int) {
                return (`, p.Ident("math/bits", "Len64"), `(x | 1) + 6)/ 7
	}`)
	p.P(`func soz(x uint64) (n int) {
		return sov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
	}`)
}

func (p *size) messageSize(varName, sizeName string, message *protogen.Message) {
	local := p.IsLocalMessage(message)

	if local {
		p.P(`l = `, varName, `.`, sizeName, `()`)
	} else {
		p.P(`if size, ok := interface{}(`, varName, `).(interface{`)
		p.P(sizeName, `() int`)
		p.P(`}); ok{`)
		p.P(`l = size.`, sizeName, `()`)
		p.P(`} else {`)
		p.P(`l = `, p.Ident(generator.ProtoPkg, "Size"), `(`, varName, `)`)
		p.P(`}`)
	}
}

func (p *size) field(oneof bool, field *protogen.Field, sizeName string) {
	fieldname := field.GoName
	nullable := p.IsNullable(field)
	repeated := field.Desc.Cardinality() == protoreflect.Repeated
	if repeated {
		p.P(`if len(m.`, fieldname, `) > 0 {`)
	} else if nullable && !oneof {
		p.P(`if m.`, fieldname, ` != nil {`)
	}
	packed := field.Desc.IsPacked()
	wireType := generator.ProtoWireType(field.Desc.Kind())
	fieldNumber := field.Desc.Number()
	if packed {
		wireType = protowire.BytesType
	}
	key := generator.KeySize(fieldNumber, wireType)
	switch field.Desc.Kind() {
	case protoreflect.DoubleKind, protoreflect.Fixed64Kind, protoreflect.Sfixed64Kind:
		if packed {
			p.P(`n+=`, strconv.Itoa(key), `+sov(uint64(len(m.`, fieldname, `)*8))`, `+len(m.`, fieldname, `)*8`)
		} else if repeated {
			p.P(`n+=`, strconv.Itoa(key+8), `*len(m.`, fieldname, `)`)
		} else if !oneof && !nullable {
			p.P(`if m.`, fieldname, ` != 0 {`)
			p.P(`n+=`, strconv.Itoa(key+8))
			p.P(`}`)
		} else {
			p.P(`n+=`, strconv.Itoa(key+8))
		}
	case protoreflect.FloatKind, protoreflect.Fixed32Kind, protoreflect.Sfixed32Kind:
		if packed {
			p.P(`n+=`, strconv.Itoa(key), `+sov(uint64(len(m.`, fieldname, `)*4))`, `+len(m.`, fieldname, `)*4`)
		} else if repeated {
			p.P(`n+=`, strconv.Itoa(key+4), `*len(m.`, fieldname, `)`)
		} else if !oneof && !nullable {
			p.P(`if m.`, fieldname, ` != 0 {`)
			p.P(`n+=`, strconv.Itoa(key+4))
			p.P(`}`)
		} else {
			p.P(`n+=`, strconv.Itoa(key+4))
		}
	case protoreflect.Int64Kind, protoreflect.Uint64Kind, protoreflect.Uint32Kind, protoreflect.EnumKind, protoreflect.Int32Kind:
		if packed {
			p.P(`l = 0`)
			p.P(`for _, e := range m.`, fieldname, ` {`)
			p.P(`l+=sov(uint64(e))`)
			p.P(`}`)
			p.P(`n+=`, strconv.Itoa(key), `+sov(uint64(l))+l`)
		} else if repeated {
			p.P(`for _, e := range m.`, fieldname, ` {`)
			p.P(`n+=`, strconv.Itoa(key), `+sov(uint64(e))`)
			p.P(`}`)
		} else if nullable {
			p.P(`n+=`, strconv.Itoa(key), `+sov(uint64(*m.`, fieldname, `))`)
		} else if oneof {
			p.P(`n+=`, strconv.Itoa(key), `+sov(uint64(m.value.AsInt64()))`)
		} else {
			p.P(`if m.`, fieldname, ` != 0 {`)
			p.P(`n+=`, strconv.Itoa(key), `+sov(uint64(m.`, fieldname, `))`)
			p.P(`}`)
		}
	case protoreflect.BoolKind:
		if packed {
			p.P(`n+=`, strconv.Itoa(key), `+sov(uint64(len(m.`, fieldname, `)))`, `+len(m.`, fieldname, `)*1`)
		} else if repeated {
			p.P(`n+=`, strconv.Itoa(key+1), `*len(m.`, fieldname, `)`)
		} else if !oneof && !nullable {
			p.P(`if m.`, fieldname, ` {`)
			p.P(`n+=`, strconv.Itoa(key+1))
			p.P(`}`)
		} else {
			p.P(`n+=`, strconv.Itoa(key+1))
		}
	case protoreflect.StringKind:
		if repeated {
			p.P(`for _, s := range m.`, fieldname, ` { `)
			p.P(`l = len(s)`)
			p.P(`n+=`, strconv.Itoa(key), `+l+sov(uint64(l))`)
			p.P(`}`)
		} else if nullable {
			p.P(`l=len(*m.`, fieldname, `)`)
			p.P(`n+=`, strconv.Itoa(key), `+l+sov(uint64(l))`)
		} else if oneof {
			p.P(`l=len(m.value.AsString())`)
			p.P(`n+=`, strconv.Itoa(key), `+l+sov(uint64(l))`)
		} else {
			p.P(`l=len(m.`, fieldname, `)`)
			p.P(`if l > 0 {`)
			p.P(`n+=`, strconv.Itoa(key), `+l+sov(uint64(l))`)
			p.P(`}`)
		}
	case protoreflect.GroupKind:
		p.messageSize("m."+fieldname, sizeName, field.Message)
		p.P(`n+=l+`, strconv.Itoa(2*key))
	case protoreflect.MessageKind:
		if field.Desc.IsMap() {
			fieldKeySize := generator.KeySize(field.Desc.Number(), generator.ProtoWireType(field.Desc.Kind()))
			keyKeySize := generator.KeySize(1, generator.ProtoWireType(field.Message.Fields[0].Desc.Kind()))
			valueKeySize := generator.KeySize(2, generator.ProtoWireType(field.Message.Fields[1].Desc.Kind()))
			p.P(`for k, v := range m.`, fieldname, ` { `)
			p.P(`_ = k`)
			p.P(`_ = v`)
			sum := []string{strconv.Itoa(keyKeySize)}

			switch field.Message.Fields[0].Desc.Kind() {
			case protoreflect.DoubleKind, protoreflect.Fixed64Kind, protoreflect.Sfixed64Kind:
				sum = append(sum, `8`)
			case protoreflect.FloatKind, protoreflect.Fixed32Kind, protoreflect.Sfixed32Kind:
				sum = append(sum, `4`)
			case protoreflect.Int64Kind, protoreflect.Uint64Kind, protoreflect.Uint32Kind, protoreflect.EnumKind, protoreflect.Int32Kind:
				sum = append(sum, `sov(uint64(k))`)
			case protoreflect.BoolKind:
				sum = append(sum, `1`)
			case protoreflect.StringKind, protoreflect.BytesKind:
				sum = append(sum, `len(k)+sov(uint64(len(k)))`)
			case protoreflect.Sint32Kind, protoreflect.Sint64Kind:
				sum = append(sum, `soz(uint64(k))`)
			}

			switch field.Message.Fields[1].Desc.Kind() {
			case protoreflect.DoubleKind, protoreflect.Fixed64Kind, protoreflect.Sfixed64Kind:
				sum = append(sum, strconv.Itoa(valueKeySize))
				sum = append(sum, strconv.Itoa(8))
			case protoreflect.FloatKind, protoreflect.Fixed32Kind, protoreflect.Sfixed32Kind:
				sum = append(sum, strconv.Itoa(valueKeySize))
				sum = append(sum, strconv.Itoa(4))
			case protoreflect.Int64Kind, protoreflect.Uint64Kind, protoreflect.Uint32Kind, protoreflect.EnumKind, protoreflect.Int32Kind:
				sum = append(sum, strconv.Itoa(valueKeySize))
				sum = append(sum, `sov(uint64(v))`)
			case protoreflect.BoolKind:
				sum = append(sum, strconv.Itoa(valueKeySize))
				sum = append(sum, `1`)
			case protoreflect.StringKind:
				sum = append(sum, strconv.Itoa(valueKeySize))
				sum = append(sum, `len(v)+sov(uint64(len(v)))`)
			case protoreflect.BytesKind:
				p.P(`l = `, strconv.Itoa(valueKeySize), ` + len(v)+sov(uint64(len(v)))`)
				sum = append(sum, `l`)
			case protoreflect.Sint32Kind, protoreflect.Sint64Kind:
				sum = append(sum, strconv.Itoa(valueKeySize))
				sum = append(sum, `soz(uint64(v))`)
			case protoreflect.MessageKind:
				p.P(`l = 0`)
				p.P(`if v != nil {`)
				p.messageSize("v", sizeName, field.Message.Fields[1].Message)
				p.P(`}`)
				p.P(`l += `, strconv.Itoa(valueKeySize), ` + sov(uint64(l))`)
				sum = append(sum, `l`)
			}
			p.P(`mapEntrySize := `, strings.Join(sum, "+"))
			p.P(`n+=mapEntrySize+`, fieldKeySize, `+sov(uint64(mapEntrySize))`)
			p.P(`}`)
		} else if field.Desc.IsList() {
			p.P(`for _, e := range m.`, fieldname, ` { `)
			p.messageSize("e", sizeName, field.Message)
			p.P(`n+=`, strconv.Itoa(key), `+l+sov(uint64(l))`)
			p.P(`}`)
		} else if oneof {
			p.messageSize(`(*`+p.QualifiedGoIdent(field.Message.GoIdent)+`)(m.value.AsPointer())`, sizeName, field.Message)
			p.P(`n+=`, strconv.Itoa(key), `+l+sov(uint64(l))`)
		} else {
			p.messageSize(`m.`+fieldname, sizeName, field.Message)
			p.P(`n+=`, strconv.Itoa(key), `+l+sov(uint64(l))`)
		}
	case protoreflect.BytesKind:
		if repeated {
			p.P(`for _, b := range m.`, fieldname, ` { `)
			p.P(`l = len(b)`)
			p.P(`n+=`, strconv.Itoa(key), `+l+sov(uint64(l))`)
			p.P(`}`)
		} else if oneof {
			p.P(`l=len(m.value.AsBytes())`)
			p.P(`n+=`, strconv.Itoa(key), `+l+sov(uint64(l))`)
		} else {
			p.P(`l=len(m.`, fieldname, `)`)
			p.P(`if l > 0 {`)
			p.P(`n+=`, strconv.Itoa(key), `+l+sov(uint64(l))`)
			p.P(`}`)
		}
	case protoreflect.Sint32Kind, protoreflect.Sint64Kind:
		if packed {
			p.P(`l = 0`)
			p.P(`for _, e := range m.`, fieldname, ` {`)
			p.P(`l+=soz(uint64(e))`)
			p.P(`}`)
			p.P(`n+=`, strconv.Itoa(key), `+sov(uint64(l))+l`)
		} else if repeated {
			p.P(`for _, e := range m.`, fieldname, ` {`)
			p.P(`n+=`, strconv.Itoa(key), `+soz(uint64(e))`)
			p.P(`}`)
		} else if nullable {
			p.P(`n+=`, strconv.Itoa(key), `+soz(uint64(*m.`, fieldname, `))`)
		} else if oneof {
			p.P(`n+=`, strconv.Itoa(key), `+soz(uint64(m.value.AsInt64()))`)
		} else {
			p.P(`if m.`, fieldname, ` != 0 {`)
			p.P(`n+=`, strconv.Itoa(key), `+soz(uint64(m.`, fieldname, `))`)
			p.P(`}`)
		}
	default:
		panic("not implemented")
	}
	if repeated || (nullable && !oneof) {
		p.P(`}`)
	}
}

func (p *size) message(message *protogen.Message) {
	for _, nested := range message.Messages {
		p.message(nested)
	}

	if message.Desc.IsMapEntry() {
		return
	}

	p.once = true

	ccTypeName := message.GoIdent

	p.P(`func (m *`, ccTypeName, `) SizeVT() (n int) {`)
	p.P(`if m == nil {`)
	p.P(`return 0`)
	p.P(`}`)
	p.P(`var l int`)
	p.P(`_ = l`)
	oneofs := make(map[string]struct{})
	for _, field := range message.Fields {
		if !p.IsOneof(field) {
			p.field(false, field, `SizeVT`)
		} else {
			fieldname := field.Oneof.GoName
			if _, ok := oneofs[fieldname]; !ok {
				oneofs[fieldname] = struct{}{}
				p.P(`n+=m.`, fieldname, `.SizeVT()`)
			}
		}
	}
	if p.HasUnknownFields() {
		p.P(`if m.unknownFields != nil {`)
		p.P(`n+=len(m.unknownFields)`)
		p.P(`}`)
	}
	p.P(`return n`)
	p.P(`}`)
	p.P()

	for _, oneof := range message.Oneofs {
		if oneof.Desc.IsSynthetic() {
			continue
		}
		p.P(`func (m *`, oneof.GoIdent, `) SizeVT() (n int) {`)
		p.P(`if m == nil {`)
		p.P(`return 0`)
		p.P(`}`)
		p.P(`var l int`)
		p.P(`_ = l`)

		p.P(`switch m.valueType {`)
		for i := range oneof.Fields {
			p.P(`case `, p.OneofFieldType(oneof, i), `:`)
			p.field(true, oneof.Fields[i], "SizeVT")
		}
		p.P(`}`)

		p.P(`return n`)
		p.P(`}`)
	}
}
