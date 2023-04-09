package dbrepositories

import (
	"serviceNotification/models"
)


type Categories struct {

	Id int `json:"id"`
	Name string `json:"name"`
  }

  /*
  * TODO HANDLE WITH MODEL AND MIGRATION
  */
func GetAllNotification() []Categories {
	/*
	* this is sample for use raw query
	*/
	db := models.SetupDatabase()
	var result []Categories   
    db.Raw("SELECT * FROM categories").Scan(&result)
	return result
}