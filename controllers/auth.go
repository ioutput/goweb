package controllers

import (
	"net/http"
	"goweb/service"
	"goweb/middleware"
	"github.com/gin-gonic/gin"
)



func Login(c *gin.Context) {
	var (
		token string
	 	data interface{}
	)
	user := service.Login(c.PostForm("username"),c.PostForm("password"))
	code := http.StatusOK
	if user["code"] == 200 {
		data = user["data"]
		token,_ = middleware.CreateToken()
	}else{
		code = 400
		data = user["msg"]
	}
	c.JSON(http.StatusOK, gin.H{"status": code, "data": data,"token":token})
}
