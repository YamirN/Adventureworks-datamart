package transform

import (
	"ETL_adventure/internal/etlutil"
	"ETL_adventure/internal/models/dim"
	"ETL_adventure/internal/models/raw"
	"time"
)

func TransformTerritorio(rawRows []raw.TerritorioRAW) []dim.DimTerritorio {

	out := make([]dim.DimTerritorio, 0, len(rawRows))
	now := time.Now().UTC()

	seen := make(map[int]struct{})

	for _, r := range rawRows {
		if _, exists := seen[r.TerritoryID]; exists {
			continue
		}
		seen[r.TerritoryID] = struct{}{}

		row := dim.DimTerritorio{
			TerritorioID:     r.TerritoryID,
			NombreTerritorio: etlutil.NormalizeString(r.Name, "Sin territorio"),
			CodigoPais:       etlutil.NormalizeCountryCode(r.CodigoPais, "NA"),
			Pais:             etlutil.NormalizeString(r.Pais, "Desconocido"),
			Continente:       etlutil.NormalizeString(r.Continente, "Desconocido"),
			FechaCarga:       now,
		}

		out = append(out, row)
	}

	return out
}
