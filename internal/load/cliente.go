package load

import (
	"context"
	"database/sql"

	"ETL_adventure/internal/models/dim"

	mssql "github.com/microsoft/go-mssqldb"
)

func LoadClientesBulk(ctx context.Context, db *sql.DB, rows []dim.DimCliente) error {
	// Iniciar tx
	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	// Truncar (full load)
	if _, err := tx.ExecContext(ctx, "DELETE FROM dbo.DIM_Cliente"); err != nil {
		tx.Rollback()
		return err
	}

	// Bulk CopyIn
	stmt, err := tx.Prepare(mssql.CopyIn(
		"DIM_Cliente",
		mssql.BulkOptions{},
		"ClienteID",
		"PersonaID",
		"TiendaID",
		"Nombre",
		"TipoCliente",
		"TipoPersona",
		"Ciudad",
		"Estado",
		"Pais",
		"fecha_carga",
	))
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, r := range rows {
		if _, err := stmt.Exec(r.ClienteID, r.PersonaID, r.TiendaID, r.Nombre, r.TipoCliente, r.TipoPersona, r.Ciudad, r.Estado, r.Pais, r.FechaCarga); err != nil {
			stmt.Close()
			tx.Rollback()
			return err
		}
	}

	// finish
	if _, err := stmt.Exec(); err != nil {
		stmt.Close()
		tx.Rollback()
		return err
	}
	if err := stmt.Close(); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
