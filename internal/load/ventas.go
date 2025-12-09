package load

import (
	"ETL_adventure/internal/models/fact"
	"database/sql"

	mssql "github.com/microsoft/go-mssqldb"
)

func LoadFactVentasBulk(db *sql.DB, rows []fact.FactVenta) error {

	stmt, err := db.Prepare(mssql.CopyIn(
		"FACT_Ventas",
		mssql.BulkOptions{},
		"ProductoKey",
		"ClienteKey",
		"TerritorioKey",
		"TiempoKey",
		"OrdenVentaID",
		"Cantidad",
		"PrecioUnitario",
		"DescuentoPrecioUnitario",
		"Total",
		"FechaCarga",
	))
	if err != nil {
		return err
	}

	for _, r := range rows {
		_, err = stmt.Exec(
			r.ProductoKey,
			r.ClienteKey,
			r.TerritorioKey,
			r.TiempoKey,
			r.OrdenVentaID,
			r.Cantidad,
			r.PrecioUnitario,
			r.DescuentoPrecioUnitario,
			r.Total,
			r.FechaCarga,
		)
		if err != nil {
			return err
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return stmt.Close()
}
