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
	//router.OPTIONS("/login",controllers.Login)
	router.POST("/login",controllers.Login)
	v1 := router.Group("/api/",middleware.CheckToken)
	{
		v1.GET("user",controllers.GetUser)
		v1.GET("user/:id",controllers.ViewUser)
		v1.PUT("user/:id",controllers.UpdateUser)
		v1.DELETE("user/:id",controllers.DeleteUser)
		v1.POST("user",controllers.CreateUser)

		v1.GET("menu",controllers.GetMenu)
		v1.GET("menu/:id",controllers.ViewMenu)
		v1.PUT("menu/:id",controllers.UpdateMenu)
		v1.DELETE("menu/:id",controllers.DeleteMenu)
		v1.POST("menu",controllers.CreateMenu)

		v1.GET("role",controllers.GetRole)
		v1.GET("role/:id",controllers.ViewRole)
		v1.PUT("role/:id",controllers.UpdateRole)
		v1.DELETE("role/:id",controllers.DeleteRole)
		v1.POST("role",controllers.CreateRole)
	}	
	//defer models.db.Close()
	router.Run(":3008")
}
