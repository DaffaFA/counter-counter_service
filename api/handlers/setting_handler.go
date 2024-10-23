package handlers

import (
	"fmt"

	"github.com/DaffaFA/counter-counter_service/pkg/entities"
	"github.com/DaffaFA/counter-counter_service/pkg/setting"
	"github.com/DaffaFA/counter-counter_service/utils"
	"github.com/gofiber/fiber/v2"
)

func GetSettings(service setting.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := utils.Tracer.Start(c.UserContext(), fmt.Sprintf("%s %s", c.Method(), c.OriginalURL()))
		defer span.End()

		var queryFilter entities.FetchFilter

		if err := c.QueryParser(&queryFilter); err != nil {
			span.RecordError(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		alias := c.Params("alias", "all")

		item, err := service.FetchSetting(ctx, alias, &queryFilter)
		if err != nil {
			span.RecordError(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(item)
	}
}

// curl -X POST -H "Content-Type: application/json" -d '{"key": "key", "value": "value"}' http://localhost:3000/api/v1/setting/alias
func CreateSetting(service setting.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := utils.Tracer.Start(c.UserContext(), fmt.Sprintf("%s %s", c.Method(), c.OriginalURL()))
		defer span.End()

		var setting entities.Setting

		if err := c.BodyParser(&setting); err != nil {
			span.RecordError(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		alias := c.Params("alias")

		if err := service.CreateSetting(ctx, alias, &setting); err != nil {
			span.RecordError(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Setting created successfully",
		})
	}
}

func DeleteSetting(service setting.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := utils.Tracer.Start(c.UserContext(), fmt.Sprintf("%s %s", c.Method(), c.OriginalURL()))
		defer span.End()

		alias := c.Params("alias")
		id, err := c.ParamsInt("id")
		if err != nil {
			span.RecordError(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := service.DeleteSetting(ctx, alias, id); err != nil {
			span.RecordError(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Setting deleted successfully",
		})
	}
}
