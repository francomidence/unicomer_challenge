package router

import (
	"github.com/gin-gonic/gin"
	"time"
	"unicomer_challenge/server/controllers"
	"unicomer_challenge/server/middlewares"
	"unicomer_challenge/services/holidays"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	// Initialize the rate limiter
	rateLimiter := middlewares.NewRateLimiter(1, time.Second*5, 1)

	// Apply middleware for CORS, request ID generation, and rate limiting
	router.Use(
		middlewares.CORSMiddleware(),
		middlewares.RequestIDMiddleware(),
		middlewares.RateLimiterMiddleware(rateLimiter),
	)

	holidaysService := holidays.NewHolidaysService()
	holidaysController := controllers.NewHolidaysController(holidaysService)

	api := router.Group("api")
	{
		v1 := api.Group("v1")
		{
			holidaysRouter := v1.Group("/holidays")
			{
				holidaysRouter.GET("/", holidaysController.GetHolidays)
			}
		}
	}

	return router

}
