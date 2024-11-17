package analytic

import (
	"context"
	"fmt"
	"time"

	"github.com/DaffaFA/counter-counter_service/pkg/entities"
	"github.com/DaffaFA/counter-counter_service/utils"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

type Repository interface {
	FetchAnalyticItems(context.Context, *entities.FetchFilter) (entities.AnalyticItemPagination, error)
	FetchCountChart(context.Context, int, *entities.DashboardAnalyticFilter) ([]entities.ItemCountChart, error)
	FetchAggregateByFactory(context.Context, int) ([]entities.AggregateByFactory, error)
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

func (r *repository) FetchAnalyticItems(ctx context.Context, filter *entities.FetchFilter) (entities.AnalyticItemPagination, error) {
	ctx, span := utils.Tracer.Start(ctx, "analytic.repository.FetchAnalyticItems")
	defer span.End()

	rawData := psql.Select(
		"s.id as id",
		"s.name as style",
		"b.value as buyer",
		"s.amount as amount",
	).
		From("counter.styles s").
		LeftJoin("counter.settings b ON s.buyer_id = b.id")

	if len(filter.Sort) > 0 {
		for _, sort := range filter.Sort {
			if sort[0] == '-' {
				rawData = rawData.OrderBy(fmt.Sprintf("%s DESC", sort[1:]))
			} else {
				rawData = rawData.OrderBy(fmt.Sprintf("%s ASC", sort))
			}
		}
	}

	if filter.ID != 0 {
		rawData = rawData.Where(squirrel.Eq{"s.id": filter.ID})
	}

	if filter.ID == 0 && filter.Query != "" {
		rawData = rawData.Where(squirrel.Or{
			squirrel.ILike{"b.value": fmt.Sprintf("%%%s%%", filter.Query)},
			squirrel.ILike{"s.name": fmt.Sprintf("%%%s%%", filter.Query)},
		})
	}

	rawData = rawData.Prefix("WITH raw_data AS (").
		Suffix(")")

	pagination := psql.Select("*").
		From("raw_data").
		Limit(filter.Limit).
		Offset(filter.Cursor * filter.Limit).
		PrefixExpr(rawData).Prefix(", with_pagination AS (").Suffix(")")

	query := psql.Select("(SELECT COUNT(*) FROM raw_data) as total",
		"(SELECT json_agg(with_pagination) FROM with_pagination) as data").PrefixExpr(pagination)

	var res entities.AnalyticItemPagination

	sqln, args, err := query.ToSql()
	if err != nil {
		return res, err
	}

	if err := r.DB.QueryRow(ctx, sqln, args...).Scan(&res.Total, &res.Items); err != nil {
		return res, err
	}

	return res, nil
}

func (r *repository) FetchCountChart(ctx context.Context, styleId int, filter *entities.DashboardAnalyticFilter) ([]entities.ItemCountChart, error) {
	from, err := time.Parse(time.RFC3339, filter.From)
	if err != nil {
		return nil, err
	}

	to, err := time.Parse(time.RFC3339, filter.To)
	if err != nil {
		return nil, err
	}

	ctx, span := utils.Tracer.Start(ctx, "item.repository.FetchCountChart")
	defer span.End()

	time := to.Sub(from)

	interval := "1 hour"

	if time.Hours() > 24 {
		interval = "1 day"
	} else if time.Hours() > 168 {
		interval = "1 week"
	} else if time.Hours() > 720 {
		interval = "1 month"
	}

	query := `
		SELECT time_bucket_gapfill($3, time) as bucket, COUNT(time) as data
		FROM counter.item_scans
		WHERE time BETWEEN $1 AND $2
		AND qr_code_code IN (SELECT code FROM counter.items WHERE style_id = $4)
		GROUP BY bucket
	`

	rows, err := r.DB.Query(ctx, query, from, to, interval, styleId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []entities.ItemCountChart

	for rows.Next() {
		var count entities.ItemCountChart
		if err := rows.Scan(&count.Bucket, &count.Count); err != nil {
			return nil, err
		}

		res = append(res, count)
	}

	return res, nil
}

func (r *repository) FetchAggregateByFactory(ctx context.Context, factoryId int) ([]entities.AggregateByFactory, error) {
	query := `
		WITH raw_data AS (SELECT f.value as factory, i.code, c.value as color, z.value as size, count(si.time) as total
											FROM counter.settings f
															LEFT JOIN counter.settings m ON m.parent_id = f.id
															LEFT JOIN counter.item_scans si ON si.machine_id = m.id
															LEFT JOIN counter.items i ON si.qr_code_code = i.code
															LEFT JOIN counter.settings b ON b.id = i.buyer_id
															LEFT JOIN counter.settings c ON c.id = i.color_id
															LEFT JOIN counter.settings z ON z.id = i.size_id
											WHERE f.setting_type_alias = 'factory'
											AND i.style_id = $1
											GROUP BY f.value, i.code, c.value, z.value),
				grouped AS (SELECT factory,
														jsonb_agg(json_build_object(
																		'code', code,
																		'color', color,
																		'size', size,
																		'count', total
																			)) as rows,
														sum(total)   as total
										FROM raw_data
										GROUP BY factory)
		SELECT *
		FROM grouped
	`

	var res []entities.AggregateByFactory

	rows, err := r.DB.Query(ctx, query, factoryId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var agg entities.AggregateByFactory
		if err := rows.Scan(&agg.Factory, &agg.Rows, &agg.Total); err != nil {
			return nil, err
		}

		res = append(res, agg)
	}

	return res, nil

}
