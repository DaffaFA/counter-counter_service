package item

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
	FetchItem(context.Context, *entities.FetchFilter) (entities.ItemPagination, error)
	CreateItem(context.Context, *entities.ItemCreateParam) error
	UpdateItem(context.Context, string, *entities.ItemCreateParam) error
	FetchCountChart(context.Context, *entities.DashboardAnalyticFilter) ([]entities.ItemCountChart, error)
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

func (r *repository) FetchItem(ctx context.Context, filter *entities.FetchFilter) (entities.ItemPagination, error) {
	ctx, span := utils.Tracer.Start(ctx, "item.repository.FetchItem")
	defer span.End()

	res := entities.ItemPagination{}

	entities.SetDefaultFilter(filter)

	rawData := psql.Select(
		"i.code",
		"i.buyer_id",
		"i.style_id",
		"i.color_id",
		"i.size_id",
		"i.created_at",
		"i.updated_at",
		"row_to_json(b) as buyer",
		"row_to_json(s) as style",
		"row_to_json(c) as color",
		"row_to_json(z) as size",
	).From("counter.items i").
		LeftJoin("counter.settings b ON i.buyer_id = b.id").
		LeftJoin("counter.settings s ON i.style_id = s.id").
		LeftJoin("counter.settings c ON i.color_id = c.id").
		LeftJoin("counter.settings z ON i.size_id = z.id")

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
			squirrel.ILike{"i.code": fmt.Sprintf("%%%s%%", filter.Query)},
			squirrel.ILike{"b.value": fmt.Sprintf("%%%s%%", filter.Query)},
			squirrel.ILike{"s.value": fmt.Sprintf("%%%s%%", filter.Query)},
			squirrel.ILike{"c.value": fmt.Sprintf("%%%s%%", filter.Query)},
			squirrel.ILike{"z.value": fmt.Sprintf("%%%s%%", filter.Query)},
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

	if err := r.DB.QueryRow(ctx, sqln, args...).Scan(&res.Total, &res.Items); err != nil {
		return res, err
	}

	return res, nil
}

func (r *repository) CreateItem(ctx context.Context, item *entities.ItemCreateParam) error {
	ctx, span := utils.Tracer.Start(ctx, "item.repository.CreateItem")
	defer span.End()

	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}

	settingQuery := "INSERT INTO counter.settings (setting_type_alias, value, parent_id) VALUES ($1, $2, $3) ON CONFLICT (setting_type_alias, value, parent_id) DO UPDATE SET value = $2 RETURNING id"

	var buyerID, styleID, colorID, sizeID int

	if err := tx.QueryRow(ctx, settingQuery, "buyer", item.Buyer, 0).
		Scan(&buyerID); err != nil {
		tx.Rollback(ctx)
		return err
	}

	if err := tx.QueryRow(ctx, settingQuery, "style", item.Style, buyerID).
		Scan(&styleID); err != nil {
		tx.Rollback(ctx)
		return err
	}

	if err := tx.QueryRow(ctx, settingQuery, "color", item.Color, styleID).
		Scan(&colorID); err != nil {
		tx.Rollback(ctx)
		return err
	}

	if err := tx.QueryRow(ctx, settingQuery, "size", item.Size, styleID).
		Scan(&sizeID); err != nil {
		tx.Rollback(ctx)
		return err
	}

	query := psql.Insert("counter.items").
		Columns("code", "buyer_id", "style_id", "color_id", "size_id").
		Values(item.Code, buyerID, styleID, colorID, sizeID)

	sqln, args, err := query.ToSql()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(ctx, sqln, args...); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateItem(ctx context.Context, code string, item *entities.ItemCreateParam) error {
	ctx, span := utils.Tracer.Start(ctx, "item.repository.UpdateItem")
	defer span.End()

	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}

	settingQuery := "INSERT INTO counter.settings (setting_type_alias, value, parent_id) VALUES ($1, $2, $3) ON CONFLICT (setting_type_alias, value, parent_id) DO UPDATE SET value = $2 RETURNING id"

	var buyerID, styleID, colorID, sizeID int

	if err := tx.QueryRow(ctx, settingQuery, "buyer", item.Buyer, nil).
		Scan(&buyerID); err != nil {
		tx.Rollback(ctx)
		return err
	}

	if err := tx.QueryRow(ctx, settingQuery, "style", item.Style, buyerID).
		Scan(&styleID); err != nil {
		tx.Rollback(ctx)
		return err
	}

	if err := tx.QueryRow(ctx, settingQuery, "color", item.Color, styleID).
		Scan(&colorID); err != nil {
		tx.Rollback(ctx)
		return err
	}

	if err := tx.QueryRow(ctx, settingQuery, "size", item.Size, colorID).
		Scan(&sizeID); err != nil {
		tx.Rollback(ctx)
		return err
	}

	query := psql.Update("counter.items").
		Set("buyer_id", buyerID).
		Set("style_id", styleID).
		Set("color_id", colorID).
		Set("size_id", sizeID).
		Where(squirrel.Eq{"code": code})

	sqln, args, err := query.ToSql()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(ctx, sqln, args...); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return err
}

func (r *repository) FetchCountChart(ctx context.Context, filter *entities.DashboardAnalyticFilter) ([]entities.ItemCountChart, error) {
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
	GROUP BY bucket
	`

	rows, err := r.DB.Query(ctx, query, from, to, interval)
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
