package filesystem

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type LocalFileSystem struct {
	rootDir string
}

func (lfs *LocalFileSystem) checkScope(path string) {
	rootAbs, _ := filepath.Abs(lfs.rootDir)
	fp, _ := filepath.Abs(filepath.Join(lfs.rootDir, path))
	logrus.Debug(rootAbs)
	logrus.Debug(fp)
	if !strings.HasPrefix(fp, rootAbs) {
		panic("access denied! out of scope !!!")
	}
}

func (lfs *LocalFileSystem) ReadFile(filepath string) ([]byte, error) {
	lfs.checkScope(filepath)
	filePath := path.Join(lfs.rootDir, filepath)
	return ioutil.ReadFile(filePath)
}

func (lfs *LocalFileSystem) UpdateFile(filepath string, content []byte) error {
	lfs.checkScope(filepath)
	return ioutil.WriteFile(path.Join(lfs.rootDir, filepath), content, 0666)
}

func (lfs *LocalFileSystem) List(rel string) ([]os.FileInfo, error) {
	lfs.checkScope(rel)
	logrus.Debug(path.Join(lfs.rootDir, rel))
	return ioutil.ReadDir(path.Join(lfs.rootDir, rel))
}

func (lfs *LocalFileSystem) FileExists(filepath string) (bool, error) {
	lfs.checkScope(filepath)
	if _, err := os.Stat(path.Join(lfs.rootDir, filepath)); os.IsNotExist(err) {
		return false, nil
	} else if os.IsExist(err) {
		return true, nil
	} else {
		return err == nil, err
	}
}

func (lfs *LocalFileSystem) Delete(filepath string) error {
	lfs.checkScope(filepath)
	return os.Remove(path.Join(lfs.rootDir, filepath))
}

func NewLocalFileSystem(root string) *LocalFileSystem {
	return &LocalFileSystem{rootDir:root}
}