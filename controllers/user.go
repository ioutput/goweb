package controllers

import (
	"strconv"
	"net/http"
	"demo/models"
	"demo/service"
	"github.com/gin-gonic/gin"
)



func GetUser(c *gin.Context) {
	
	param := make(map[string]string)
	param["page"] = c.DefaultQuery("page","1")
	param["page_size"] = c.DefaultQuery("page_size","10")
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
	if err != nil {
		result = "更新失败"
	}else{
		result = "更新成功"
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": result})
}

func ViewUser(c *gin.Context) {
	var user models.User
	uid,_ := strconv.Atoi(c.Param("id"))
	data := models.DB.First(&user,uid)
	//data := models.ViewUser(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data})
}

func DeleteUser(c *gin.Context) {
	//var result string
	var data models.User
	uid,_ := strconv.Atoi(c.Param("id"))
	err := models.DB.Where("id=?",uid).Delete(&data)
	/* if err != nil {
		result = "删除失败"
	}else{
		result = "删除成功"
	} */
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": err})
}

