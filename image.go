package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func resizeImageBySeparateRatios(inputPath, outputPath string, xRatio, yRatio float64) error {
	// 入力ファイルを開く
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 画像をデコード
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	// リサイズの新しい幅と高さを計算
	newWidth := uint(float64(img.Bounds().Dx()) * xRatio)
	newHeight := uint(float64(img.Bounds().Dy()) * yRatio)

	// 画像をリサイズ
	resizedImg := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)

	// リサイズした画像を出力ファイルに保存
	out, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer out.Close()

	return jpeg.Encode(out, resizedImg, nil)
}

// ConvertImageFormat は指定された入力ファイルを指定された出力形式に変換します。
func ConvertImageFormat(inputPath, outputPath, outputFormat string) error {
	// 入力ファイルを開く
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 画像をデコード
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	// 出力ファイルを作成
	out, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer out.Close()

	// 指定されたフォーマットでエンコード
	switch strings.ToLower(outputFormat) {
	case "jpeg":
		err = jpeg.Encode(out, img, nil)
	case "png":
		err = png.Encode(out, img)
	default:
		return fmt.Errorf("unsupported output format: %s", outputFormat)
	}

	return err
}
