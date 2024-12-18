package routes

import (
	"github.com/DaffaFA/counter-counter_service/api/handlers"
	"github.com/DaffaFA/counter-counter_service/pkg/setting"
	"github.com/gofiber/fiber/v2"
)

func SettingRoutes(app fiber.Router, service setting.Service) {
	routes := app.Group("/setting")

	routes.Get("/:alias", handlers.GetSettings(service))
	routes.Post("/:alias", handlers.CreateSetting(service))
	routes.Get("/:alias/:id", handlers.GetSettings(service))

	app.Get("/monitor/machine/:id", handlers.GetMachineDetail(service))
}
