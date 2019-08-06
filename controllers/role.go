package controllers

import (
	"strconv"
	"net/http"
	"github.com/ioutput/goweb/models"
	"github.com/ioutput/goweb/service"
	"github.com/gin-gonic/gin"
)



func GetRole(c *gin.Context) {
	
	param := make(map[string]string)
	param["page"] = c.DefaultQuery("page","1")
	param["page_size"] = c.DefaultQuery("page_size","10")
	param["name"] = c.DefaultQuery("name","")
	param["status"] = c.DefaultQuery("status","")
	data := service.ListRole(param)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data})
}

func CreateRole(c *gin.Context) { 
	var (
		role models.Role
		data interface{}
	)
	c.ShouldBind(&role)
	err :=models.DB.Create(&role).Error
	code := http.StatusOK
	data = role
	if err != nil{
		code = http.StatusBadRequest
		data = err
	}
	c.JSON(http.StatusOK, gin.H{"status": code, "data": data})
}

func UpdateRole(c *gin.Context) {
	var data models.Role
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

func ViewRole(c *gin.Context) {
	var model models.Role
	uid,_ := strconv.Atoi(c.Param("id"))
	data := models.DB.First(&model,uid).Value
	//data := models.ViewUser(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data})
}

func DeleteRole(c *gin.Context) {
	var data models.Role
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

