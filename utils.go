package main

import (
	"path/filepath"
	"strings"
)

func GetFileExtension(fileName string) string {
	extensionWithDot := filepath.Ext(fileName)
	extension := strings.TrimPrefix(extensionWithDot, ".")
	return extension
}

func GetFileNameWithoutExtension(fileName string) string {
	fileNameWithoutExt := strings.TrimSuffix(filepath.Base(fileName), filepath.Ext(fileName))
	return fileNameWithoutExt
}
