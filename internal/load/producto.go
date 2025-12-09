package load

import (
	"context"
	"database/sql"

	"ETL_adventure/internal/models/dim"

	mssql "github.com/microsoft/go-mssqldb"
)

// LoadProductosBulk realiza bulk insert en la tabla destino. La tabla debe tener ProductoKey INT IDENTITY.
func LoadProductosBulk(ctx context.Context, db *sql.DB, rows []dim.DimProducto) error {
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
	if _, err := tx.ExecContext(ctx, "DELETE FROM dbo.DIM_Producto"); err != nil {
		tx.Rollback()
		return err
	}

	// Bulk CopyIn
	stmt, err := tx.Prepare(mssql.CopyIn(
		"DIM_Producto",
		mssql.BulkOptions{},
		"ProductoID",
		"nombre_producto",
		"costo_estandar",
		"categoria",
		"subcategoria",
		"fecha_carga",
	))
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, r := range rows {
		if _, err := stmt.Exec(r.ProductoID, r.Nombre, r.CostoEstandar, r.Categoria, r.Subcategoria, r.FechaCarga); err != nil {
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
