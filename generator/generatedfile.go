// Copyright (c) 2021 PlanetScale Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package generator

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type GeneratedFile struct {
	*protogen.GeneratedFile
	Ext           *Extensions
	LocalPackages map[string]bool
}

func (p *GeneratedFile) Ident(path, ident string) string {
	return p.QualifiedGoIdent(protogen.GoImportPath(path).Ident(ident))
}

func (p *GeneratedFile) Alloc(vname string, message *protogen.Message) {
	p.P(vname, " := new(", message.GoIdent, `)`)
}

func (p *GeneratedFile) FieldGoType(field *protogen.Field) (goType string, pointer bool) {
	if field.Desc.IsWeak() {
		return "struct{}", false
	}

	pointer = field.Desc.HasPresence()
	switch field.Desc.Kind() {
	case protoreflect.BoolKind:
		goType = "bool"
	case protoreflect.EnumKind:
		goType = p.QualifiedGoIdent(field.Enum.GoIdent)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		goType = "int32"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		goType = "uint32"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		goType = "int64"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		goType = "uint64"
	case protoreflect.FloatKind:
		goType = "float32"
	case protoreflect.DoubleKind:
		goType = "float64"
	case protoreflect.StringKind:
		goType = "string"
	case protoreflect.BytesKind:
		goType = "[]byte"
		pointer = false // rely on nullability of slices for presence
	case protoreflect.MessageKind, protoreflect.GroupKind:
		goType = ""
		if p.IsNullable(field) {
			goType += "*"
		}
		goType += p.QualifiedGoIdent(field.Message.GoIdent)
		pointer = false // pointer captured as part of the type
	}
	switch {
	case field.Desc.IsList():
		return "[]" + goType, false
	case field.Desc.IsMap():
		keyType, _ := p.FieldGoType(field.Message.Fields[0])
		valType, _ := p.FieldGoType(field.Message.Fields[1])
		return fmt.Sprintf("map[%v]%v", keyType, valType), false
	}
	return goType, pointer
}

func (p *GeneratedFile) IsLocalMessage(message *protogen.Message) bool {
	// pkg := string(message.Desc.ParentFile().Package())
	// return p.LocalPackages[pkg]
	return true
}

func (p *GeneratedFile) HasUnknownFields() bool {
	// pkg := string(message.Desc.ParentFile().Package())
	// return p.LocalPackages[pkg]
	return false
}

func (p *GeneratedFile) OneofStruct(oneof *protogen.Oneof) string {
	return `optional_` + p.QualifiedGoIdent(oneof.GoIdent)
}

func (p *GeneratedFile) OneofType(oneof *protogen.Oneof) string {
	return p.QualifiedGoIdent(oneof.GoIdent) + `_Type`
}

func (p *GeneratedFile) OneofFieldNotPresent(oneof *protogen.Oneof) string {
	return p.OneofType(oneof) + `_notPresent`
}

func (p *GeneratedFile) OneofFieldType(oneof *protogen.Oneof, ix int) string {
	return p.OneofType(oneof) + `_` + oneof.Fields[ix].GoName
}

func (p *GeneratedFile) OneofFieldTypeFromField(field *protogen.Field) string {
	return p.OneofType(field.Oneof) + `_` + field.GoName
}

func (p *GeneratedFile) IsOneof(field *protogen.Field) bool {
	return field.Oneof != nil && !field.Oneof.Desc.IsSynthetic()
}

func (p *GeneratedFile) IsNullable(field *protogen.Field) bool {
	// return field.Message != nil || (field.Oneof != nil && field.Oneof.Desc.IsSynthetic())
	return field.Oneof != nil && field.Oneof.Desc.IsSynthetic()
}
