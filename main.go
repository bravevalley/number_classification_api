package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// Configure and start webserver.
	r := gin.Default()

	// CORS
	r.Use(CorsMiddleware())

	r.GET("/api/classify-number", DigiProp)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Nothing to see here :)",
		})
	})

	r.Run("0.0.0.0:80")

}
