package config

import (
	"encoding/json"
	internalError "github.com/Monkeyman520/GoModleGenerator/utils/error"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"os"
)

// InitConfigWithFile use viper to read config file
func InitConfigWithFile() (*ModelGenerator, error) {
	config := viper.New()
	config.SetConfigName("config")
	config.AddConfigPath(".")

	if err := config.ReadInConfig(); err != nil {
		return nil, internalError.ViperReadError
	}

	var modelGenerator ModelGenerator
	err := config.Unmarshal(&modelGenerator)
	if err != nil {
		return nil, err
	}

	return &modelGenerator, nil
}

// GenerateDefaultConfigFile generate a default config file
func GenerateDefaultConfigFile() error {
	generator := NewModelGenerator(
		NewDatabaseConfig(
			WithHost("localhost"),
			WithPort(3306),
			WithUsername("username"),
			WithPassword("password"),
			WithSchema("test"),
		),
		NewConverterConfig(),
		NewGeneratorConfig(
			WithTables([]string{"user", "role"}),
			WithTagKeys([]string{"yaml", "gorm"}),
		),
	)
	ans, _ := json.Marshal(generator)
	err := os.WriteFile("./template/config.json", ans, 0755)
	if err != nil {
		return err
	}

	yml, _ := yaml.Marshal(generator)
	err = os.WriteFile("./template/config.yml", yml, 0755)
	if err != nil {
		return err
	}
	return nil
}
