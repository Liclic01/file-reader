package service

import (
	"fmt"
	"os"
)

type FileReader interface {
	ReadFileContent(path string, fileType string, fileVersion string) (string, error)
}

type DefaultFileReader struct{}

func (fr *DefaultFileReader) ReadFileContent(path string, fileType string, fileVersion string) (string, error) {
	filePath := fmt.Sprintf("%s/%s/%s.json", path, fileType, fileVersion)
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}
	return string(fileContent), nil
}
