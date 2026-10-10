package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func cityName(ctx context.Context, db *sql.DB, id int) (string, error) {
	var name string
	err := db.QueryRowContext(ctx, "SELECT name FROM cities WHERE id = $1", id).Scan(&name)
	if err != nil {
		return "", fmt.Errorf("city %d: %w", id, err)
	}
	return name, nil
}
func main() {
	err := fmt.Errorf("lookup: %w", sql.ErrNoRows)
	fmt.Println(errors.Is(err, sql.ErrNoRows))
}
