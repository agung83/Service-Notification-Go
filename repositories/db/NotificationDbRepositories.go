package dbrepositories

import (
	"gorm.io/gorm"
)


type Categories struct {

	Id int `json:"id"`
	Name string `json:"name"`
  }

  /*
  * TODO HANDLE WITH MODEL AND MIGRATION
  */
func GetAllNotification(db *gorm.DB) []Categories {
	/*
	* this is sample for use raw query
	*/
	var result []Categories   
    db.Raw("SELECT * FROM categories").Scan(&result)
	return result
}