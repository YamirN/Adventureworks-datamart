package extract

import (
	"database/sql"

	"ETL_adventure/internal/models/raw"
)

// ExtractProductos extrae productos desde la BD origen
func ExtractProductos(db *sql.DB) ([]raw.ProductoRAW, error) {
	q := `SELECT P.ProductID, P.Name, P.StandardCost,CA.Name AS CategoryName, SUB.Name AS SubcatName
FROM Production.Product P
LEFT JOIN Production.ProductSubcategory SUB ON P.ProductSubcategoryID = SUB.ProductSubcategoryID
LEFT JOIN Production.ProductCategory CA ON SUB.ProductCategoryID = CA.ProductCategoryID
WHERE P.ProductSubcategoryID IS NOT NULL AND CA.Name IS NOT NULL`

	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []raw.ProductoRAW
	for rows.Next() {
		var r raw.ProductoRAW
		if err := rows.Scan(&r.ProductID, &r.Name, &r.StandardCost, &r.CategoryName, &r.SubcatName); err != nil {
			return nil, err
		}
		out = append(out, r)
	}
	return out, rows.Err()
}
