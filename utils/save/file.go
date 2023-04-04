package save

import (
	internalError "github.com/Monkeyman520/GoModleGenerator/utils/error"
	"os"
)

func CheckPath(path string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return internalError.MakePathError
			}
		}
		return internalError.UnknowPathError
	}
	return nil
}
