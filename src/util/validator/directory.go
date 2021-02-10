package validator

import (
	"os"
)

func IsDirExist(dir string) error {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// IsFileExist returns a boolean indicating whether the error is known to
// report that a file or directory does not exist.
func IsFileExist(uri string) bool {
	_, err := os.Stat(uri)
	return !os.IsNotExist(err)
}