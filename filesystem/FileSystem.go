package filesystem

import "os"

type FileSystem interface {
	ReadFile(filename string) ([]byte, error)
	// update file will create it if it's not existed
	UpdateFile(filename string, content []byte) error
	List(rel string) ([]os.FileInfo, error)
	FileExists(filename string) (bool, error)
}
