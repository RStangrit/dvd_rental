package file

import (
	"errors"
	"log"
	"os"
	"strings"
)

func isValidFilePath(filepath string) error {
	if strings.Contains(filepath, "..") || strings.Contains(filepath, "~") {
		return errors.New("invalid path")
	}
	return nil
}

func isFileExists(filepath string) (bool, error) {
	if _, err := os.Stat(filepath); err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, err
	} else {
		log.Fatal(err)
		return false, err
	}
}
