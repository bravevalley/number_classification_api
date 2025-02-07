package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DigiProp(c *gin.Context) {

	// Get number
	num := c.Query("number")

	// Sanity check
	if num == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"number": "invalid",
			"error":  true,
		})
		return
	}

	// verify whether an int was provided
	number, err := strconv.Atoi(num)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"number": "alphabet",
			"error":  true,
		})
		return
	}

	numberStat, err := SetupNum(number)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"number": "Valid",
			"error":  true,
		})
	}

	c.JSON(http.StatusOK, numberStat)

}

// CORS
func CorsMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "https://myfrontend.com")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, Authorization")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PATCH, DELETE, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusNoContent)
            return
        }

        c.Next()
    }
}
