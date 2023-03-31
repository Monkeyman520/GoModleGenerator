package config

// ConverterConfig defines the configuration for converting database tables and columns to go structures and fields
type ConverterConfig struct {
	// DateToTime Whether to convert the type date to time default true
	DateToTime bool `json:"DateToTime" yaml:"DateToTime"`

	// JsonTagToHump Whether to specify the json tag as a big hump default false
	JsonTagToHump bool `json:"JsonTagToHump" yaml:"JsonTagToHump"`

	// RmTagIfUcFirst If the first letter of the field is already in upper case, no tag is added default false
	RmTagIfUcFirst bool `json:"RmTagIfUcFirst" yaml:"RmTagIfUcFirst"`

	// StructNameToHump Whether the structure name is converted to humpback default false
	StructNameToHump bool `json:"StructNameToHump" yaml:"StructNameToHump"`

	// TagToLower Whether the field name of tag is converted to lower case default false
	TagToLower bool `json:"TagToLower" yaml:"TagToLower"`

	// UcFirstOnly Field first letter capitalised convert other letters to lower case default false
	UcFirstOnly bool `json:"UcFirstOnly" yaml:"UcFirstOnly"`
}

// NewConverterConfig Generate a ConverterConfig with the options passed in
func NewConverterConfig(options ...ConverterOption) *ConverterConfig {
	converterConfig := &ConverterConfig{
		DateToTime: true,
	}
	for _, option := range options {
		option(converterConfig)
	}
	return converterConfig
}

type ConverterOption func(*ConverterConfig)

// WithoutDateToTime Init GeneratorConfig without converting the type date to time
func WithoutDateToTime() ConverterOption {
	return func(cc *ConverterConfig) {
		cc.DateToTime = false
	}
}

// WithRmTagIfUcFirst Init GeneratorConfig with rm tag if uc first
func WithRmTagIfUcFirst() ConverterOption {
	return func(cc *ConverterConfig) {
		cc.RmTagIfUcFirst = true
	}
}

// WithStructNameToHump Init GeneratorConfig with struct name to hump
func WithStructNameToHump() ConverterOption {
	return func(cc *ConverterConfig) {
		cc.StructNameToHump = true
	}
}

// WithTagToLower Init GeneratorConfig with tag to lower
func WithTagToLower() ConverterOption {
	return func(cc *ConverterConfig) {
		cc.TagToLower = true
	}
}

// WithUcFirstOnly Init GeneratorConfig with UcFirstOnly
func WithUcFirstOnly() ConverterOption {
	return func(cc *ConverterConfig) {
		cc.UcFirstOnly = true
	}
}
