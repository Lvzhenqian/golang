package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	v1 := r.Group("v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200,gin.H{
				"message": "pong",
			})
		})
		v1.POST("/post", func(c *gin.Context) {
			message := c.PostForm("message")
			nick := c.DefaultPostForm("nick","null")

			c.JSON(200,gin.H{
				"status": "postd",
				"message": message,
				"nick": nick,
			})
		})
	}



	r.Run()
}
