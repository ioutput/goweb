package controllers

import (
	"log"
	"net/http"
	"github.com/ioutput/goweb/service"
	"github.com/ioutput/goweb/middleware"
	"github.com/gin-gonic/gin"
)



func Login(c *gin.Context) {
	var token string
	var data interface{}
	log.Println(c.PostForm("username"),c.PostForm("password"))
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
