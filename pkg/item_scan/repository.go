package item_scan

import (
	"context"
	"errors"

	"github.com/DaffaFA/counter-counter_service/pkg/entities"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

type Repository interface {
	FetchLatestScan(context.Context, int) ([]entities.LatestScan, error)
	ScanItem(context.Context, int, string) (entities.ScannedItem, error)
	UndoLastCounter(context.Context, string, string) (entities.ScannedItem, error)
}

type repository struct {
	DB *pgxpool.Pool
}

func NewRepo(DB *pgxpool.Pool) Repository {
	return &repository{
		DB: DB,
	}
}

func (r *repository) FetchLatestScan(ctx context.Context, machineID int) ([]entities.LatestScan, error) {
	query := psql.Select("isc.time", "i.code", "b.value as buyer", "s.value as style", "z.value as size", "c.value as color").
		From("counter.item_scans isc").
		LeftJoin("counter.items i ON isc.qr_code_code = i.code").
		LeftJoin("counter.settings b ON i.buyer_id = b.id").
		LeftJoin("counter.settings s ON i.style_id = s.id").
		LeftJoin("counter.settings z ON i.size_id = z.id").
		LeftJoin("counter.settings c ON i.color_id = c.id").
		Where(squirrel.Eq{"isc.machine_id": machineID}).
		OrderBy("isc.time DESC").
		Limit(15)

	sqln, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var scans []entities.LatestScan

	rows, err := r.DB.Query(ctx, sqln, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var scan entities.LatestScan
		err := rows.Scan(&scan.Time, &scan.QrCode, &scan.Buyer, &scan.Style, &scan.Size, &scan.Color)
		if err != nil {
			return nil, err
		}
		scans = append(scans, scan)
	}

	return scans, nil
}

func (r *repository) ScanItem(ctx context.Context, machineId int, code string) (entities.ScannedItem, error) {
	tx, err := r.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		tx.Rollback(ctx)
		return entities.ScannedItem{}, err
	}
	defer tx.Rollback(ctx)

	// check if qr code exists
	var qrCodeExists bool
	if err := tx.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM counter.items WHERE code = $1)", code).Scan(&qrCodeExists); err != nil {
		tx.Rollback(ctx)
		return entities.ScannedItem{}, err
	}

	if !qrCodeExists {
		tx.Rollback(ctx)
		return entities.ScannedItem{}, errors.New("QR Code not found")
	}

	insertItemScan := psql.Insert("counter.item_scans").
		Columns("time", "machine_id", "qr_code_code").
		Values("NOW()", machineId, code).Prefix("WITH scanned_item AS (").Suffix("RETURNING *)")

	data := psql.
		Select("si.time", "i.code", "b.value as buyer", "s.value as style", "c.value as color", "z.value as size", "(SELECT COUNT(time) + 1 FROM counter.item_scans is2 WHERE is2.qr_code_code = $3) AS count").
		From("scanned_item si").
		LeftJoin("counter.items i ON si.qr_code_code = i.code").
		LeftJoin("counter.settings b ON i.buyer_id = b.id").
		LeftJoin("counter.settings s ON i.style_id = s.id").
		LeftJoin("counter.settings c ON i.color_id = c.id").
		LeftJoin("counter.settings z ON i.size_id = z.id").
		PrefixExpr(insertItemScan)

	var scannedItem entities.ScannedItem

	sqln, args, err := data.ToSql()
	if err != nil {
		tx.Rollback(ctx)
		return scannedItem, err
	}

	if err := tx.QueryRow(ctx, sqln, args...).Scan(&scannedItem.Time, &scannedItem.QrCode, &scannedItem.Buyer, &scannedItem.Style, &scannedItem.Color, &scannedItem.Size, &scannedItem.Count); err != nil {
		tx.Rollback(ctx)
		return scannedItem, err
	}

	tx.Commit(ctx)

	return scannedItem, nil
}

func (r *repository) UndoLastCounter(ctx context.Context, time string, code string) (entities.ScannedItem, error) {
	var scannedItem entities.ScannedItem

	tx, err := r.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return scannedItem, err
	}

	var rowExists bool
	if err := tx.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM counter.item_scans WHERE qr_code_code = $1 AND time = $2)", code, time).Scan(&rowExists); err != nil {
		tx.Rollback(ctx)
		return entities.ScannedItem{}, err
	}

	if !rowExists {
		tx.Rollback(ctx)
		return entities.ScannedItem{}, errors.New("the item already undone")
	}

	deleteItemScan := psql.Delete("counter.item_scans").Where(squirrel.And{
		squirrel.Eq{"qr_code_code": code},
		squirrel.Eq{"time": time},
	}).Prefix("WITH scanned_item AS (").Suffix("RETURNING *)")

	data := psql.
		Select("si.time", "i.code", "b.value as buyer", "s.value as style", "c.value as color", "z.value as size", "(SELECT COUNT(time) - 1 FROM counter.item_scans is2 WHERE is2.qr_code_code = $1) AS count").
		From("scanned_item si").
		LeftJoin("counter.items i ON si.qr_code_code = i.code").
		LeftJoin("counter.settings b ON i.buyer_id = b.id").
		LeftJoin("counter.settings s ON i.style_id = s.id").
		LeftJoin("counter.settings c ON i.color_id = c.id").
		LeftJoin("counter.settings z ON i.size_id = z.id").
		PrefixExpr(deleteItemScan)

	sqln, args, err := data.ToSql()
	if err != nil {
		tx.Rollback(ctx)
		return scannedItem, err
	}

	if err := tx.QueryRow(ctx, sqln, args...).Scan(&scannedItem.Time, &scannedItem.QrCode, &scannedItem.Buyer, &scannedItem.Style, &scannedItem.Color, &scannedItem.Size, &scannedItem.Count); err != nil {
		tx.Rollback(ctx)
		return scannedItem, err
	}

	tx.Commit(ctx)

	return scannedItem, nil
}
