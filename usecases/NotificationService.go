package usecases

import (
	"net/http"
	dbrepositories "serviceNotification/repositories/db"

	"github.com/gin-gonic/gin"
)


func GetAllNotificationServices(c*gin.Context)   {
		resDb := dbrepositories.GetAllNotification()
		c.Header("agung", "laksmana")
		c.JSON(http.StatusOK, gin.H{
		  "message": "Success",
		  "response" : resDb,
		  "oke nee" : "www mantul",
		})
}



func StoreOneNotificationServices() int  {
	return 1
}