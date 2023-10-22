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

type TransformImageSetting struct {
	ID           int64
	ImageName    string
	OutputType   string
	ResizeWidth  float32
	ResizeHeight float32
}

func GetSettings(imageName string) (TransformImageSetting, error) {
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASS"),
		Net:                  "tcp",
		Addr:                 os.Getenv("MYSQL_HOST") + ":" + os.Getenv("PORT"),
		DBName:               os.Getenv("DBNAME"),
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

	setting, err := TransformImageSettingByImageName(imageName)
	if err != nil {
		log.Fatal(err)
	}

	return setting, err
}

func TransformImageSettingByImageName(imageName string) (TransformImageSetting, error) {
	var setting TransformImageSetting

	row := db.QueryRow("SELECT * FROM transform_image_settings WHERE image_name = ?", imageName)
	if err := row.Scan(&setting.ID, &setting.ImageName, &setting.OutputType, &setting.ResizeHeight, &setting.ResizeWidth); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return setting, fmt.Errorf("TransformImageSettingsByImageName %d: no such image", imageName)
		}
		return setting, fmt.Errorf("TransformImageSettingsByImageName %d: %v", imageName, err)
	}
	return setting, nil
}
