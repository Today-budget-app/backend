// boilerplate yab
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/hello-world", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": true})
	})
	//add r.POST, etc...

	r.Run(":8080")
}
