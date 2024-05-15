package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyaodream/ws-gin/initializers"
	"github.com/wangyaodream/ws-gin/models"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Create a post
	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// Return

	c.JSON(200, gin.H{
		"post": post,
	})
}

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
