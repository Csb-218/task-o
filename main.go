package main

import (
	"errors"
	"fmt"
	"todo/app/routes"

	"github.com/gofiber/fiber/v2"
)

func mightFail(input int) (string, error) {
	if input < 0 {
		// Return an error value if something goes wrong
		return "", errors.New("input cannot be negative")
	}
	// Return a result and nil error if successful
	return "Success!", nil
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		result, err := mightFail(1)
		if err != nil {
			// Explicitly check and handle the error
			fmt.Printf("Error occurred: %v\n", err)
			return err
		}
		fmt.Println(result)
		return c.SendString("Hello, World!")
	})

	routes.TodoRoutes(app)

	app.Listen(":3000")

}
