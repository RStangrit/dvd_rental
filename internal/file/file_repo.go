package file

import (
	"os"
)

// Dependency Injection Principle violated here, needs to be rewritten to interfaces
type FileRepository struct {
}

func NewFileRepository() *FileRepository {
	return &FileRepository{}
}

func (repo *FileRepository) FileExists(filepath string) (bool, error) {
	_, err := os.Stat(filepath)
	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
