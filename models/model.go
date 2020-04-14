package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"goweb/conf"
	"goweb/utils"
)

type Model struct {
	Id        uint            `gorm:"primary_key" json:"id"`
	CreatedAt utils.JSONTime  `json:"created_at"`
	UpdatedAt utils.JSONTime  `json:"updated_at"`
	DeletedAt *utils.JSONTime `sql:"index" json:"-" form:"deleted_at"`
}

var DB *gorm.DB

func init() {
	var err error
	mysql := conf.Settings.Mysql
	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", mysql.Username, mysql.Password, mysql.Host, mysql.Port, mysql.Name))
	//defer db.Close()
	if err != nil {
		panic(err)
	}

}
