package config

// ConverterConfig defines the configuration for converting database tables and columns to go structures and fields
type ConverterConfig struct {
	// DateToTime Whether to convert the type date to time default true
	DateToTime bool `json:"DateToTime" yaml:"DateToTime"`

	// FieldToHump Whether to specify the field as a big hump default true
	FieldToHump bool `json:"FieldToHump" yaml:"FieldToHump"`

	// StructNameToHump Whether the structure name is converted to humpback default ture
	StructNameToHump bool `json:"StructNameToHump" yaml:"StructNameToHump"`

	// TagToHump Whether to specify the tag as a big hump default true
	TagToHump bool `json:"TagToHump" yaml:"TagToHump"`

	// TagToLower Whether the field name of tag is converted to lower case default false
	TagToLower bool `json:"TagToLower" yaml:"TagToLower"`
}

// NewConverterConfig Generate a ConverterConfig with the options passed in
func NewConverterConfig(options ...ConverterOption) *ConverterConfig {
	converterConfig := &ConverterConfig{
		DateToTime:       true,
		StructNameToHump: true,
		TagToHump:        true,
		FieldToHump:      true,
	}
	for _, option := range options {
		option(converterConfig)
	}
	return converterConfig
}

type ConverterOption func(*ConverterConfig)

// WithoutDateToTime Init ConverterConfig without converting the type date to time
func WithoutDateToTime() ConverterOption {
	return func(cc *ConverterConfig) {
		cc.DateToTime = false
	}
}

// WithoutFieldToHump Init ConverterConfig without converting the type field to hump
func WithoutFieldToHump() ConverterOption {
	return func(cc *ConverterConfig) {
		cc.FieldToHump = false
	}
}

// WithoutStructNameToHump Init ConverterConfig without struct name to hump
func WithoutStructNameToHump() ConverterOption {
	return func(cc *ConverterConfig) {
		cc.StructNameToHump = false
	}
}

// WithoutTagToHump Init ConverterConfig without tag to hump
func WithoutTagToHump() ConverterOption {
	return func(cc *ConverterConfig) {
		cc.TagToHump = false
	}
}

// WithTagToLower Init ConverterConfig with tag to lower
func WithTagToLower() ConverterOption {
	return func(cc *ConverterConfig) {
		cc.TagToLower = true
	}
}
