package routes

import (
	"github.com/DaffaFA/counter-counter_service/api/handlers"
	"github.com/DaffaFA/counter-counter_service/pkg/item"
	"github.com/gofiber/fiber/v2"
)

func ItemRouter(app fiber.Router, service item.Service) {
	item := app.Group("/item")

	item.Get("/", handlers.GetItems(service))
	item.Post("/", handlers.CreateItem(service))
	item.Get("/:code", handlers.GetItem(service))
	item.Put("/:code", handlers.UpdateItem(service))
}
