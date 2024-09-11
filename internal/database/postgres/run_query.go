package postgres

import (
	"context"
	"database/sql"
)

func (pr *postgresRepo) RunQuery(ctx context.Context, query string) ([][]string, error) {
	pr.RLock()
	defer pr.RUnlock()

	rows, err := pr.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([][]string, 0)

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	result = append(result, columns)

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}

		var value string
		row := make([]string, len(columns))
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			row[i] = value
		}
		result = append(result, row)
	}

	return result, nil
}
