package database

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	// SQLite na memória
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&User{})

}
