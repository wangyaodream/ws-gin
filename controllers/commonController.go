package controllers

import "github.com/gin-gonic/gin"


func Test(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "this is a test",
    })
}

func GetUserInfo(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "you have information with user",
    })
}
