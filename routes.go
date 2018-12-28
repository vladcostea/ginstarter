package ginstarter

import "github.com/gin-gonic/gin"

func ping(c *gin.Context) {
	c.String(200, "pong")
}

func NewRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", ping)
	return r
}
