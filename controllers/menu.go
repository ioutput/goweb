package controllers

import (
	"log"
	"strconv"
	"net/http"
	"github.com/ioutput/goweb/models"
	"github.com/ioutput/goweb/service"
	"github.com/gin-gonic/gin"
)



func GetMenu(c *gin.Context) {
	
	param := make(map[string]string)
	param["page"] = c.DefaultQuery("page","1")
	param["page_size"] = c.DefaultQuery("page_size","10")
	param["name"] = c.DefaultQuery("name","")
	param["status"] = c.DefaultQuery("status","")
	data := service.ListMenu(param)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data})
}

func CreateMenu(c *gin.Context) { 
	
	var (
		menu models.Menu
		data string
	)
	c.ShouldBind(&menu)
	err :=models.DB.Create(&menu).Error
	code := http.StatusOK
	data = "添加成功"
	if err != nil{
		code = http.StatusBadRequest
		data = "添加失败"
		log.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"status": code, "data": data})
}

func UpdateMenu(c *gin.Context) {
	var data models.Menu
	var result string
	uid,_ := strconv.Atoi(c.Param("id"))
	/*绑定数据到模型*/
	c.ShouldBind(&data)
	err := models.DB.Model(&data).Where("id=?",uid).Update(data).Error
	code := http.StatusOK
	result = "更新成功"
	if err != nil {
		result = "更新失败"
		code = http.StatusBadRequest
	}
	c.JSON(http.StatusOK, gin.H{"status": code, "data": result})
}

func ViewMenu(c *gin.Context) {
	var menu models.Menu
	uid,_ := strconv.Atoi(c.Param("id"))
	data := models.DB.First(&menu,uid).Value
	//data := models.ViewUser(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data})
}

func DeleteMenu(c *gin.Context) {
	var data models.Menu
	uid,_ := strconv.Atoi(c.Param("id"))
	err := models.DB.Where("id=?",uid).Delete(&data).Error
	status := http.StatusOK
	result := "删除成功"
	if err != nil {
		result = "删除失败"
		status = http.StatusBadRequest
	}
	c.JSON(http.StatusOK, gin.H{"status": status, "data": result})
}

func LevelMenu(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": service.LevelMenu()})
}