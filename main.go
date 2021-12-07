package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"

	language "cloud.google.com/go/language/apiv1"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
)

func main() {
	app := fiber.New()

	app.Static("/", "./index.html")

	app.Post("/sentiment", func(c *fiber.Ctx) error {
		fmt.Println("Processing POST on /sentiment")
		input := c.FormValue("sentimentInput")
		if input == "" {
			c.SendString("Empty input")
			return c.SendStatus(400)
		}
		client, _ := language.NewClient(c.Context())
		result, err := analyzeSentiment(c.Context(), client, input)

		client.Close()
		if err != nil {
			fmt.Println(err.Error())
			return c.SendStatus(500)
		}

		fmt.Println("Finished processing.")
		return c.SendString(result.String())
	})

	app.Listen(":3000")
}

func analyzeSentiment(ctx context.Context, client *language.Client, text string) (*languagepb.AnalyzeSentimentResponse, error) {
	return client.AnalyzeSentiment(ctx, &languagepb.AnalyzeSentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
	})
}
