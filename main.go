package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyaodream/ws-gin/controllers"
	"github.com/wangyaodream/ws-gin/initializers"
	"github.com/wangyaodream/ws-gin/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

    // add group
    adminGroup := r.Group("/admin")
    apiGroup := r.Group("/api")

    authorized := r.Group("/auth")

    authorized.Use(middleware.UserAuth())
    {
        authorized.GET("/getinfo", controllers.GetInfo)
    }

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/index", controllers.PostIndex)
	r.POST("/show", controllers.PostShow)
	r.PUT("/update", controllers.PostUpdate)
	r.DELETE("/delete/:id", controllers.PostDelete)

    adminGroup.GET("/login", controllers.GetInfo)
    apiGroup.GET("/test", controllers.Test)

	r.Run()
}
