package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
    app := setupRouter()

    logrus.Fatal(app.Run())
}


func setupRouter() *gin.Engine  {
  app := gin.Default()
  app.GET("/test/:name", func(c *gin.Context) {
	name := c.Param("name")
    c.JSON(http.StatusOK, gin.H{
      "message": "mantap",
	  "name" : name,
    })
  })
  return app
}


