package format

import (
	"strings"
)

func BigCamel(str string) string {
	words := strings.Split(str, "_")

	var camelCaseString string
	for _, word := range words {
		camelCaseString += ToCamelCase(word)
	}

	return camelCaseString
}

func SmallCamel(str string) string {
	words := strings.Split(str, "_")

	var camelCaseString string
	for i, word := range words {
		if i == 0 {
			camelCaseString += strings.ToLower(word)
		} else {
			camelCaseString += ToCamelCase(word)
		}
	}

	return camelCaseString
}

func ToCamelCase(str string) string {
	return strings.ToUpper(str[:1]) + str[1:]
}

func ToLower(str string) string {
	return strings.ToLower(str)
}

func ToUpper(str string) string {
	return strings.ToUpper(str)
}
