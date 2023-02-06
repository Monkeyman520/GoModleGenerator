package main

import "gorm.io/gorm"

// ModelGenerator A generator tool be used to convert tables to go structs.
// TODO: Add fields comment.
type ModelGenerator struct {
	dsn            string
	savePath       string
	db             *gorm.DB
	tables         []string
	prefix         string
	config         *GeneratorConfig
	err            []error
	realNameMethod bool
	enableJsonTag  bool   // 是否添加json的tag, 默认不添加
	packageName    string // 生成struct的包名(默认为空的话, 则取名为: package model)
	tagKey         string // tag字段的key值,默认是orm
	dateToTime     bool   // 是否将 date相关字段转换为 time.Time,默认否

}

// NewModelGenerator New a ModelGenerator with args.
func NewModelGenerator(options ...option) *ModelGenerator {
	// TODO: Add default value.
	m := new(ModelGenerator)
	for _, opt := range options {
		opt(m)
	}
	return m
}

// option NewModelGenerator's initialization options.
type option func(*ModelGenerator)

// WithDSN Init ModelGenerator with DSN
func WithDSN(dsn string) option {
	return func(m *ModelGenerator) {
		m.dsn = dsn
	}
}

// WithSavePath Init ModelGenerator with savePath.
func WithSavePath(path string) option {
	return func(m *ModelGenerator) {
		m.savePath = path
	}
}

// WithTables Init ModelGenerator with selected tables. The default value is all tables.
func WithTables(tables ...string) option {
	return func(m *ModelGenerator) {
		m.tables = append(m.tables, tables...)
	}
}

// WithPrefix Init ModelGenerator with prefix.
func WithPrefix(prefix string) option {
	return func(m *ModelGenerator) {
		m.prefix = prefix
	}
}

// WithGeneratorConfig Init ModelGenerator with config.
func WithGeneratorConfig(config *GeneratorConfig) option {
	return func(m *ModelGenerator) {
		m.config = config
	}
}

// WithRealNameMethod Init ModelGenerator with realNameMethod.
func WithRealNameMethod() option {
	return func(m *ModelGenerator) {
		m.realNameMethod = true
	}
}

// WithJsonTag Init ModelGenerator with jsonTag.
func WithJsonTag(tag string) option {
	return func(m *ModelGenerator) {
		m.enableJsonTag = true
	}
}

// WithPackageName Init ModelGenerator with packageName. The default value is model.
func WithPackageName(name string) option {
	return func(m *ModelGenerator) {
		m.packageName = name
	}
}

// WithTagKey Init ModelGenerator with tagKey. The default value is gorm.
func WithTagKey(tagKey string) option {
	return func(m *ModelGenerator) {
		m.tagKey = tagKey
	}
}

// WithDateToTime Init ModelGenerator with converting date related fields to time.Time.
func WithDateToTime() option {
	return func(m *ModelGenerator) {
		m.dateToTime = true
	}
}

// GeneratorConfig A generator config to modify the out file.
// TODO: Add fields comment.
type GeneratorConfig struct {
	structNameToHump bool // 结构体名称是否转为驼峰式，默认为false
	rmTagIfUcFirst   bool // 如果字段首字母本来就是大写, 就不添加tag, 默认false添加, true不添加
	tagToLower       bool // tag的字段名字是否转换为小写, 如果本身有大写字母的话, 默认false不转
	jsonTagToHump    bool // json tag是否转为驼峰，默认为false，不转换
	ucFirstOnly      bool // 字段首字母大写的同时, 是否要把其他字母转换为小写,默认false不转换
	separateFile     bool // 每个struct放入单独的文件,默认false,放入同一个文件
}

// NewGeneratorConfig New a generator config with args.
func NewGeneratorConfig(generatorOptions ...generatorOption) *GeneratorConfig {
	// TODO: Add default value.
	gc := new(GeneratorConfig)
	for _, opt := range generatorOptions {
		opt(gc)
	}
	return gc
}

// generatorOption GeneratorConfig's initialization options.
type generatorOption func(*GeneratorConfig)

// WithStructNameToHump Init GeneratorConfig with convert The first letter of the field is capitalized without adding the tag.
func WithStructNameToHump() generatorOption {
	return func(gc *GeneratorConfig) {
		gc.jsonTagToHump = true
	}
}

// WithRmTagIfUcFirst Init GeneratorConfig with convert structure name to hump.
func WithRmTagIfUcFirst() generatorOption {
	return func(gc *GeneratorConfig) {
		gc.rmTagIfUcFirst = true
	}
}

// WithTagToLower Init GeneratorConfig with convert tag field name to lower case.
func WithTagToLower() generatorOption {
	return func(gc *GeneratorConfig) {
		gc.tagToLower = true
	}
}

// WithKJsonTagToHump Init GeneratorConfig with convert json tag to hump.
func WithKJsonTagToHump() generatorOption {
	return func(gc *GeneratorConfig) {
		gc.jsonTagToHump = true
	}
}

// WithUcFirstOnly Init GeneratorConfig with convert structure fields the first letter of the field is capitalized and the rest of the letters are lowercase.
func WithUcFirstOnly() generatorOption {
	return func(gc *GeneratorConfig) {
		gc.ucFirstOnly = true
	}
}

// WithSeparateFile Init GeneratorConfig with each struct in a separate file.
func WithSeparateFile() generatorOption {
	return func(gc *GeneratorConfig) {
		gc.separateFile = true
	}
}
