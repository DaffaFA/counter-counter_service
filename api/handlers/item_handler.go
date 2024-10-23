package handlers

import (
	"fmt"

	"github.com/DaffaFA/counter-counter_service/pkg/entities"
	"github.com/DaffaFA/counter-counter_service/pkg/item"
	"github.com/DaffaFA/counter-counter_service/utils"
	"github.com/gofiber/fiber/v2"
)

func GetItem(service item.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := utils.Tracer.Start(c.UserContext(), fmt.Sprintf("%s %s", c.Method(), c.OriginalURL()))
		defer span.End()

		code := c.Params("code", "-1")
		if code == "-1" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "code cannot be empty",
			})
		}

		items, err := service.FetchItem(ctx, &entities.FetchFilter{Alias: code})
		if err != nil {
			span.RecordError(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if len(items.Items) < 1 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "item not found",
			})
		}

		return c.Status(fiber.StatusOK).JSON(items.Items[0])
	}
}

func GetItems(service item.Service) fiber.Handler {
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

		item, err := service.FetchItem(ctx, &queryFilter)
		if err != nil {
			span.RecordError(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(item)
	}
}

func CreateItem(service item.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := utils.Tracer.Start(c.UserContext(), fmt.Sprintf("%s %s", c.Method(), c.OriginalURL()))
		defer span.End()

		var item entities.Item

		if err := c.BodyParser(&item); err != nil {
			span.RecordError(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := service.CreateItem(ctx, &item); err != nil {
			span.RecordError(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Item created successfully",
		})
	}
}

func UpdateItem(service item.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := utils.Tracer.Start(c.UserContext(), fmt.Sprintf("%s %s", c.Method(), c.OriginalURL()))
		defer span.End()

		code := c.Params("code")

		var item entities.Item

		if err := c.BodyParser(&item); err != nil {
			span.RecordError(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := service.UpdateItem(ctx, code, &item); err != nil {
			span.RecordError(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Item updated successfully",
		})
	}
}
