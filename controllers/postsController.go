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

func PostIndex(c *gin.Context) {
	// Get posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with posts
	c.JSON(200, gin.H{
		"posts": posts,
	})

}

func PostShow(c *gin.Context) {
	// Get post
	var post models.Post

	var body struct {
		ID uint `json:"id"`
	}

	c.Bind(&body)

	initializers.DB.First(&post, body.ID)

	// Respond with post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	// Get the id off the params

	// Get data off req body
	var body struct {
		ID    uint   `json:"id"`
		Body  string `json:"body"`
		Title string `json:"title"`
	}

	c.Bind(&body)

	var post models.Post
	initializers.DB.First(&post, body.ID)
	// Update the post
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	// Respond with post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {

	// Delete the post
	initializers.DB.Delete(&models.Post{}, c.Param("id"))

	// Respond with post
	c.Status(200)
}
