package routers

import (
	"net/http"
	"nwd/controller/waiting"

	"github.com/gin-gonic/gin"
)

func Routers() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ping pong!")
	})

	v1 := r.Group("/api/v1")
	{
		v1.POST("/waiting", waiting.Create)
	}

	r.Run(":8000")
}
