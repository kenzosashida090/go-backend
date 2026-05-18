package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	authRoutes := server.Group("/")
	authRoutes.Use(validateAuth) // Protected Routes

	authRoutes.GET("/events/:id", getEvent)
	authRoutes.GET("/events", getEvents)
	authRoutes.POST("/events", postEvent)
	authRoutes.PUT("/events/:id", updateEvent)
	authRoutes.DELETE("/events/:id", deleteEvent)

	/////////////////////////////////////////

	server.POST("/users", signUp)
	server.POST("/login", login)
	authRoutes.POST("/logout", logout)

}
