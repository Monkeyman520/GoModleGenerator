package goxes

import (
	"fmt"
	"github.com/Monkeyman520/GoModleGenerator/utils/format"
	"github.com/goplus/gox"
	"go/token"
	"go/types"
	"strings"
)

func (gg *GoxGenerator) CtxRef(name string) gox.Ref {
	return gg.pkg.CB().Scope().Lookup(name)
}

func (gg *GoxGenerator) GetPkg() *gox.Package {
	return gg.pkg
}

func (gg *GoxGenerator) NewPackage(path, name string) *gox.Package {
	return gox.NewPackage(path, name, nil)
}

type GeneratedStructure struct {
	Name   string
	Fields StructFields
	Tags   []string
}

func (gg *GoxGenerator) NewStructure(name string, fields *StructFields, tags []string) {
	gg.pkg.NewType(name).InitType(gg.pkg, types.NewStruct(gg.ConvertToFields(fields), tags))
}

func (gg *GoxGenerator) NewRealNameMethod(typeName, realName string) {
	gg.pkg.NewFunc(
		types.NewParam(token.NoPos, gg.pkg.Types, "", gg.CtxRef(typeName).Type()),
		"TableName",
		nil,
		types.NewTuple(types.NewVar(token.NoPos, gg.pkg.Types, "", types.Typ[types.String])),
		false,
	).BodyStart(gg.pkg).
		Val(realName).Return(1).EndStmt().
		End()
}

func (gg *GoxGenerator) Flush() {
	gg.pkg = gox.NewPackage(gg.PackagePath, gg.PackageName, nil)
}

func (gg *GoxGenerator) Write(path string, name string) error {
	err := gg.pkg.WriteFile(fmt.Sprintf("%s/%s", path, name))
	if err != nil {
		return err
	}
	return nil
}

type StructField struct {
	FieldName string
	Type      string
}

type StructFields []StructField

func (gg *GoxGenerator) ConvertToFields(fields *StructFields) []*types.Var {
	res := make([]*types.Var, 0)
	if gg.GormModel {
		gormImport := gg.GetPkg().Import("gorm.io/gorm")
		res = append(res, types.NewField(token.NoPos, gg.pkg.Types, "", gormImport.Ref("Model").Type(), true))
	}
	for _, v := range *fields {
		fieldType := TypeConvertorMap[v.Type]
		field := new(types.Var)
		if fieldType <= 25 {
			field = types.NewField(token.NoPos, gg.pkg.Types, v.FieldName, types.Typ[fieldType], false)
		} else if gg.DateToTime {
			timeImport := gg.GetPkg().Import("time")
			field = types.NewField(token.NoPos, gg.pkg.Types, v.FieldName, timeImport.Ref("Time").Type(), false)
		} else {
			field = types.NewField(token.NoPos, gg.pkg.Types, v.FieldName, types.Typ[fieldType%25], false)
		}
		res = append(res, field)
	}
	return res
}

type (
	TagKeys = []string
	Tags    = []string
	Tag     = string
)

func convertKeysToTag(tagKeys TagKeys, content string) Tag {
	tags := new(Tags)
	for _, k := range tagKeys {
		*tags = append(*tags, fmt.Sprintf("%s:\"%s\"", k, content))
	}

	return strings.Join(*tags, " ")
}

func (gg *GoxGenerator) ConvertToTags(
	tagKeys TagKeys, fields *StructFields,
) *Tags {
	tags := new(Tags)

	if gg.GormModel {
		*tags = append(*tags, "")
	}

	for _, field := range *fields {
		tagContent := field.FieldName
		if gg.TagToHump {
			tagContent = format.BigCamel(tagContent)
		} else {
			tagContent = format.ToUnderscore(format.BigCamel(tagContent))
		}
		if gg.TagToLower {
			tagContent = format.ToLower(tagContent)
		}
		*tags = append(*tags, convertKeysToTag(tagKeys, tagContent))
	}
	return tags
}
