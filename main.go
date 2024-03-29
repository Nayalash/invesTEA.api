package main

import (
	"GoProject/database"
	"GoProject/external"
	"GoProject/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my API")
}

func setupRoutes(app *fiber.App) {
	// Welcome
	app.Get("/api", welcome)
	//User
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("api/users/:id", routes.GetUser)
	app.Put("api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)

	//Favourite
	app.Post("/api/favourite", routes.CreateFavourite)
	app.Get("/api/favourite/", routes.GetFavourites)

	//Finhub Integration
	app.Get("/api/stocks/price/:name", external.GetCurrentStockPrice)
	app.Get("/api/stocks/forecast/:name", external.BuyOrSell)
}

//Need to implment company routes left : video-> 51:24

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
