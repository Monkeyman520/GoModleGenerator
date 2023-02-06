package main

import (
	"errors"
	"fmt"
	"github.com/Monkeyman520/GoModleGenerator/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// typeConvertorMap map for converting mysql type to golang types
var typeConvertorMap = map[string]string{
	"int":                "int64",
	"integer":            "int64",
	"tinyint":            "int64",
	"smallint":           "int64",
	"mediumint":          "int64",
	"bigint":             "int64",
	"int unsigned":       "int64",
	"integer unsigned":   "int64",
	"tinyint unsigned":   "int64",
	"smallint unsigned":  "int64",
	"mediumint unsigned": "int64",
	"bigint unsigned":    "int64",
	"bit":                "int64",
	"bool":               "bool",
	"enum":               "string",
	"set":                "string",
	"varchar":            "string",
	"char":               "string",
	"tinytext":           "string",
	"mediumtext":         "string",
	"text":               "string",
	"longtext":           "string",
	"blob":               "string",
	"tinyblob":           "string",
	"mediumblob":         "string",
	"longblob":           "string",
	"date":               "time.Time", // time.Time or string
	"datetime":           "time.Time", // time.Time or string
	"timestamp":          "time.Time", // time.Time or string
	"time":               "time.Time", // time.Time or string
	"float":              "float64",
	"double":             "float64",
	"decimal":            "float64",
	"binary":             "string",
	"varbinary":          "string",
}

func (m *ModelGenerator) connectDB() {
	db, _ := gorm.Open(mysql.Open(m.dsn), &gorm.Config{})
	sql, _ := db.DB()
	err := sql.Ping()
	if err != nil {
		panic(err)
	}
	m.db = db
}

func (m *ModelGenerator) getAllTables() ([]string, error) {
	tables, err := m.db.Migrator().GetTables()
	if err != nil {
		return nil, err
	}
	return tables, nil
}

func (m *ModelGenerator) checkTables() ([]error, bool) {
	errorList := make([]error, 0)
	for _, table := range m.tables {
		if !m.db.Migrator().HasTable(table) {
			errorList = append(errorList, errors.New(table))
		}
	}
	return errorList, len(errorList) == 0
}

func (m *ModelGenerator) getRealNameFunc(structName, tableName string) string {
	return fmt.Sprintf("func (%s) TableName() string {\n  return \"%s\"\n\n}", structName, tableName)
}

type column struct {
	Name          string
	Type          string
	Nullable      string
	ColumnComment string
	Tag           string
}

func (c column) toStruct() string {
	return ""
}

type columns []column

func (c columns) toStruct() string {
	return ""
}

func (m *ModelGenerator) getColumInfo(table string) *columns {
	return nil
}

func (m *ModelGenerator) generatePackageHeader() string {
	return ""
}

type exportOption struct {
	writeFunction func(string, string) error
}

func (m *ModelGenerator) Export(option exportOption) {
	tables := m.tables
	if len(tables) != 0 {
		err, ok := m.checkTables()
		if !ok {
			panic(err)
		}
	} else {
		allTables, err := m.getAllTables()
		if err != nil {
			panic(err)
		}
		tables = allTables
	}

	for _, table := range tables {
		_ = m.getColumInfo(table)
	}
}

func main() {
	util.Parse()
}
