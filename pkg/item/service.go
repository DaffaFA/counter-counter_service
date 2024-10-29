package item

import (
	"context"

	"github.com/DaffaFA/counter-counter_service/pkg/entities"
	"github.com/DaffaFA/counter-counter_service/utils"
)

type Service interface {
	FetchItem(context.Context, *entities.FetchFilter) (entities.ItemPagination, error)
	CreateItem(context.Context, *entities.ItemCreateParam) error
	UpdateItem(context.Context, string, *entities.ItemCreateParam) error
	FetchCountChart(context.Context, *entities.DashboardAnalyticFilter) ([]entities.ItemCountChart, error)
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

func (s *service) FetchItem(ctx context.Context, filter *entities.FetchFilter) (entities.ItemPagination, error) {
	ctx, span := utils.Tracer.Start(ctx, "item.service.FetchItem")
	defer span.End()

	return s.repository.FetchItem(ctx, filter)
}

func (s *service) CreateItem(ctx context.Context, item *entities.ItemCreateParam) error {
	ctx, span := utils.Tracer.Start(ctx, "item.service.CreateItem")
	defer span.End()

	return s.repository.CreateItem(ctx, item)
}

func (s *service) UpdateItem(ctx context.Context, code string, item *entities.ItemCreateParam) error {
	ctx, span := utils.Tracer.Start(ctx, "item.service.UpdateItem")
	defer span.End()

	return s.repository.UpdateItem(ctx, code, item)
}

func (s *service) FetchCountChart(ctx context.Context, filter *entities.DashboardAnalyticFilter) ([]entities.ItemCountChart, error) {
	ctx, span := utils.Tracer.Start(ctx, "item.service.FetchCountChart")
	defer span.End()

	return s.repository.FetchCountChart(ctx, filter)
}
