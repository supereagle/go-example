package main

import (
	"fmt"
	"os"
)

func FileExists(path string) bool {
	// First check the err, if the file exists, err is nil
	if _, err := os.Stat(path); err != nil {
		return true
	}

	return false
}

func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func DirExists(path string) bool {
	fileInfo, err := os.Stat(path)
	return (err == nil || os.IsExist(err)) && fileInfo.IsDir()
}

func AppendToFile(path, content string) (err error) {
	// Check the exist of file, if not exist, create it
	if !FileExists(path) {
		file, cErr := os.Create(path)
		defer file.Close()
		if cErr != nil {
			return fmt.Errorf("File %s does not exist, fail to create it as %s", path, cErr.Error())
		}
	}

	// Open the file as O_APPEND modal
	file, err := os.OpenFile(path, os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		return fmt.Errorf("Fail to open the file %s\n", path)
	}

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("Fail to write the content %s into file %s", content, path)
	}

	return
}
