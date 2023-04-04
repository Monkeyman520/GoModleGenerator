package information

type Column struct {
	TableCatalog           string `gorm:"column:TABLE_CATALOG;type:varchar(64)"`
	TableSchema            string `gorm:"column:TABLE_SCHEMA;type:varchar(64)"`
	TableName              string `gorm:"column:TABLE_NAME;type:varchar(64)"`
	ColumnName             string `gorm:"column:COLUMN_NAME;type:varchar(64)"`
	OrdinalPosition        int    `gorm:"column:ORDINAL_POSITION;type:int"`
	ColumnDefault          string `gorm:"column:COLUMN_DEFAULT;type:text"`
	IsNullAble             string `gorm:"column:IS_NULLABLE;type:varchar(3)"`
	DataType               string `gorm:"column:DATA_TYPE;type:longtext"`
	CharacterMaximumLength int    `gorm:"column:CHARACTER_MAXIMUM_LENGTH;type:bigint"`
	CharacterOctetLength   int    `gorm:"column:CHARACTER_OCTET_LENGTH;type:bigint"`
	NumericPrecision       int    `gorm:"column:NUMERIC_PRECISION;type:bigint"`
	NumericScale           int    `gorm:"column:NUMERIC_SCALE;type:bigint"`
	DateTimePrecision      int    `gorm:"column:DATETIME_PRECISION;type:int"`
	CharacterSetName       string `gorm:"column:CHARACTER_SET_NAME;type:varchar(64)"`
	CollationName          string `gorm:"column:COLLATION_NAME;type:varchar(64)"`
	ColumnType             string `gorm:"column:COLUMN_TYPE;type:mediumtext"`
	ColumnKey              string `gorm:"column:COLUMN_KEY;type:enum"`
	Privileges             string `gorm:"column:PRIVILEGES;type:varchar(154)"`
	Extra                  string `gorm:"column:EXTRA;type:varchar(256)"`
	ColumnComment          string `gorm:"COLUMN_COMMENT;type:text"`
	GenerationExpression   string `gorm:"column:GENERATION_EXPRESSION;type:longtext"`
	SrsID                  string `gorm:"column:SRS_ID;type:int"`
}

type Columns []Column
