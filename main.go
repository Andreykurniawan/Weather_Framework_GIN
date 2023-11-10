package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.Default()

	route := app
	route.GET("/ping", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"Hallo": "World",
		})
		return
	})
	app.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
