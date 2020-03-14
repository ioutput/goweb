package models

import (
	"fmt"
	"goweb/utils"
	"goweb/conf" 
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
	mysql := conf.Settings.Mysql
	DB,err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",mysql.Username, mysql.Password, mysql.Host,mysql.Port , mysql.Name))
	//defer db.Close()
	 if err != nil {
		panic(err)
	} 
}
