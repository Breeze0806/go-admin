package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var opens = map[string]func(string) gorm.Dialector{
	"sqlite3":   sqlite.Open,
}
