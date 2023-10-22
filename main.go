package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"transform-image-batch/db"
	"transform-image-batch/gcs"
)

func main() {
	var bucketName, downloadObjectName string

	flag.StringVar(&bucketName, "b", "", "Name of the bucket")
	flag.StringVar(&downloadObjectName, "o", "", "Name of the object to download")

	flag.Parse()

	if bucketName == "" || downloadObjectName == "" {
		fmt.Println("All options are required")
		flag.Usage()
		return
	}

	fmt.Println("Bucket Name:", bucketName)
	fmt.Println("Download Object Name:", downloadObjectName)

	ctx, client, err := gcs.InitGCSClient()
	if err != nil {
		fmt.Printf("Error initDBClient: %v\n", err)
		return
	}

	setting, err := db.GetSettings(filepath.Base(downloadObjectName))
	if err != nil {
		fmt.Printf("Error getSettings: %v\n", err)
		return
	}
	fmt.Printf("setting: %v\n", setting)

	destFileName, err := gcs.DownloadFile(ctx, client, os.Stdout, bucketName, downloadObjectName)
	if err != nil {
		fmt.Printf("Error downloading file: %v\n", err)
		return
	}

	resizedFilePath, err := resizeImageBySeparateRatios(destFileName, setting.ResizeWidth, setting.ResizeHeight)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Image resized successfully!")
	}

	formattedFilePath, err := ConvertImageFormat(resizedFilePath, getFileExtension(filepath.Base(resizedFilePath)))
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = gcs.UploadFile(os.Stdout, bucketName, formattedFilePath)
	if err != nil {
		fmt.Printf("Error uploading file: %v\n", err)
		return
	}
}

func getFileExtension(fileName string) string {
	extensionWithDot := filepath.Ext(fileName)
	extension := strings.TrimPrefix(extensionWithDot, ".")
	return extension
}
