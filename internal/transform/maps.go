package transform

import "ETL_adventure/internal/models/dim"

func BuildTimeMap(rows []dim.DimTiempo) map[int]int {
	m := make(map[int]int, len(rows))

	for _, r := range rows {
		key := r.Fecha.Year()*10000 + int(r.Fecha.Month())*100 + r.Fecha.Day()
		m[key] = r.TiempoKey
	}

	return m
}
