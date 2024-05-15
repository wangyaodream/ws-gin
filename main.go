package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyaodream/ws-gin/controllers"
	"github.com/wangyaodream/ws-gin/initializers"
)

func init() {
    initializers.LoadEnvVariables()
    initializers.ConnectToDB()
}

func main() {
    r := gin.Default()

    r.POST("/posts", controllers.PostsCreate) 

    r.Run()
}
