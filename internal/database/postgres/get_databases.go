package postgres

import (
	"context"
)

func (pr *postgresRepo) GetAllDatabases(ctx context.Context) ([]string, error) {
	query := `SELECT datname FROM pg_catalog.pg_database WHERE datistemplate=$1`

	rows, err := pr.db.QueryContext(ctx, query, false)
	if err != nil {
		return nil, err
	}

	result := []string{}

	for rows.Next() {
		var database string

		err := rows.Scan(&database)
		if err != nil {
			return nil, err
		}

		result = append(result, database)
	}

	return result, nil
}
