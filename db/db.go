package db

import (
	"gorm.io/driver/sqlite" // Sqlite driver based on GGO
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnDB 连接数据库
func ConnDB(filepath string) error {

	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open(filepath), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db
	return nil
}
