package config

import (
	"fmt"
	internalError "github.com/Monkeyman520/GoModleGenerator/utils/error"
	"github.com/Monkeyman520/GoModleGenerator/utils/network"
)

// DatabaseConfig defines the configuration for connecting to the database
type DatabaseConfig struct {
	// Host the database host address
	Host string `json:"Host" yaml:"Host"`

	// Port the database port number default 3306
	Port int `json:"Port" yaml:"Port"`

	// User the database user default root
	Username string `json:"Username" yaml:"Username"`

	// Password the database password
	Password string `json:"Password" yaml:"Password"`

	// Schema the database schema
	Schema string `json:"Schema" yaml:"Schema"`
}

// DSN return the database source name to connect to the database
func (dc *DatabaseConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dc.Username, dc.Password, dc.Host, dc.Port, dc.Schema,
	)
}

// PreCheck check the database configuration info is valid
func (dc *DatabaseConfig) PreCheck() error {
	if valid := network.CheckHost(dc.Host); dc.Host == "" || !valid {
		return internalError.InvalidDatabaseHost
	}

	if dc.Port <= 80 || dc.Port >= 65530 {
		return internalError.InvalidDatabasePort
	}

	if dc.Username == "" {
		return internalError.InvalidDatabaseUsername
	}

	if dc.Password == "" {
		return internalError.InvalidDatabasePassword
	}

	if dc.Schema == "" {
		return internalError.InvalidDatabaseSchema
	}
	return nil
}

// NewDatabaseConfig Generate a NewDatabaseConfig with the options passed in
func NewDatabaseConfig(options ...DatabaseOption) *DatabaseConfig {
	databaseConfig := &DatabaseConfig{
		Port:     3306,
		Username: "root",
	}
	for _, option := range options {
		option(databaseConfig)
	}
	return databaseConfig
}

type DatabaseOption func(*DatabaseConfig)

// WithHost Init DatabaseConfig with the specific host
// It can be IP addr or Domain
func WithHost(host string) DatabaseOption {
	return func(dc *DatabaseConfig) {
		dc.Host = host
	}
}

// WithPort Init DatabaseConfig with the specific host
// The default value is 3306
func WithPort(port int) DatabaseOption {
	return func(dc *DatabaseConfig) {
		dc.Port = port
	}
}

// WithUsername Init DatabaseConfig with the specific username
// The default value is root
func WithUsername(username string) DatabaseOption {
	return func(dc *DatabaseConfig) {
		dc.Username = username
	}
}

// WithPassword Init DatabaseConfig with the specific password
func WithPassword(password string) DatabaseOption {
	return func(dc *DatabaseConfig) {
		dc.Password = password
	}
}

// WithSchema Init DatabaseConfig with the specific schema
func WithSchema(schema string) DatabaseOption {
	return func(dc *DatabaseConfig) {
		dc.Schema = schema
	}
}
