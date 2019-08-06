package service

import (
	"fmt"
	"strconv"
	"github.com/ioutput/goweb/models"
)

type LeMenu struct {
	models.Menu
	Key   	uint `json:"key"`
	Value   uint `json:"value"`
	Title   string `json:"title"`
	Child []LeMenu `json:"children"`
}

//查询列表
func ListMenu(params map[string]string) (result []models.Menu) {
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

func LevelMenu() (lists []LeMenu) {
	var data []LeMenu
	
	err := models.DB.Table("menus").Scan(&data).Error
	if err != nil {
		panic(err)
	}
	lists = getChilds(data,0)
	/* if len(lists) > 0 {
		for k,v := range lists {
			v.
			data[k] = v
		}
	} */
	return
}

func getChilds(data []LeMenu , pid uint) (result []LeMenu) {
	
	for _,v := range data {
		if v.Pid == pid {
			v.Key = v.Id
			v.Value = v.Id
			v.Title = v.Name
			v.Child = getChilds(data,v.Id)
			result = append(result,v)
		}
		
	}
	return
}