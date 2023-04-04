package information

import "time"

type Table struct {
	TableCatalog   string    `gorm:"column:TABLE_CATALOG;type:varchar(64)"`
	TableSchema    string    `gorm:"column:TABLE_SCHEMA;type:varchar(64)"`
	TableName      string    `gorm:"column:TABLE_NAME;type:varchar(64)"`
	TableType      string    `gorm:"column:TABLE_TYPE;type:enum"`
	Engine         string    `gorm:"column:ENGINE;type:varchar(64)"`
	Version        int       `gorm:"column:VERSION;type:int"`
	RowFormat      string    `gorm:"column:ROW_FORMAT;type:enum"`
	TableRows      int       `gorm:"column:ROW_ROWS;type:bigint"`
	AvgRowLength   int       `gorm:"column:AVG_ROW_LENGTH;type:bigint"`
	DataLength     int       `gorm:"column:DATA_LENGTH;type:bigint"`
	MaxDataLength  int       `gorm:"column:MAX_DATA_LENGTH;type:bigint"`
	IndexLength    int       `gorm:"column:INDEX_LENGTH;type:bigint"`
	DataFree       int       `gorm:"column:DATA_FREE;type:bigint"`
	AutoIncrement  int       `gorm:"column:AUTO_INCREMENT;type:bigint"`
	CreateTime     time.Time `gorm:"column:CREATE_TIME;type:timestamp"`
	UpdateTime     time.Time `gorm:"column:UPDATE_TIME;type:datetime"`
	CheckTime      time.Time `gorm:"column:CHECK_TIME;type:datetime"`
	TableCollation string    `gorm:"column:TABLE_COLLATION;type:varchar(64)"`
	CheckSum       int       `gorm:"column:CHECK_SUM;type:bigint"`
	CreateOptions  string    `gorm:"column:CREATE_OPTIONS;type:varchar(25)"`
	TableComment   string    `gorm:"column:TABLE_COMMENT;type:text"`
}

type Tables []Table
