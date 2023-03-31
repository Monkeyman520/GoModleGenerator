package config

import (
	internalError "github.com/Monkeyman520/GoModleGenerator/utils/error"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

// ModelGenerator A generator tool be used to convert tables to go structs.
type ModelGenerator struct {
	Database  *DatabaseConfig  `json:"Database" yaml:"Database"`
	Converter *ConverterConfig `json:"Converter" yaml:"Converter"`
	Generator *GeneratorConfig `json:"Generator" yaml:"Generator"`
}

var _ = (*ModelGenerator)(nil)

var (
	_dbOnce        sync.Once
	_dbLock                 = new(sync.RWMutex)
	DB             *gorm.DB = nil
	_ModelOnce     sync.Once
	_generatorLock                 = new(sync.RWMutex)
	Generator      *ModelGenerator = nil
)

// NewModelGenerator creates a new ModelGenerator with the given configurations
func NewModelGenerator(databaseConfig *DatabaseConfig, converterConfig *ConverterConfig, generatorConfig *GeneratorConfig) *ModelGenerator {
	_ModelOnce.Do(func() {
		generator := &ModelGenerator{databaseConfig, converterConfig, generatorConfig}
		err := generator.PreCheck()
		if err != nil {
			panic(err)
		}
		SetGenerator(generator)
	})
	return GetGenerator()
}

// SetGenerator sets the generator
func SetGenerator(generator *ModelGenerator) {
	_generatorLock.Lock()
	Generator = generator
	_generatorLock.Unlock()
}

// GetGenerator returns the preloaded generator
func GetGenerator() *ModelGenerator {
	_generatorLock.RLock()
	generator := Generator
	_generatorLock.RUnlock()
	return generator
}

// SetDB sets the database connection
func SetDB(db *gorm.DB) {
	_dbLock.Lock()
	DB = db
	_dbLock.Unlock()
}

// GetDB returns the preloaded database connection
func GetDB() *gorm.DB {
	_dbLock.RLock()
	db := DB
	_dbLock.RUnlock()
	return db
}

// PreCheck to ensure that the database source name is correct and accessible
func (mg *ModelGenerator) PreCheck() error {
	err := mg.Database.PreCheck()
	if err != nil {
		return err
	}

	db, err := gorm.Open(mysql.Open(mg.Database.DSN()), &gorm.Config{})
	if err != nil {
		return internalError.DbConnectionError
	}
	sql, err := db.DB()
	if err != nil {
		return internalError.DbConnectionError
	}
	err = sql.Ping()
	if err != nil {
		return internalError.DbConnectionError
	}

	_dbOnce.Do(func() {
		SetDB(db)
	})
	return nil
}
