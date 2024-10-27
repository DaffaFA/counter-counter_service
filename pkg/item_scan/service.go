package item_scan

import (
	"context"

	"github.com/DaffaFA/counter-counter_service/pkg/entities"
	"github.com/DaffaFA/counter-counter_service/utils"
)

type Service interface {
	FetchLatestScan(context.Context, int) ([]entities.LatestScan, error)
	ScanItem(context.Context, int, string) (entities.ScannedItem, error)
	UndoLastCounter(context.Context, string, string) error
}

type service struct {
	repository Repository
}

// NewRepo is the single instance repo that is being created.
func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) FetchLatestScan(ctx context.Context, machineID int) ([]entities.LatestScan, error) {
	ctx, span := utils.Tracer.Start(ctx, "item_scan.service.FetchLatestScan")
	defer span.End()

	return s.repository.FetchLatestScan(ctx, machineID)
}

func (s *service) ScanItem(ctx context.Context, machineID int, qrCode string) (entities.ScannedItem, error) {
	ctx, span := utils.Tracer.Start(ctx, "item_scan.service.ScanItem")
	defer span.End()

	return s.repository.ScanItem(ctx, machineID, qrCode)
}

func (s *service) UndoLastCounter(ctx context.Context, time string, qrCode string) error {
	ctx, span := utils.Tracer.Start(ctx, "item_scan.service.UndoLastCounter")
	defer span.End()

	return s.repository.UndoLastCounter(ctx, time, qrCode)
}
