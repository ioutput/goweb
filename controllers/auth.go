package controllers

import (
	//"strconv"
	"net/http"
	"demo/service"
	"demo/middleware"
	"github.com/gin-gonic/gin"
)



func Login(c *gin.Context) {
	data := make(map[string]interface{})
	user := service.Login(map[string]string{"username":c.PostForm("username"),"password":c.PostForm("password")})
	code := http.StatusOK
	if user["code"] == 200 {
		data["data"] = user
		data["token"],_ = middleware.CreateToken()
	}else{
		code = 400
		data["data"] = user["msg"]
	}
	c.JSON(http.StatusOK, gin.H{"status": code, "data": data})
}
