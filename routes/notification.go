package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Categories struct {

	Id int `json:"id"`
	Name string `json:"name"`
  }

func InitNotificationRoutes(db *gorm.DB , route *gin.Engine)   {

	groupRoute := route.Group("/api/1.0.0/notification")
	groupRoute.GET("/getall",func(c *gin.Context) {
		/*
		* TODO handle in usecases
		*/
	var result []Categories   
    db.Raw("SELECT * FROM categories").Scan(&result)

  // // user := Categories{}
  // // db.Raw("INSERT INTO categories (id,name) VALUES (8,'LAKSMANA jumat merah')").Scan(&user)

    c.JSON(http.StatusOK, gin.H{
      "message": "mantap",
      "response" : result,
    })

	
})
	

	
}