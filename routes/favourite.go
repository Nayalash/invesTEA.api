package routes

import (
	"GoProject/database"
	"GoProject/models"
	"github.com/gofiber/fiber/v2"
)

type FavouriteSerializer struct {
	ID     uint           `json:"id"`
	Ticker string         `json:"ticker"`
	User   UserSerializer `json:"user"`
}

func CreateResponseFavourite(favourite models.Favourite, user UserSerializer) FavouriteSerializer {
	return FavouriteSerializer{ID: favourite.ID, Ticker: favourite.Ticker, User: user}
}

func CreateFavourite(c *fiber.Ctx) error {
	var favourite models.Favourite

	if err := c.BodyParser(&favourite); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User

	if err := findUser(favourite.UserRefer, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&favourite)

	responseUser := CreateResponseUser(user)
	responseOrder := CreateResponseFavourite(favourite, responseUser)

	return c.Status(200).JSON(responseOrder)
}

func GetFavourites(c *fiber.Ctx) error {

	favs := []models.Favourite{}
	database.Database.Db.Find(&favs)
	responseFavs := []FavouriteSerializer{}

	for _, f := range favs {
		var user models.User
		database.Database.Db.Find(&user, "id = ?", f.UserRefer)
		responseFav := CreateResponseFavourite(f, CreateResponseUser(user))
		responseFavs = append(responseFavs, responseFav)
	}

	return c.Status(200).JSON(responseFavs)
}
