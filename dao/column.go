package dao

import (
	"github.com/Monkeyman520/GoModleGenerator/config"
	"github.com/Monkeyman520/GoModleGenerator/schema/information"
)

func GetColumn(table *information.Table) (*information.Columns, error) {
	db := config.GetDB()
	cfg := config.GetGenerator()
	columns := new(information.Columns)

	if err := db.Table("information_schema.columns").
		Where("table_schema = ?", cfg.Database.Schema).
		Where("table_name = ?", table.TableName).Find(columns).Error; err != nil {
		return nil, err
	}

	return columns, nil
}
