package service

import (
	"fmt"
	"strconv"
	"github.com/goweb/models"
)



//查询列表
func ListRole(params map[string]string) (result []models.Role) {
	//转string为int
	page,_ := strconv.Atoi(params["page"])
	page_size,_ := strconv.Atoi(params["page_size"])
	offset := (page - 1) * page_size
	query := models.DB.Offset(offset).Limit(page_size)
	if params["name"] !="" {
		query = query.Where("Name LIKE ?" , fmt.Sprintf("%s%s%s","%",params["name"],"%"))
	}
	if params["status"] !="" {
		query = query.Where("Status =?",params["status"])
	}
	err :=query.Find(&result).Error
	if err != nil {
		panic(err)
	}
	return 
}
