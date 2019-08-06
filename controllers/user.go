package controllers

import (
	"strconv"
	"net/http"
	"github.com/ioutput/goweb/models"
	"github.com/ioutput/goweb/service"
	"github.com/gin-gonic/gin"
)



func GetUser(c *gin.Context) {
	
	param := make(map[string]string)
	param["page"] = c.DefaultQuery("page","1")
	param["page_size"] = c.DefaultQuery("page_size","10")
	param["username"] = c.DefaultQuery("username","")
	param["status"] = c.DefaultQuery("status","")
	data := service.ListUser(param)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data})
}

func CreateUser(c *gin.Context) { 
	
	var user models.User
	c.ShouldBind(&user)
	result := service.Register(user)
	//result := models.CreateUser(user)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": result})
}

func UpdateUser(c *gin.Context) {
	var data models.User
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

func ViewUser(c *gin.Context) {
	var user models.User
	uid,_ := strconv.Atoi(c.Param("id"))
	data := models.DB.First(&user,uid).Value
	//data := models.ViewUser(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data})
}

func DeleteUser(c *gin.Context) {
	var data models.User
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

