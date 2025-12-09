package transform

import (
	"ETL_adventure/internal/etlutil"
	"ETL_adventure/internal/models/fact"
	"ETL_adventure/internal/models/raw"
	"time"
)

func TransformVentas(
	rawRows []raw.VentaRAW,
	prodMap map[int]int,
	cliMap map[int]int,
	terrMap map[int]int,
	timeMap map[int]int,
) []fact.FactVenta {

	now := time.Now().UTC()
	out := make([]fact.FactVenta, 0, len(rawRows))

	for _, r := range rawRows {

		// ----- DIMENSIONES -----
		pk := prodMap[r.ProductID]
		ck := cliMap[r.CustomerID]
		tk := terrMap[r.TerritoryID]

		// ----- FECHA -----
		var fecha time.Time
		if r.OrderDate.Valid {
			fecha = r.OrderDate.Time
		} else {
			fecha = time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
		}

		// Convertir a YYYYMMDD
		key := etlutil.FechaToKey(fecha)

		// Obtener TiempoKey o usar default (19000101)
		tkTime := etlutil.KeyOrDefault(timeMap, key, 19000101)

		// ----- FACT -----
		row := fact.FactVenta{
			ProductoKey:             pk,
			ClienteKey:              ck,
			TerritorioKey:           tk,
			TiempoKey:               tkTime,
			OrdenVentaID:            r.SalesOrderID,
			Cantidad:                r.OrderQty,
			PrecioUnitario:          r.UnitPrice.Round(2),
			DescuentoPrecioUnitario: r.UnitPriceDiscount.Round(2),
			Total:                   r.LineTotal.Round(2),
			FechaCarga:              now,
		}

		out = append(out, row)
	}

	return out
}
