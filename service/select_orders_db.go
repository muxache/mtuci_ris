package service

import (
	"context"
	"database/sql"

	"github.com/muxache/mtuci_ris/data"
)

func SelectFromORDERS(db *sql.DB) ([]data.Orders, error) {
	ctx := context.Background()
	var orders []data.Orders
	request := "SELECT order_id, description, order_date, close_date, master_date FROM ORDERS"
	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := db.QueryContext(ctx, request)
	if err != nil {
		return []data.Orders{}, err
	}
	var o data.Orders
	for rows.Next() {

		// Get values from row.
		err := rows.Scan(&o.Order_ID, &o.Description, &o.Order_date, &o.Close_date, &o.Master_date)
		if err != nil {
			return orders, err
		}

		orders = append(orders, o)
	}
	return orders, nil
}
