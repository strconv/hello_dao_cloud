package main

import "github.com/gin-gonic/gin"

func HelloHandler(c *gin.Context)  {
	c.JSON(200, gin.H{
		"msg": "hello dao cloud",
	})
}

func main() {
	r := gin.Default()
	r.GET("/ping", HelloHandler)
	r.Run(":8082")
}