package external

import (
	"context"
	finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
	"github.com/gofiber/fiber/v2"
)

func configClient() *finnhub.DefaultApiService {
	key := GetKey()
	cfg := finnhub.NewConfiguration()
	cfg.AddDefaultHeader("X-Finnhub-Token", key)
	finnhubClient := finnhub.NewAPIClient(cfg).DefaultApi
	return finnhubClient
}

func BuyOrSell(c *fiber.Ctx) error {

	name := c.Params("name")

	client := configClient()

	res, _, err := client.RecommendationTrends(context.Background()).Symbol(name).Execute()

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(res)
}

func GetCurrentStockPrice(c *fiber.Ctx) error {

	name := c.Params("name")

	client := configClient()

	res, _, err := client.Quote(context.Background()).Symbol(name).Execute()

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(res)
}
