package routes

import (
	"todo/app/dal"

	"github.com/gofiber/fiber/v2"
)

func TodoRoutes(app *fiber.App) {
	r := app.Group("/todo")

	r.Post("/", func(c *fiber.Ctx) error {
		t := dal.Todo{}
		if err := c.BodyParser(&t); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		created, err := dal.CreateTodo(t)
		if err != nil {
			return c.Status(400).JSON(err.Error())
		}
		return c.Status(201).JSON(created)
	})

	r.Get("/", func(c *fiber.Ctx) error {
		t, err := dal.ReadTodo()
		if err != nil {
			return c.Status(404).JSON(err.Error())
		}
		return c.Status(200).JSON(t)
	})

	r.Get("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(400).JSON("Invalid ID")
		}
		t, err := dal.ReadTodoById(id)
		if err != nil {
			return c.Status(404).JSON(err.Error())
		}
		return c.Status(200).JSON(t)
	})

	r.Put("/", func(c *fiber.Ctx) error {
		t := dal.Todo{}
		if err := c.BodyParser(&t); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		updated, err := dal.UpdateTodo(t)
		if err != nil {
			return c.Status(404).JSON(err.Error())
		}
		return c.Status(200).JSON(updated)
	})

	r.Delete("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(400).JSON("Invalid ID")
		}
		if err := dal.DeleteTodo(id); err != nil {
			return c.Status(404).JSON(err.Error())
		}
		return c.Status(200).JSON("Deleted")
	})
}
