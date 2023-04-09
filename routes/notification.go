package routes

import (
	"serviceNotification/usecases"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Categories struct {

	Id int `json:"id"`
	Name string `json:"name"`
  }

func InitNotificationRoutes(db *gorm.DB , route *gin.Engine)   {

	groupRoute := route.Group("/api/1.0.0/notification")
	groupRoute.GET("/getall",	func (c*gin.Context){
		usecases.GetAllNotificationServices(db,c)
	})
	
}