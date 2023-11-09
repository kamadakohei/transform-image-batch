package main

import (
	"io"
	"os"
	"path/filepath"
	"testing"
)

func createTestDir(t *testing.T) string {
	t.Helper()
	dir, err := os.MkdirTemp("", "image_test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	return dir
}

func copyTestFile(t *testing.T, src, dst string) {
	t.Helper()
	sourceFile, err := os.Open(src)
	if err != nil {
		t.Fatalf("テストファイルのオープンに失敗しました: %v", err)
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		t.Fatalf("テストファイルの作成に失敗しました: %v", err)
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		t.Fatalf("テストファイルのコピーに失敗しました: %v", err)
	}
}

func TestResizeImageBySeparateRatios(t *testing.T) {
	testDir := createTestDir(t)
	defer os.RemoveAll(testDir)

	testImagePath := "./testdata/test.jpeg"
	inputPath := filepath.Join(testDir, "test.jpeg")
	copyTestFile(t, testImagePath, inputPath)
	xRatio, yRatio := float32(0.5), float32(0.5)

	outputPath, err := resizeImageBySeparateRatios(inputPath, xRatio, yRatio)
	if err != nil {
		t.Errorf("resizeImageBySeparateRatios returned an error: %v", err)
	}

	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Fatalf("The output file does not exist: %s", outputPath)
	}

	originalImg, _, err := loadImage(testImagePath)
	if err != nil {
		t.Fatalf("Failed to load original image: %v", err)
	}

	resizedImg, _, err := loadImage(outputPath)
	if err != nil {
		t.Fatalf("Failed to load resized image: %v", err)
	}

	expectedWidth := uint(float32(originalImg.Bounds().Dx()) * xRatio)
	expectedHeight := uint(float32(originalImg.Bounds().Dy()) * yRatio)

	if uint(resizedImg.Bounds().Dx()) != expectedWidth || uint(resizedImg.Bounds().Dy()) != expectedHeight {
		t.Errorf("Resized image dimensions are incorrect: got width %d, height %d, want width %d, height %d",
			resizedImg.Bounds().Dx(), resizedImg.Bounds().Dy(), expectedWidth, expectedHeight)
	}
}

func TestConvertImageFormat(t *testing.T) {
	testDir := createTestDir(t)
	defer os.RemoveAll(testDir)

	testImagePath := "./testdata/test.jpeg"
	inputPath := filepath.Join(testDir, "test.jpeg")
	copyTestFile(t, testImagePath, inputPath)

	outputFormat := "png"

	outputPath, err := ConvertImageFormat(inputPath, outputFormat)
	if err != nil {
		t.Errorf("ConvertImageFormat returned an error: %v", err)
	}

	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Errorf("Output file does not exist: %s", outputPath)
	}

	outputFileExtention := GetFileExtension(outputPath)
	expectedFileExtention := "png"
	if outputFileExtention != expectedFileExtention {
		t.Errorf("Expected output file extention at %s, but got %s", expectedFileExtention, outputFileExtention)
	}
}
