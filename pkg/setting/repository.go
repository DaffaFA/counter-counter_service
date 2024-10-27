package setting

import (
	"context"
	"fmt"

	"github.com/DaffaFA/counter-counter_service/pkg/entities"
	"github.com/DaffaFA/counter-counter_service/utils"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

type Repository interface {
	FetchSetting(context.Context, string, *entities.FetchFilter) (entities.SettingPagination, error)
	CreateSetting(context.Context, string, *entities.Setting) error
	DeleteSetting(context.Context, string, int) error
	FetchMachineDetail(context.Context, int) (entities.MachineDetail, error)
}

type repository struct {
	DB *pgxpool.Pool
}

// NewRepo is the single instance repo that is being created.
func NewRepo(db *pgxpool.Pool) Repository {
	return &repository{
		DB: db,
	}
}

func (r *repository) FetchSetting(ctx context.Context, settingAlias string, filter *entities.FetchFilter) (entities.SettingPagination, error) {
	ctx, span := utils.Tracer.Start(ctx, "setting.repository.FetchSetting")
	defer span.End()

	res := entities.SettingPagination{}

	entities.SetDefaultFilter(filter)

	rawData := psql.Select(
		"id",
		"value",
		"setting_type_alias",
		"parent_id",
	).From("counter.settings")

	if settingAlias != "all" {
		rawData = rawData.Where(squirrel.Eq{"setting_type_alias": settingAlias})
	}

	if filter.Query != "" {
		rawData = rawData.Where(squirrel.Or{
			squirrel.ILike{
				"setting_type_alias": fmt.Sprintf("%%%s%%", filter.Query),
				"value":              fmt.Sprintf("%%%s%%", filter.Query),
			},
		})
	}

	rawData = rawData.Prefix("WITH raw_data AS (").
		Suffix(")")

	pagination := psql.Select("*").
		From("raw_data").
		Limit(filter.Limit).
		Offset(filter.Cursor*filter.Limit - filter.Limit).
		PrefixExpr(rawData).Prefix(", with_pagination AS (").Suffix(")")

	query := psql.Select("(SELECT COUNT(*) FROM raw_data) as total",
		"(SELECT json_agg(with_pagination) FROM with_pagination) as data").PrefixExpr(pagination)

	sqln, args, err := query.ToSql()
	if err != nil {
		return res, err
	}

	if err := r.DB.QueryRow(ctx, sqln, args...).Scan(&res.Total, &res.Settings); err != nil {
		return res, err
	}

	return res, nil
}

func (r *repository) CreateSetting(ctx context.Context, settingAlias string, setting *entities.Setting) error {
	ctx, span := utils.Tracer.Start(ctx, "setting.repository.CreateSetting")
	defer span.End()

	query := psql.Insert("counter.settings").
		Columns("value", "setting_type_alias", "parent_id").
		Values(setting.Value, settingAlias, setting.ParentID)

	sqln, args, err := query.ToSql()
	if err != nil {
		return err
	}

	if _, err := r.DB.Exec(ctx, sqln, args...); err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteSetting(ctx context.Context, settingAlias string, settingID int) error {
	ctx, span := utils.Tracer.Start(ctx, "setting.repository.DeleteSetting")
	defer span.End()

	query := psql.Delete("counter.settings").
		Where(squirrel.Eq{"id": settingID, "setting_type_alias": settingAlias})

	sqln, args, err := query.ToSql()
	if err != nil {
		return err
	}

	if _, err := r.DB.Exec(ctx, sqln, args...); err != nil {
		return err
	}

	return nil
}

func (r *repository) FetchMachineDetail(ctx context.Context, machineID int) (entities.MachineDetail, error) {
	ctx, span := utils.Tracer.Start(ctx, "setting.repository.FetchMachineDetail")
	defer span.End()

	res := entities.MachineDetail{}

	query := psql.Select("f.value", "m.value").
		From("counter.settings m").
		LeftJoin("counter.settings f ON m.parent_id = f.id").
		Where(squirrel.Eq{"m.id": machineID})

	sqln, args, err := query.ToSql()
	if err != nil {
		return res, err
	}

	if err := r.DB.QueryRow(ctx, sqln, args...).Scan(&res.Factory, &res.Machine); err != nil {
		return res, err
	}

	return res, nil
}
