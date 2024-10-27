package setting

import (
	"context"

	"github.com/DaffaFA/counter-counter_service/pkg/entities"
	"github.com/DaffaFA/counter-counter_service/utils"
)

type Service interface {
	FetchSetting(context.Context, string, *entities.FetchFilter) (entities.SettingPagination, error)
	CreateSetting(context.Context, string, *entities.Setting) error
	DeleteSetting(context.Context, string, int) error
	FetchMachineDetail(context.Context, int) (entities.MachineDetail, error)
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

func (s *service) FetchSetting(ctx context.Context, settingAlias string, filter *entities.FetchFilter) (entities.SettingPagination, error) {
	ctx, span := utils.Tracer.Start(ctx, "item.service.FetchItem")
	defer span.End()

	return s.repository.FetchSetting(ctx, settingAlias, filter)
}

func (s *service) CreateSetting(ctx context.Context, settingAlias string, setting *entities.Setting) error {
	ctx, span := utils.Tracer.Start(ctx, "item.service.CreateItem")
	defer span.End()

	return s.repository.CreateSetting(ctx, settingAlias, setting)
}

func (s *service) DeleteSetting(ctx context.Context, settingAlias string, settingID int) error {
	ctx, span := utils.Tracer.Start(ctx, "item.service.DeleteItem")
	defer span.End()

	return s.repository.DeleteSetting(ctx, settingAlias, settingID)
}

func (s *service) FetchMachineDetail(ctx context.Context, machineID int) (entities.MachineDetail, error) {
	ctx, span := utils.Tracer.Start(ctx, "item.service.FetchMachineDetail")
	defer span.End()

	return s.repository.FetchMachineDetail(ctx, machineID)
}
