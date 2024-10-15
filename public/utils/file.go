package utils

import (
	"errors"
	"os"
)

func GetFilePath(paths ...string) (string, error) {
	var err error
	for _, path := range paths {
		_, _err := os.Stat(path)
		if _err == nil {
			return path, nil
		} else {
			err = errors.Join(err, _err)

		}
	}
	return "", err
}
