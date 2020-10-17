package fs

import (
	"fmt"
	"os"
)

//noinspection GoUnusedExportedFunction
func Exists(path string) (bool, error) {
	fileInfo, err := check(path)
	return fileInfo != nil, err
}

func RequireDirectory(path string) error {
	fileInfo, err := check(path)
	if err != nil {
		return err
	}
	if fileInfo == nil {
		return fmt.Errorf("'%s' did not exist", path)
	}
	if !(*fileInfo).IsDir() {
		return fmt.Errorf("'%s' existed, but was not a Directory", path)
	}
	return nil
}

//noinspection GoUnusedExportedFunction
func IsDirectory(path string) (bool, error) {
	fileInfo, err := check(path)
	if (err != nil) || (fileInfo == nil) {
		return false, err
	}
	return (*fileInfo).IsDir(), nil
}

//noinspection GoUnusedExportedFunction
func IsFile(path string) (bool, error) {
	fileInfo, err := check(path)
	if (err != nil) || (fileInfo == nil) {
		return false, err
	}
	return !(*fileInfo).IsDir(), nil
}

func check(path string) (*os.FileInfo, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		// Checking if the given file exists or not
		// Using IsNotExist() function
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	return &fileInfo, nil
}
