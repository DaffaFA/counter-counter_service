package analytic

import (
	"context"
	"errors"

	"github.com/DaffaFA/counter-counter_service/pkg/entities"
	"github.com/DaffaFA/counter-counter_service/utils"
)

type Service interface {
	FetchAnalyticItems(context.Context, *entities.FetchFilter) (entities.AnalyticItemPagination, error)
	FetchCountChart(context.Context, int, *entities.DashboardAnalyticFilter) ([]entities.ItemCountChart, error)
	FetchAnalyticItemsByID(context.Context, int) (entities.AnalyticItem, error)
	FetchAggregateByFactory(context.Context, int) ([]entities.AggregateByFactory, error)
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

func (s *service) FetchAnalyticItems(ctx context.Context, filter *entities.FetchFilter) (entities.AnalyticItemPagination, error) {
	ctx, span := utils.Tracer.Start(ctx, "analytic.service.FetchAnalyticItems")
	defer span.End()

	return s.repository.FetchAnalyticItems(ctx, filter)
}

func (s *service) FetchCountChart(ctx context.Context, styleId int, filter *entities.DashboardAnalyticFilter) ([]entities.ItemCountChart, error) {
	ctx, span := utils.Tracer.Start(ctx, "analytic.service.FetchCountChart")
	defer span.End()

	return s.repository.FetchCountChart(ctx, styleId, filter)
}

func (s *service) FetchAnalyticItemsByID(ctx context.Context, styleId int) (entities.AnalyticItem, error) {
	ctx, span := utils.Tracer.Start(ctx, "analytic.service.FetchAnalyticItemsByID")
	defer span.End()

	items, err := s.repository.FetchAnalyticItems(ctx, &entities.FetchFilter{
		ID:     styleId,
		Limit:  1,
		Cursor: 1,
	})
	if err != nil {
		return entities.AnalyticItem{}, err
	}

	if len(items.Items) == 0 {
		return entities.AnalyticItem{}, errors.New("item not found")
	}

	return items.Items[0], nil
}

func (s *service) FetchAggregateByFactory(ctx context.Context, styleId int) ([]entities.AggregateByFactory, error) {
	ctx, span := utils.Tracer.Start(ctx, "analytic.service.FetchAggregateByFactory")
	defer span.End()

	return s.repository.FetchAggregateByFactory(ctx, styleId)
}
