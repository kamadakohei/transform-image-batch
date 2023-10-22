package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	TransformDir = "./gcs/tmp/transform"
)

func resizeImageBySeparateRatios(inputPath string, xRatio, yRatio float32) (string, error) {
	fmt.Println(inputPath)
	outputPath := filepath.Dir(inputPath) + "/resized_" + filepath.Base(inputPath)
	inputFileFormat := GetFileExtension(inputPath)

	img, _, err := loadImage(inputPath)
	if err != nil {
		return "", err
	}

	newWidth := uint(float32(img.Bounds().Dx()) * xRatio)
	newHeight := uint(float32(img.Bounds().Dy()) * yRatio)

	resizedImg := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)

	out, err := createFile(outputPath)
	if err != nil {
		return "", err
	}

	return outputPath, encodeImage(out, resizedImg, inputFileFormat)
}

func ConvertImageFormat(inputPath, outputFormat string) (string, error) {
	currentTime := time.Now()
	timeString := currentTime.Format("20060102150405")

	outputPath := filepath.Join(TransformDir, timeString+"_formatted_"+GetFileNameWithoutExtension(filepath.Base(inputPath))+"."+outputFormat)

	img, _, err := loadImage(inputPath)
	if err != nil {
		return "", err
	}

	out, err := createFile(outputPath)
	if err != nil {
		return "", err
	}

	err = encodeImage(out, img, outputFormat)
	if err != nil {
		return "", err
	}

	return outputPath, err
}

func loadImage(inputPath string) (image.Image, string, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	return img, format, err
}

func createFile(path string) (*os.File, error) {
	out, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func encodeImage(w io.Writer, img image.Image, format string) error {
	switch strings.ToLower(format) {
	case "jpeg":
		return jpeg.Encode(w, img, nil)
	case "png":
		return png.Encode(w, img)
	default:
		return fmt.Errorf("unsupported output format: %s", format)
	}
}
