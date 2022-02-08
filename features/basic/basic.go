// Copyright (c) 2021 PlanetScale Inc. All rights reserved.
// Copyright (c) 2013, The GoGo Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package basic

import (
	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/bogdandrutu/lazyproto/generator"
)

func init() {
	generator.RegisterFeature("basic", func(gen *generator.GeneratedFile) generator.FeatureGenerator {
		return &basic{GeneratedFile: gen}
	})
}

type basic struct {
	*generator.GeneratedFile
}

var _ generator.FeatureGenerator = (*basic)(nil)

func (p *basic) Name() string {
	return "basic"
}

func (p *basic) GenerateFile(file *protogen.File) bool {
	for _, enum := range file.Enums {
		p.generateEnum(enum)
	}

	for _, message := range file.Messages {
		p.generateMessage(message)
	}

	return true
}

func (p *basic) GenerateHelpers() {}

func (p *basic) generateEnum(enum *protogen.Enum) {
	p.P(enum.Comments.Leading, `type `, enum.GoIdent, ` int32`)
	p.P()
	p.P(`const (`)
	for _, value := range enum.Values {
		value.GoIdent.GoName = toEnumValueName(string(value.Desc.Name()))
		p.P(value.Comments.Leading, value.GoIdent, ` `, enum.GoIdent, ` = `, value.Desc.Number())
	}
	p.P(`)`)
	p.P()
}

func (p *basic) generateMessage(message *protogen.Message) {
	if message.Desc.IsMapEntry() {
		return
	}

	p.P(message.Comments.Leading, `type `, message.GoIdent, ` struct {`)
	if p.HasUnknownFields() {
		p.P(`unknownFields []byte`)
	}

	// Ensures one field per one oneof entry in the message.
	oneofs := make(map[string]struct{})
	for _, field := range message.Fields {
		field.GoName = strcase.ToLowerCamel(field.GoName)

		oneof := field.Oneof
		if oneof == nil || oneof.Desc.IsSynthetic() {
			p.generateField(field)
			continue
		}
		oneof.GoName = strcase.ToLowerCamel(oneof.GoName)
		if _, ok := oneofs[oneof.GoName]; !ok {
			oneofs[oneof.GoName] = struct{}{}
			oneof.GoIdent.GoName = `optional_` + oneof.GoIdent.GoName
			p.P(oneof.Comments.Leading, oneof.GoName, ` `, oneof.GoIdent)
		}
	}
	p.P(`}`)
	p.P()

	for _, oneof := range message.Oneofs {
		if oneof.Desc.IsSynthetic() {
			continue
		}
		p.P(`type `, oneof.GoIdent, ` struct {`)
		p.P(`value `, p.Ident("github.com/bogdandrutu/lazyproto/variant", "Value"))
		p.P(`valueType `, oneof.GoIdent, `_Type`)
		p.P(`}`)
		p.P()
		p.P(`type `, p.OneofType(oneof), ` int32`)
		p.P(`const (`)
		p.P(p.OneofFieldNotPresent(oneof), ` `, p.OneofType(oneof), ` = iota`)
		for i := range oneof.Fields {
			p.P(p.OneofFieldType(oneof, i))
		}
		p.P(`)`)
		p.P()
	}

	for _, nested := range message.Messages {
		p.generateMessage(nested)
	}
}

func (p *basic) generateField(field *protogen.Field) {
	fieldType, optional := p.FieldGoType(field)
	if optional {
		p.P(field.GoName, ` *`, fieldType)
	} else {
		p.P(field.GoName, ` `, fieldType)
	}
}
