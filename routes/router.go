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

	ex := u.Group("/exchange")
	ex.Use(services.AuthorizationRequired)
	ex.GET("/all", controllers.ShowAllExchange)
	ex.POST("/create", controllers.CreateExchange)
	ex.GET("/show/:id", controllers.ShowExchangeByID)
	ex.PUT("/update/:id", controllers.UpdateExchange)
	ex.DELETE("/delete/:id", controllers.DeleteExchange)

	cat := u.Group("/category")
	cat.Use(services.AuthorizationRequired)
	cat.GET("/all", controllers.ShowAllCategories)
	cat.POST("/create", controllers.CreateCategory)
	cat.GET("/show/:id", controllers.ShowCategoryByID)
	cat.PUT("/update/:id", controllers.UpdateCategory)
	cat.DELETE("/delete/:id", controllers.DeleteCategory)

	crypto := u.Group("crypto")
	crypto.GET("/all", controllers.ShowAllCryptoCurrency)
	crypto.POST("/create", controllers.CreateCryptoCurrency)
	crypto.GET("/show/:id", controllers.ShowCryptoCurrencyByID)
	crypto.PUT("/update/:id", controllers.UpdateCryptoCurrency)
	crypto.DELETE("/delete/:id", controllers.DeleteCryptoCurrency)
}
