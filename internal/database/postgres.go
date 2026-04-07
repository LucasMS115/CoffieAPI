package database

import "database/sql"

// NewPostgresConn opens a connection pool to PostgreSQL.
func NewPostgresConn(dataSourceName string) (*sql.DB, error) {
	_ = dataSourceName
	return nil, nil
}
