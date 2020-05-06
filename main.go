package main

import "github.com/gin-gonic/gin"

func HelloHandler(c *gin.Context)  {
	c.JSON(200, gin.H{
		"msg": "hello dao cloud",
	})
}

func ByeHandler(c *gin.Context)  {
	name := c.Query("name")
	if name == "" {
		name = "friend"
	}
	c.JSON(200, gin.H{
		"msg": "bye" + name + "!",
	})
}

func main() {
	r := gin.Default()
	r.GET("/ping", HelloHandler)
	r.GET("/bye", ByeHandler)
	r.Run(":8082")
}