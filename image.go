package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func resizeImageBySeparateRatios(inputPath string, xRatio, yRatio float32) (string, error) {

	outputPath := filepath.Dir(inputPath) + "/resized_" + filepath.Base(inputPath)

	file, err := os.Open(inputPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return "", err
	}

	newWidth := uint(float32(img.Bounds().Dx()) * xRatio)
	newHeight := uint(float32(img.Bounds().Dy()) * yRatio)

	resizedImg := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)

	out, err := os.Create(outputPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	return outputPath, jpeg.Encode(out, resizedImg, nil)
}

func ConvertImageFormat(inputPath, outputFormat string) (string, error) {
	currentTime := time.Now()
	timeString := currentTime.Format("20060102150405")

	outputPath := "./gcs/tmp/transform" + "/" + timeString + "_" + "formatted_" + filepath.Base(inputPath)

	file, err := os.Open(inputPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return "", err
	}

	out, err := os.Create(outputPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	switch strings.ToLower(outputFormat) {
	case "jpeg":
		err = jpeg.Encode(out, img, nil)
	case "png":
		err = png.Encode(out, img)
	default:
		return "", fmt.Errorf("unsupported output format: %s", outputFormat)
	}

	return outputPath, err
}
