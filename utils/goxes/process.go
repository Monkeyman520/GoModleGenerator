package goxes

import (
	"github.com/Monkeyman520/GoModleGenerator/schema/information"
	"github.com/Monkeyman520/GoModleGenerator/utils/format"
)

func (gg *GoxGenerator) Process(tableName string, columns *information.Columns) error {
	gg.generate(tableName, columns)
	if gg.SeparateFile {
		err := gg.Write(
			gg.SavePath, format.Format("%s.go",
				format.StringRemovePrefix(
					gg.Prefix,
					format.ToLower(format.ToUnderscore(tableName)),
				),
			),
		)
		if err != nil {
			return err
		}
		gg.Flush()
		return nil
	}
	err := gg.Write(
		gg.SavePath, "model.go",
	)
	if err != nil {
		return err
	}
	return nil
}

func (gg *GoxGenerator) generate(tableName string, columns *information.Columns) {
	name := tableName
	if gg.StructNameToHump {
		name = format.BigCamel(name)
	} else {
		name = format.ToUpper(name)
	}
	fields := new(StructFields)
	for _, c := range *columns {
		columnName := c.ColumnName
		if gg.FieldToHump {
			columnName = format.BigCamel(columnName)
		} else {
			columnName = format.ToUpper(columnName)
		}
		*fields = append(*fields, StructField{
			FieldName: columnName,
			Type:      c.DataType,
		})
	}

	gg.NewStructure(name, fields, *gg.ConvertToTags(gg.TagKeys, fields))
	if gg.RealNameMethod {
		gg.NewRealNameMethod(name, tableName)
	}
}
