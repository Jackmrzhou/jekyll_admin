package filesystem

import "os"

type FileSystem interface {
	ReadFile(filepath string) ([]byte, error)
	// update file will create it if it's not existed
	UpdateFile(filepath string, content []byte) error
	List(rel string) ([]os.FileInfo, error)
	FileExists(filepath string) (bool, error)
	Delete(filepath string) error
}
