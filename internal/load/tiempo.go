package load

import (
	"ETL_adventure/internal/models/dim"
	"context"
	"database/sql"

	mssql "github.com/microsoft/go-mssqldb"
)

func LoadDimTiempo(ctx context.Context, db *sql.DB, rows []dim.DimTiempo) error {

	txn, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	stmt, err := txn.PrepareContext(ctx, mssql.CopyIn(
		"DIM_Tiempo",
		mssql.BulkOptions{},
		"TiempoKey",
		"Fecha",
		"Anio",
		"Mes",
		"NombreMes",
		"Dia",
		"DiaSemana",
		"NombreDia",
		"Trimestre",
		"SemanaISO",
		"EsFinDeSemana",
		"FechaCarga",
	))
	if err != nil {
		txn.Rollback()
		return err
	}

	for _, r := range rows {
		_, err := stmt.ExecContext(
			ctx,
			r.TiempoKey,
			r.Fecha,
			r.Anio,
			r.Mes,
			r.NombreMes,
			r.Dia,
			r.DiaSemana,
			r.NombreDia,
			r.Trimestre,
			r.SemanaISO,
			r.EsFinDeSemana,
			r.FechaCarga,
		)
		if err != nil {
			txn.Rollback()
			return err
		}
	}

	// Indica fin del Bulk Insert
	_, err = stmt.ExecContext(ctx)
	if err != nil {
		txn.Rollback()
		return err
	}

	err = stmt.Close()
	if err != nil {
		txn.Rollback()
		return err
	}

	return txn.Commit()
}
