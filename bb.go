package main

import (
	"github.com/gin-gonic/gin"
	"github.com/manjeshpv/bb/routes"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.Use(Cors())

	routes.Init(r)

	r.Run(":8001")
}

