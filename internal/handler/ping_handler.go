package handler

import "github.com/gin-gonic/gin"

func Ping(c *gin.Context) {
	c.SetCookie("ping_cookie", "ping_value", 3600, "/", "localhost", false, true)
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
