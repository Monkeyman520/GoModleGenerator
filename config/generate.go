package config

// GeneratorConfig A generator config to modify the out file.
type GeneratorConfig struct {
	// GormModel Whether to add gorm.Model default false
	GormModel bool `json:"GormModel" yaml:"GormModel"`

	// JsonTag Whether to add json tag default false
	JsonTag bool `json:"JsonTag" yaml:"JsonTag"`

	// PackagePath Whether to specify the package path default ""
	PackagePath string `json:"PackagePath" yaml:"PackagePath"`

	// PackageName Whether to specify a package name default "model"
	PackageName string `json:"PackageName" yaml:"PackageName"`

	// Prefix Whether to specify a field prefix default ""
	Prefix string `json:"Prefix" yaml:"Prefix"`

	// RealNameMethod Whether to implement the table name interface default false
	RealNameMethod bool `json:"RealNameMethod" yaml:"RealNameMethod"`

	// SavePath Whether to specify the path to save the generated file default "./model"
	SavePath string `json:"SavePath" yaml:"SavePath"`

	// SeparateFile Whether each exported structure needs to be placed in a separate file default true
	SeparateFile bool `json:"SeparateFile" yaml:"SeparateFile"`

	// Tables Specify the name of the table to be exported from the database default []string{} which means all
	Tables []string `json:"Tables" yaml:"Tables"`

	// TagKeys Specify the name of the tag to be added default "gorm"
	TagKeys []string `json:"TagKeys" yaml:"TagKeys"`
}

// NewGeneratorConfig Generate a GeneratorConfig with the options passed in
func NewGeneratorConfig(options ...GeneratorOption) *GeneratorConfig {
	generatorConfig := &GeneratorConfig{
		PackageName:  "model",
		SavePath:     "./model/",
		SeparateFile: true,
		TagKeys:      []string{"gorm"},
	}
	for _, option := range options {
		option(generatorConfig)
	}
	return generatorConfig
}

type GeneratorOption func(*GeneratorConfig)

// WithGormModel Init GeneratorConfig with gorm.Model
func WithGormModel() GeneratorOption {
	return func(gc *GeneratorConfig) {
		gc.GormModel = true
	}
}

// WithJsonTag Init GeneratorConfig with json tag
func WithJsonTag() GeneratorOption {
	return func(gc *GeneratorConfig) {
		gc.JsonTag = true
		gc.TagKeys = append(gc.TagKeys, "json")
	}
}

// WithRealNameMethod Init GeneratorConfig with realName method
func WithRealNameMethod() GeneratorOption {
	return func(gc *GeneratorConfig) {
		gc.RealNameMethod = true
	}
}

// WithPackageName Init GeneratorConfig with the specific package name
// The default value is model.
func WithPackageName(packageName string) GeneratorOption {
	return func(gc *GeneratorConfig) {
		gc.PackageName = packageName
	}
}

// WithPackagePath Init GeneratorConfig with the specific package path
func WithPackagePath(packagePath string) GeneratorOption {
	return func(gc *GeneratorConfig) {
		gc.PackagePath = packagePath
	}
}

// WithPrefix Init GeneratorConfig with the specific prefix
func WithPrefix(prefix string) GeneratorOption {
	return func(gc *GeneratorConfig) {
		gc.Prefix = prefix
	}
}

// WithSavePath Init GeneratorConfig with the specific save path
// The default value is "./"
func WithSavePath(savePath string) GeneratorOption {
	return func(gc *GeneratorConfig) {
		gc.SavePath = savePath
	}
}

// WithoutSeparateFile Init GeneratorConfig without each struct in a separate file
func WithoutSeparateFile() GeneratorOption {
	return func(gc *GeneratorConfig) {
		gc.SeparateFile = false
	}
}

// WithTables Init GeneratorConfig with the specific tables
func WithTables(tables []string) GeneratorOption {
	return func(gc *GeneratorConfig) {
		gc.Tables = append(gc.Tables, tables...)
	}
}

// WithTagKeys Init GeneratorConfig with other tags
// such as "orm" "yaml"
func WithTagKeys(tagKeys []string) GeneratorOption {
	return func(gc *GeneratorConfig) {
		gc.TagKeys = append(gc.TagKeys, tagKeys...)
	}
}
