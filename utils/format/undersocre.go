package format

import "strings"

func ToUnderscore(str string) string {
	var underscoreString string
	for i, char := range str {
		if i > 0 && char >= 'A' && char <= 'Z' {
			underscoreString += "_"
		}
		underscoreString += strings.ToLower(string(char))
	}

	return underscoreString
}

func Uppercase(str string) string {
	words := strings.Split(str, "_")

	var camelCaseString []string
	for _, word := range words {
		camelCaseString = append(camelCaseString, ToCamelCase(word))
	}

	return strings.Join(camelCaseString, "_")
}
