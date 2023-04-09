package usecases

import (
	"net/http"
	dbrepositories "serviceNotification/repositories/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func GetAllNotificationServices(db*gorm.DB,c*gin.Context)   {

		resDb := dbrepositories.GetAllNotification(db)
		c.Header("agung", "laksmana")
		c.JSON(http.StatusOK, gin.H{
		  "message": "Success",
		  "response" : resDb,
		  "oke nee" : "www",
		})
}



func StoreOneNotificationServices() int  {
	return 1
}