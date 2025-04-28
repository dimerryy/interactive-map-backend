package routes

import (
	"interactive-map/controllers"
	"interactive-map/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	r.GET("/countries", middleware.AuthMiddleware(), controllers.GetCountries)
	r.POST("/countries", middleware.AuthMiddleware(), controllers.UpdateCountryStatus)
	r.DELETE("/countries/:countryISO", middleware.AuthMiddleware(), controllers.DeleteCountry)
	r.POST("/ai", controllers.HandleAI)

}
