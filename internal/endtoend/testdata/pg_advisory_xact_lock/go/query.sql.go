// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const advisoryLock = `-- name: AdvisoryLock :many
SELECT pg_advisory_xact_lock($1)
`

func (q *Queries) AdvisoryLock(ctx context.Context, key int64) ([]sql.NullBool, error) {
	rows, err := q.db.QueryContext(ctx, advisoryLock, key)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sql.NullBool
	for rows.Next() {
		var pg_advisory_xact_lock sql.NullBool
		if err := rows.Scan(&pg_advisory_xact_lock); err != nil {
			return nil, err
		}
		items = append(items, pg_advisory_xact_lock)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}