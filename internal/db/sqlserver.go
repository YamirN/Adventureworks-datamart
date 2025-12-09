package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/microsoft/go-mssqldb"
)

// OpenSQLServer abre una conexión estándar a SQL Server.
func OpenSQLServer(server, database string, trusted bool, timeout time.Duration) (*sql.DB, error) {
	// build connection string
	conn := fmt.Sprintf("server=%s;database=%s;encrypt=disable", server, database)
	if trusted {
		conn = fmt.Sprintf("server=%s;database=%s;trusted_connection=yes;encrypt=disable", server, database)
	}
	// opcional: puedes añadir user/password si no usas trusted
	return sql.Open("sqlserver", conn)
}
