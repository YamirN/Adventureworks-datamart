package load

import (
	"context"
	"database/sql"
	"fmt"

	"ETL_adventure/internal/models/dim"

	mssql "github.com/microsoft/go-mssqldb"
)

func LoadTerritorioBulk(ctx context.Context, db *sql.DB, rows []dim.DimTerritorio) error {

	txn, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	stmt, err := txn.PrepareContext(ctx, mssql.CopyIn(
		"DIM_Territorio",
		mssql.BulkOptions{},
		"TerritorioID",
		"NombreTerritorio",
		"CodigoPais",
		"Pais",
		"Continente",
		"FechaCarga",
	))
	if err != nil {
		return err
	}

	for _, row := range rows {

		// Aquí aplicamos ctx también
		_, err = stmt.ExecContext(
			ctx,
			row.TerritorioID,
			row.NombreTerritorio,
			row.CodigoPais,
			row.Pais,
			row.Continente,
			row.FechaCarga,
		)
		if err != nil {
			return fmt.Errorf("bulk insert error: %w", err)
		}
	}

	_, err = stmt.ExecContext(ctx)
	if err != nil {
		return err
	}

	if err = stmt.Close(); err != nil {
		return err
	}

	return txn.Commit()
}
