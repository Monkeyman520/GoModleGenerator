package format

import "fmt"

func Format(format string, content ...any) string {
	return fmt.Sprintf(format, content...)
}
