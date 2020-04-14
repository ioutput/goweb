package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"goweb/conf"
	"goweb/controllers"
	"goweb/middleware"
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {
	startCmd := &cobra.Command{
		Use:   "start",
		Short: "start web server",
		Long:  "start web server port 3009",
		Run: func(cmd *cobra.Command, args []string) {
			start()
			command := exec.Command("goweb", "start")
			command.Start()
			fmt.Printf("goweb start, [PID] %d running...\n", command.Process.Pid)
			ioutil.WriteFile("/run/goweb.pid", []byte(fmt.Sprintf("%d", command.Process.Pid)), 0666)
			os.Exit(0)

		},
	}
	var stopCmd = &cobra.Command{
		Use:   "stop",
		Short: "Stop goweb",
		Run: func(cmd *cobra.Command, args []string) {
			strb, _ := ioutil.ReadFile("/run/goweb.pid")
			command := exec.Command("kill", string(strb))
			command.Start()
			println("goweb stop")
		},
	}
	//startCmd.Flags().BoolVarP(&daemon, "deamon", "d", false, "is daemon?")
	var rootCmd = &cobra.Command{Use: "goweb"}
	rootCmd.AddCommand(startCmd, stopCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func start() {
	if err := conf.Init(); err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	router := gin.Default()
	router.Use(middleware.Cors())
	router.OPTIONS("/login", controllers.Login)
	router.POST("/login", controllers.Login)
	v1 := router.Group("/api/", middleware.CheckToken)
	{
		v1.GET("user", controllers.GetUser)
		v1.GET("user/:id", controllers.ViewUser)
		v1.PUT("user/:id", controllers.UpdateUser)
		v1.DELETE("user/:id", controllers.DeleteUser)
		v1.POST("user", controllers.CreateUser)

		v1.GET("menu", controllers.GetMenu)
		v1.GET("levelmenu", controllers.LevelMenu)
		v1.GET("menu/:id", controllers.ViewMenu)
		v1.PUT("menu/:id", controllers.UpdateMenu)
		v1.DELETE("menu/:id", controllers.DeleteMenu)
		v1.POST("menu", controllers.CreateMenu)

		v1.GET("role", controllers.GetRole)
		v1.GET("role/:id", controllers.ViewRole)
		v1.PUT("role/:id", controllers.UpdateRole)
		v1.DELETE("role/:id", controllers.DeleteRole)
		v1.POST("role", controllers.CreateRole)
	}
	//defer models.db.Close()
	router.Run(":3009")
}
