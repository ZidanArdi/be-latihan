package main

import (
	"log"

	"be_latihan/config"
	"be_latihan/docs"
	"be_latihan/model"
	"be_latihan/router"

	"github.com/gofiber/fiber/v2"
	//"gorm.io/gorm/logger"
	"strings"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	//"github.com/gofiber/swagger"
	"os"
	//"be_latihan/docs"
	//"be_latihan/routes"
)

// @title API Praktikum 13 - be_latihan
// @version 1.0
// @description Dokumentasi API backend be_latihan menggunakan Golang Fiber, GORM, PostgreSQL, dan JWT.
// @contact.name Praktikum Pemrograman III
// @contact.email praktikum@example.com
// @host 127.0.0.1:3000
// @BasePath /
// @schemes http https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	app := fiber.New()
	app.Use(logger.New())

	//Swagger Host Configuration
	swaggerHost := os.Getenv("SWAGGER_HOST")
	if swaggerHost == "" {
		swaggerHost = "127.0.0.1:3000"
	}
	docs.SwaggerInfo.Host = swaggerHost


	app.Use(cors.New(cors.Config{
    AllowOrigins: strings.Join(config.GetAllowedOrigins(), ","),
    AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
    AllowHeaders: "Origin,Content-Type,Accept,Authorization",
}))

	config.InitDB()
	config.GetDB().AutoMigrate(&model.Mahasiswa{}, &model.User{})
	router.SetupRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}