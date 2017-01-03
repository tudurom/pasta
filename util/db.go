package util

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/tudurom/pasta/model"
)

var GlobalDB *gorm.DB

func ConnectToDB() {
	var err error
	GlobalDB, err = gorm.Open("sqlite3", GlobalConfig.DatabasePath)
	Crash("Couldn't connect to database", err)
}

func MigrateSchema() {
	GlobalDB.AutoMigrate(&model.Paste{})
}
