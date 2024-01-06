package routes

import (
	"example.com/api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authentication)

	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", UpdateEvent)
	authenticated.DELETE("/events/:id", DeleteEvent)
	authenticated.POST("/events/:id/register", RegisterForEvent)
	authenticated.DELETE("/events/:id/register", CancelRegisteration)

	server.POST("/signup", SignUp)
	server.POST("/login", Login)
}
