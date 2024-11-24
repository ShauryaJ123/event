package routes

import (
	"abc.com/calc/middlewares"
	"github.com/gin-gonic/gin"
)



func RegisterRoutes(server *gin.Engine) {
    // Public routes
    server.GET("/events", getEvents)               // Fetch all events
    server.POST("/signup", signup)                 // User signup
    server.POST("/login", login)    
	server.GET("/events/:name", getEventByName)               // User login
	


    // Authenticated routes
    authenticated := server.Group("/")
    authenticated.Use(middlewares.Authenticate)

    // Fetch event by name
    authenticated.POST("/events/register/:event_id", registerEvent) // Register for an event
    authenticated.GET("/events/my-registrations", viewRegisteredEvents) // View user's registered events
    // authenticated.DELETE("/events/:id/register", cancelRegistration)    // Cancel registration for an event
}
