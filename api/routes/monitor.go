package routes

import (
	"github.com/DaffaFA/counter-counter_service/api/handlers"
	"github.com/DaffaFA/counter-counter_service/pkg/item_scan"
	"github.com/gofiber/fiber/v2"
)

func MonitorRoutes(app fiber.Router, service item_scan.Service) {
	routes := app.Group("/monitor")

	routes.Get("/machine/:id/latest-scan", handlers.GetMachineLatestScan(service))
	routes.Post("/machine/:id", handlers.ScanItem(service))
	routes.Post("/item/:code/reset", handlers.UndoLastCounter(service))
}
