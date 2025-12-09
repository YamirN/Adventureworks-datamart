package extract

import (
	"ETL_adventure/internal/models/raw"
	"database/sql"
)

func ExtractVentas(db *sql.DB) ([]raw.VentaRAW, error) {
	query := `
	SELECT 
		OD.SalesOrderID,
		OH.CustomerID,
		OD.ProductID,
		OH.TerritoryID,
		OD.OrderQty,
		OD.UnitPrice,
		OD.UnitPriceDiscount,
		OD.LineTotal,
		OH.OrderDate
	FROM Sales.SalesOrderDetail AS OD
	INNER JOIN Sales.SalesOrderHeader AS OH
	ON OD.SalesOrderID = OH.SalesOrderID;
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []raw.VentaRAW{}

	for rows.Next() {
		var r raw.VentaRAW
		err := rows.Scan(
			&r.SalesOrderID,
			&r.CustomerID,
			&r.ProductID,
			&r.TerritoryID,
			&r.OrderQty,
			&r.UnitPrice,
			&r.UnitPriceDiscount,
			&r.LineTotal,
			&r.OrderDate,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, r)
	}

	return result, nil
}
