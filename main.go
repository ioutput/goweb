package main

import (
	//"net/http"
    "github.com/gin-gonic/gin"
	"github.com/goweb/controllers"
	"github.com/goweb/middleware"
)
func main() {
	
	router := gin.Default()
	router.Use(middleware.Cors())
	router.OPTIONS("/login",controllers.Login)
	router.POST("/login",controllers.Login)
	v1 := router.Group("/",middleware.CheckToken)
	{
		v1.GET("user",controllers.GetUser)
		v1.GET("/user/:id",controllers.ViewUser)
		v1.PUT("/user/:id",controllers.UpdateUser)
		v1.DELETE("/user/:id",controllers.DeleteUser)
		v1.POST("/user",controllers.CreateUser)
	}	
	//defer models.db.Close()
	router.Run(":3008")
}
