package routes

import (
	"github.com/abe27/crypto/tracker/api/controllers"
	"github.com/abe27/crypto/tracker/api/services"
	"github.com/gin-gonic/gin"
)

func SetupRoute(r *gin.Engine) {
	r.GET("/", controllers.Hello)
	app := r.Group("/api/v1")
	app.GET("/", controllers.Hello)
	app.POST("/register", controllers.Register)
	app.POST("/login", controllers.SignIn)

	u := r.Group("/api/v1")
	u.Use(services.AuthorizationRequired)
	u.GET("/logout", controllers.SignOut)
}
