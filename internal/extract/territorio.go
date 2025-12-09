package extract

import (
	"ETL_adventure/internal/models/raw"
	"database/sql"
)

func ExtractTerritorio(db *sql.DB) ([]raw.TerritorioRAW, error) {
	q := `
		SELECT 
			T.TerritoryID,
			T.Name,
			T.CountryRegionCode AS CodigoPais,
			CR.Name AS Pais,
			T.[Group] AS Continente
		FROM Sales.SalesTerritory AS T
		LEFT JOIN Person.CountryRegion AS CR
			ON T.CountryRegionCode = CR.CountryRegionCode
	`
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []raw.TerritorioRAW{}
	for rows.Next() {
		var r raw.TerritorioRAW
		err := rows.Scan(&r.TerritoryID, &r.Name, &r.CodigoPais, &r.Pais, &r.Continente)
		if err != nil {
			return nil, err
		}
		result = append(result, r)
	}

	return result, nil
}
