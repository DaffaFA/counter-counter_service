package analytic

import (
	"context"
	"fmt"
	"log"
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
		"s.value as style",
		"b.value as buyer",
		"count(si.time) as amount",
	).
		From("counter.settings s").
		LeftJoin("counter.settings b ON s.parent_id = b.id").
		LeftJoin("counter.item_scans si ON si.qr_code_code IN (SELECT code FROM counter.items WHERE style_id = s.id)").
		Where(squirrel.Eq{"s.setting_type_alias": "style"}).
		GroupBy("s.id", "b.value", "s.value")

	if len(filter.Sort) > 0 {
		for _, sort := range filter.Sort {
			if sort[0] == '-' {
				rawData = rawData.OrderBy(fmt.Sprintf("%s DESC", sort[1:]))
			} else {
				rawData = rawData.OrderBy(fmt.Sprintf("%s ASC", sort))
			}
		}
	}

	if filter.Alias != "" {
		rawData = rawData.Where(squirrel.Eq{"i.code": filter.Alias})
	}

	if filter.Alias == "" && filter.Query != "" {
		rawData = rawData.Where(squirrel.Or{
			squirrel.ILike{"b.value": fmt.Sprintf("%%%s%%", filter.Query)},
			squirrel.ILike{"s.value": fmt.Sprintf("%%%s%%", filter.Query)},
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

	var res entities.AnalyticItemPagination

	sqln, args, err := query.ToSql()
	if err != nil {
		return res, err
	}

	log.Println(sqln, args)

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
