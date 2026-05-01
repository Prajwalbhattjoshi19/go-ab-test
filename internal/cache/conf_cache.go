package cache

import (
	"context"
	"database/sql"
	"time"
)

var featuresCache []string

func RefreshFeatures(db *sql.DB) {
	features := make([]string, 25)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// QueryContext executes a query that returns rows, typically a SELECT statement (a command used to fetch data from a database table).
	rows, err := db.QueryContext(ctx, "SELECT * FROM features")
	defer rows.Close()
	if err != nil {
		return
	}

	featuresCache = features
}
