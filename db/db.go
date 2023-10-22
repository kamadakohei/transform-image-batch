package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var db *sql.DB

type TransformImageSettings struct {
	ID           int64
	ImageName    string
	OutputType   string
	ResizeWidth  float32
	ResizeHeight float32
}

func GetSettings() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:8000",
		DBName:               "transform_image_settings",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	alb, err := TransformImageSettingsByImageName("example_file1.png")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transform_image_settings found: %v\n", alb)
}

func TransformImageSettingsByImageName(imageName string) (TransformImageSettings, error) {
	var settings TransformImageSettings

	row := db.QueryRow("SELECT * FROM transform_image_settings WHERE image_name = ?", imageName)
	if err := row.Scan(&settings.ID, &settings.ImageName, &settings.OutputType, &settings.ResizeHeight, &settings.ResizeWidth); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return settings, fmt.Errorf("TransformImageSettingsByImageName %d: no such image", imageName)
		}
		return settings, fmt.Errorf("TransformImageSettingsByImageName %d: %v", imageName, err)
	}
	return settings, nil
}
