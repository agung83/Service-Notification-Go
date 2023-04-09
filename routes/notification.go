package routes

import (
	servicesNotification "serviceNotification/usecases"

	"github.com/gin-gonic/gin"
)

type Categories struct {

	Id int `json:"id"`
	Name string `json:"name"`
  }

func InitNotificationRoutes(route *gin.Engine)   {

	groupRoute := route.Group("/api/1.0.0/notification")
	groupRoute.GET("/getall",servicesNotification.GetAllNotificationServices)
	
}