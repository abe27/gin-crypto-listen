package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/abe27/crypto/tracker/api/models"
	"github.com/abe27/crypto/tracker/api/routes"
	"github.com/abe27/crypto/tracker/api/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func init() {
	fmt.Println("Starting server....")
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	dsn := "" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	services.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tbt_", // table name prefix, table for `User` would be `t_users`
			SingularTable: false,  // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   false,  // skip the snake_casing of names
			NameReplacer:  strings.NewReplacer("CID", "Cid"),
		},
	})

	if err != nil {
		panic("Failed to connect to database")
	}

	services.DB.AutoMigrate(&models.User{})
	services.DB.AutoMigrate(&models.JwtToken{})
	services.DB.AutoMigrate(&models.Exchange{})
	services.DB.AutoMigrate(&models.Category{})
	services.DB.AutoMigrate(&models.Cryptocurrency{})
	services.DB.AutoMigrate(&models.Asset{})
	services.DB.AutoMigrate(&models.HistoricalData{})
	services.DB.AutoMigrate(&models.Interesting{})
	services.DB.AutoMigrate(&models.Investment{})
}

func main() {
	app := gin.Default()
	app.Use(cors.Default())
	routes.SetupRoute(app)
	app.Run(":3000")
}
