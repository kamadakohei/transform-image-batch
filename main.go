package main

import (
	"fmt"
	"os"
	"transform-image-batch/gcs"
)

func main() {
	bucketName := "transform-image-bucket"
	downloadObjectName := "download/sample.jpeg"
	uploadObjectName := "upload/notes.txt"
	destinationFileName := "sample.jpeg"

	err := gcs.DownloadFile(os.Stdout, bucketName, downloadObjectName, destinationFileName)
	if err != nil {
		fmt.Printf("Error downloading file: %v\n", err)
		return
	}

	//db.GetSettings()

	err = resizeImageBySeparateRatios("./gcs/tmp/download/sample.jpeg", "./gcs/tmp/download/sample_resize_test.jpg", 1.5, 2.0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Image resized successfully!")
	}

	//// PNG を JPEG に変換
	//err = ConvertImageFormat("sample.png", "output.jpg", "jpeg")
	//if err != nil {
	//	fmt.Println("Error:", err)
	//}

	// JPEG を PNG に変換
	err = ConvertImageFormat("./gcs/tmp/download/sample.jpeg", "./gcs/tmp/download/sample_trans_test.png", "png")
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = gcs.UploadFile(os.Stdout, bucketName, uploadObjectName)
	if err != nil {
		fmt.Printf("Error uploading file: %v\n", err)
		return
	}
}
