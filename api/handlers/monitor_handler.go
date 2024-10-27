package handlers

import (
	"fmt"

	"github.com/DaffaFA/counter-counter_service/pkg/entities"
	"github.com/DaffaFA/counter-counter_service/pkg/item_scan"
	"github.com/DaffaFA/counter-counter_service/pkg/setting"
	"github.com/DaffaFA/counter-counter_service/utils"
	"github.com/gofiber/fiber/v2"
)

func GetMachineDetail(service setting.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := utils.Tracer.Start(c.UserContext(), fmt.Sprintf("%s %s", c.Method(), c.OriginalURL()))
		defer span.End()

		id, err := c.ParamsInt("id", -1)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		machine, err := service.FetchMachineDetail(ctx, id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(machine)
	}
}

func GetMachineLatestScan(service item_scan.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := utils.Tracer.Start(c.UserContext(), fmt.Sprintf("%s %s", c.Method(), c.OriginalURL()))
		defer span.End()

		id, err := c.ParamsInt("id", -1)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		scans, err := service.FetchLatestScan(ctx, id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(scans)
	}
}

func ScanItem(service item_scan.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := utils.Tracer.Start(c.UserContext(), fmt.Sprintf("%s %s", c.Method(), c.OriginalURL()))
		defer span.End()

		id, err := c.ParamsInt("id", -1)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		var sip entities.ScanItemParam
		if err := c.BodyParser(&sip); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		scannedItem, err := service.ScanItem(ctx, id, sip.Code)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(scannedItem)
	}
}

// ResetScanCounter(context.Context, string) error
func ResetScanCounter(service item_scan.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := utils.Tracer.Start(c.UserContext(), fmt.Sprintf("%s %s", c.Method(), c.OriginalURL()))
		defer span.End()

		code := c.Params("code", "-1")
		if code == "-1" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "code is required",
			})
		}

		if err := service.ResetScanCounter(ctx, code); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Scan counter reset successfully",
		})
	}
}
