package routes

import (
	"github.com/DaffaFA/counter-counter_service/api/handlers"
	"github.com/DaffaFA/counter-counter_service/pkg/analytic"
	"github.com/gofiber/fiber/v2"
)

func AnalyticRouter(app fiber.Router, service analytic.Service) {
	route := app.Group("/analytic")

	route.Get("/", handlers.GetAnalyticItems(service))
	route.Get("/:style_id/chart", handlers.GetCountChart(service))
}
