package main

import (
	"os"
	"serviceNotification/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
* TODO HANDLE CLEAN ARCHITECTURE AND RESPOSITORY PATTERN
 */

func main() {
  err := godotenv.Load()
  if err != nil {
    logrus.Fatal("Error loading .env file")
  }
    app := setupRouter()
    logrus.Fatal(app.Run())
}



func setupRouter() *gin.Engine  {
  db := setupDatabase()
  app := gin.Default()

  routes.InitNotificationRoutes(db,app)
  return app
}

func setupDatabase() *gorm.DB {
    db,err := gorm.Open(mysql.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
    if err != nil {
      defer logrus.Info("Connect into Database Failed")
      logrus.Fatal(err.Error())
    }
    return db
}


