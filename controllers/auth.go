package controllers

import (
	//"strconv"
	"net/http"
	"github.com/goweb/service"
	"github.com/goweb/middleware"
	"github.com/gin-gonic/gin"
)



func Login(c *gin.Context) {
	var token string
	var data interface{}
	user := service.Login(map[string]string{"username":c.PostForm("username"),"password":c.PostForm("password")})
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
