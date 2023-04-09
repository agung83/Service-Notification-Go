package main

import (
	"serviceNotification/models"
	"serviceNotification/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

/*
* TODO HANDLE CLEAN ARCHITECTURE AND RESPOSITORY PATTERN
 */

func main() {
  err := godotenv.Load()
  if err != nil {
    logrus.Fatal("Error loading .env file")
  }
    models.SetupDatabase()
    app := setupRouter()
    logrus.Fatal(app.Run())
}

func setupRouter() *gin.Engine  {
  app := gin.Default()
  routes.InitNotificationRoutes(app)
  return app
}



