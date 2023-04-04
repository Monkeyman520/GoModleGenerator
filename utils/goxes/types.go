package goxes

import (
	"github.com/Monkeyman520/GoModleGenerator/config"
	"github.com/Monkeyman520/GoModleGenerator/utils/save"
	"github.com/goplus/gox"
)

type GoxGenerator struct {
	config.GeneratorConfig
	config.ConverterConfig
	pkg *gox.Package
}

func (gg *GoxGenerator) PreCheck() error {
	return save.CheckPath(gg.SavePath)
}

func NewGoxGenerator(config *config.ModelGenerator) (*GoxGenerator, error) {
	generator := &GoxGenerator{
		GeneratorConfig: *config.Generator,
		ConverterConfig: *config.Converter,
		pkg:             gox.NewPackage(config.Generator.PackagePath, config.Generator.PackageName, nil),
	}
	return generator, generator.PreCheck()
}

// TypeConvertorMap map for converting mysql type to golang types
var TypeConvertorMap = map[string]int{
	"int":                6,
	"integer":            6,
	"tinyint":            6,
	"smallint":           6,
	"mediumint":          6,
	"bigint":             6,
	"int unsigned":       6,
	"integer unsigned":   6,
	"tinyint unsigned":   6,
	"smallint unsigned":  6,
	"mediumint unsigned": 6,
	"bigint unsigned":    6,
	"bit":                6,
	"bool":               1,
	"enum":               17,
	"set":                17,
	"varchar":            17,
	"char":               17,
	"tinytext":           17,
	"mediumtext":         17,
	"text":               17,
	"longtext":           17,
	"blob":               17,
	"tinyblob":           17,
	"mediumblob":         17,
	"longblob":           17,
	"date":               42, // time.Time or string
	"datetime":           42, // time.Time or string
	"timestamp":          42, // time.Time or string
	"time":               42, // time.Time or string
	"float":              14,
	"double":             14,
	"decimal":            14,
	"binary":             17,
	"varbinary":          17,
}
