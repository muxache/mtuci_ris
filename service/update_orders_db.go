package service

import (
	"context"
	"database/sql"
	"time"
)

func UpdateEmployee(id int, description string, master_date, close_date time.Time, db *sql.DB) (int64, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := "UPDATE operator_api.ORDERS SET description = @description, master_date = @master_date, close_date = @close_date WHERE order_id = @odrer_id"

	// Execute non-query with named parameters
	result, err := db.ExecContext(
		ctx,
		tsql,
		sql.Named("odrer_id", id),
		sql.Named("description", description),
		sql.Named("master_date", master_date),
		sql.Named("close_date", close_date))
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}
