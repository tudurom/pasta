package util

import (
	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/tudurom/pasta/models"
)

var GlobalDB *gorm.DB

func ConnectToDB() {
	logrus.Print("Connecting to DB")
	var err error
	GlobalDB, err = gorm.Open("sqlite3", GlobalConfig.DatabasePath)
	Crash("Couldn't connect to database", err)
}

func MigrateSchema() {
	logrus.Print("Migrating schemas")
	GlobalDB.AutoMigrate(&models.Paste{})
}
