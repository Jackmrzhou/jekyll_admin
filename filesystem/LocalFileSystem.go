package filesystem

import (
	"io/ioutil"
	"os"
	"path"
)

type LocalFileSystem struct {
	rootDir string
}

func (lfs *LocalFileSystem) ReadFile(filename string) ([]byte, error) {
	filePath := path.Join(lfs.rootDir, filename)
	return ioutil.ReadFile(filePath)
}

func (lfs *LocalFileSystem) UpdateFile(filename string, content []byte) error {
	return ioutil.WriteFile(path.Join(lfs.rootDir, filename), content, 0666)
}

func (lfs *LocalFileSystem) List(rel string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(path.Join(lfs.rootDir, rel))
}

func (lfs *LocalFileSystem) FileExists(filename string) (bool, error) {
	if _, err := os.Stat(path.Join(lfs.rootDir, filename)); os.IsNotExist(err) {
		return false, nil
	} else if os.IsExist(err) {
		return true, nil
	} else {
		return err == nil, err
	}
}

func NewLocalFileSystem(root string) *LocalFileSystem {
	return &LocalFileSystem{rootDir:root}
}