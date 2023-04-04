package dao

import (
	"github.com/Monkeyman520/GoModleGenerator/config"
	"github.com/Monkeyman520/GoModleGenerator/schema/information"
	internalError "github.com/Monkeyman520/GoModleGenerator/utils/error"
	"github.com/Monkeyman520/GoModleGenerator/utils/format"
)

// GetTables get the specific part of the tables in current database
func GetTables() (*information.Tables, error) {
	db := config.GetDB()
	cfg := config.GetGenerator()
	tables := new(information.Tables)

	db = db.Table("information_schema.tables").Where("TABLE_SCHEMA = ?", cfg.Database.Schema)

	if len(cfg.Generator.Tables) > 0 {
		db = db.Where("TABLE_NAME in ?", format.ListWithPrefix(cfg.Generator.Prefix, cfg.Generator.Tables))
	}

	if err := db.Find(tables).Error; err != nil {
		return nil, internalError.TablesSearchError
	}
	return tables, nil
}
