package models

import (
	"github.com/goweb/utils"
	"github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
)


type Model struct {
	Id 			uint 				`gorm:"primary_key" json:"id"`
	CreatedAt   utils.JSONTime		`json:"created_at"`
	UpdatedAt   utils.JSONTime		`json:"updated_at"`
	DeletedAt  *utils.JSONTime 		`sql:"index" json:"-" form:"deleted_at"`
}
var DB *gorm.DB
func init() {
	var err error
	DB,err = gorm.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")
	//defer db.Close()
	if err != nil {
		panic(err)
	}
}
