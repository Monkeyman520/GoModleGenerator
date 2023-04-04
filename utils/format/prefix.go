package format

import (
	"fmt"
	"strings"
)

func ListWithPrefix(prefix string, list []string) []string {
	res := make([]string, 0)
	if prefix == "" {
		return list
	}
	for _, l := range list {
		res = append(res, StringWithPrefix(prefix, l))
	}

	return res
}

func StringWithPrefix(prefix string, str string) string {
	return fmt.Sprintf("%s_%s", prefix, str)
}

func StringRemovePrefix(prefix string, str string) string {
	return strings.Replace(str, prefix, "", 1)
}

func StringListRemovePrefix(prefix string, list []string) []string {
	res := make([]string, 0)
	for _, l := range list {
		res = append(res, StringRemovePrefix(prefix, l))
	}
	return res
}
