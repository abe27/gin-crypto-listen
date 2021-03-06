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
	"gorm.io/driver/postgres"
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

	dns := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=" + os.Getenv("SSL_MODE") +
		" TimeZone=" + os.Getenv("TZ_NAME") + ""

	if os.Getenv("DB_PASSWORD") == "" {
		dns = "host=" + os.Getenv("DB_HOST") +
			" user=" + os.Getenv("DB_USER") +
			" dbname=" + os.Getenv("DB_NAME") +
			" port=" + os.Getenv("DB_PORT") +
			" sslmode=" + os.Getenv("SSL_MODE") +
			" TimeZone=" + os.Getenv("TZ_NAME") + ""
	}
	services.DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{
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
	services.DB.AutoMigrate(&models.ApiData{})
}

func main() {
	app := gin.Default()
	app.Use(cors.Default())
	routes.SetupRoute(app)
	app.Run(":3000")
}
