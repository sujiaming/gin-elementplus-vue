package router

import (
	"wanworld/controller"

	"github.com/gin-gonic/gin"
)

func SetupAPIRouter(r *gin.RouterGroup) {

	// 定义路由
	r.GET("/users", controller.GetUsers)
	// r.POST("/users", controller.CreateUser)
	// r.GET("/users/:id", controller.GetUser)
	// r.PUT("/users/:id", controller.UpdateUser)
	// r.DELETE("/users/:id", controller.DeleteUser)
}
