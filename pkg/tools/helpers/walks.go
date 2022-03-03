package helpers

import (
	"io/ioutil"
	"os"
)

func Walk(dir string) ([]os.FileInfo, error) {
	var files []os.FileInfo
	fs, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range fs {
		filesFromDir := []os.FileInfo{file}
		if file.IsDir() {
			filesFromDir, err = Walk(file.Name())
			if err != nil {
				return nil, err
			}
		}
		files = append(files, filesFromDir...)
	}

	return files, nil
}
