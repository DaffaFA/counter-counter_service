package handlers

import (
	"fmt"

	"github.com/DaffaFA/counter-counter_service/pkg/analytic"
	"github.com/DaffaFA/counter-counter_service/pkg/entities"
	"github.com/DaffaFA/counter-counter_service/utils"
	"github.com/gofiber/fiber/v2"
)

// func GetCountChart(service item.Service) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		ctx, span := utils.Tracer.Start(c.UserContext(), fmt.Sprintf("%s %s", c.Method(), c.OriginalURL()))
// 		defer span.End()

// 		var queryFilter entities.DashboardAnalyticFilter

// 		if err := c.QueryParser(&queryFilter); err != nil {
// 			span.RecordError(err)
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 				"error": err.Error(),
// 			})
// 		}

// 		queryFilter.SetDefault()

// 		chart, err := service.FetchCountChart(ctx, &queryFilter)
// 		if err != nil {
// 			span.RecordError(err)
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"error": err.Error(),
// 			})
// 		}

// 		return c.Status(fiber.StatusOK).JSON(chart)
// 	}
// }

func GetAnalyticItems(service analytic.Service) fiber.Handler {
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

		entities.SetDefaultFilter(&queryFilter)

		items, err := service.FetchAnalyticItems(ctx, &queryFilter)
		if err != nil {
			span.RecordError(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(items)
	}
}

func GetCountChart(service analytic.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := utils.Tracer.Start(c.UserContext(), fmt.Sprintf("%s %s", c.Method(), c.OriginalURL()))
		defer span.End()

		var queryFilter entities.DashboardAnalyticFilter

		if err := c.QueryParser(&queryFilter); err != nil {
			span.RecordError(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		queryFilter.SetDefault()

		styleId, err := c.ParamsInt("style_id", -1)
		if err != nil || styleId == -1 {
			span.RecordError(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		chart, err := service.FetchCountChart(ctx, styleId, &queryFilter)
		if err != nil {
			span.RecordError(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(chart)
	}
}

func GetAnalyticItemsByID(service analytic.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := utils.Tracer.Start(c.UserContext(), fmt.Sprintf("%s %s", c.Method(), c.OriginalURL()))
		defer span.End()

		styleId, err := c.ParamsInt("style_id", -1)
		if err != nil || styleId == -1 {
			span.RecordError(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		item, err := service.FetchAnalyticItemsByID(ctx, styleId)
		if err != nil {
			span.RecordError(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(item)
	}
}

func GetAggregateByFactory(service analytic.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := utils.Tracer.Start(c.UserContext(), fmt.Sprintf("%s %s", c.Method(), c.OriginalURL()))
		defer span.End()

		styleId, err := c.ParamsInt("style_id", -1)
		if err != nil || styleId == -1 {
			span.RecordError(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		factory, err := service.FetchAggregateByFactory(ctx, styleId)
		if err != nil {
			span.RecordError(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(factory)
	}
}
