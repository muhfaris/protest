package utils

import (
	"io/fs"
	"io/ioutil"
	"os"

	e "gitlab.sicepat.tech/muhfaris/protest/pkg/errorc"
)

func CurrentDir() (path string, err error) {
	path, err = os.Getwd()
	if err != nil {
		return path, err
	}

	return path, err
}

func ReadFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, e.ErrReadTemplate.Errorf(err)
	}

	return data, nil
}

func ReadFiles(path string) ([]fs.FileInfo, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, e.ErrReadTemplate.Errorf(err)
	}

	return files, nil
}
