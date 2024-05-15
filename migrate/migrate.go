package main

import (
	"github.com/wangyaodream/ws-gin/initializers"
	"github.com/wangyaodream/ws-gin/models"
)

func init() {
    initializers.LoadEnvVariables()
    initializers.ConnectToDB()
}

func main() {
    initializers.DB.AutoMigrate(&models.Post{})
}
