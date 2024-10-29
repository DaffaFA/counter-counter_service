package analytic

import (
	"context"

	"github.com/DaffaFA/counter-counter_service/pkg/entities"
	"github.com/DaffaFA/counter-counter_service/utils"
)

type Service interface {
	FetchAnalyticItems(context.Context, *entities.FetchFilter) (entities.AnalyticItemPagination, error)
	FetchCountChart(context.Context, int, *entities.DashboardAnalyticFilter) ([]entities.ItemCountChart, error)
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
