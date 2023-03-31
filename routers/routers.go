package routers

import (
	"net/http"
	"nwd/controller/users"
	"nwd/controller/waiting"
	"nwd/middleware"
	"nwd/shared/response"

	"github.com/gin-gonic/gin"
)

func Routers() {
	r := gin.Default()

	//heart beat api
	r.GET("/ping", func(c *gin.Context) {
		response.Response(c.Writer, http.StatusOK, "ping tong", "")
	})

	r.POST("/api/login", users.Login)

	v1 := r.Group("/api/v1")
	v1.Use(middleware.Jwt())
	{

		v1.POST("/waiting", waiting.Create)
	}

	vt := r.Group("/api/vt")
	vt.Use(middleware.Jwt())
	{
		vt.GET("/ping", func(c *gin.Context) {
			response.Response(c.Writer, http.StatusOK, "ping tong", "")
		})
	}

	go func() {
		r.Run(":8000")
	}()
}
