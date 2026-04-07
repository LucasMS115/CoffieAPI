package database

import "database/sql"

// NewPostgresConn opens a connection pool to PostgreSQL.
func NewPostgresConn(dataSourceName string) (*sql.DB, error) {
	databaseConnection, openError := sql.Open("postgres", dataSourceName)
	if openError != nil {
		return nil, openError
	}

	if pingError := databaseConnection.Ping(); pingError != nil {
		return nil, pingError
	}

	return databaseConnection, nil
}
