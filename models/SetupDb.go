package models

import (
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func SetupDatabase() *gorm.DB {
    db,err := gorm.Open(mysql.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
    if err != nil {
      defer logrus.Info("Connect into Database Failed")
      logrus.Fatal(err.Error())
    }
    return db
}

