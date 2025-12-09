package load

import (
	"ETL_adventure/internal/models/dim"
	"database/sql"
)

func GetAllDimTiempo(db *sql.DB) ([]dim.DimTiempo, error) {
	rows, err := db.Query(`
        SELECT 
            TiempoKey, Fecha, Anio, Mes, NombreMes,
            Dia, DiaSemana, NombreDia,
            Trimestre, SemanaISO, EsFinDeSemana, FechaCarga
        FROM DIM_Tiempo`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := []dim.DimTiempo{}

	for rows.Next() {
		var r dim.DimTiempo
		err := rows.Scan(
			&r.TiempoKey, &r.Fecha, &r.Anio, &r.Mes, &r.NombreMes,
			&r.Dia, &r.DiaSemana, &r.NombreDia,
			&r.Trimestre, &r.SemanaISO, &r.EsFinDeSemana, &r.FechaCarga,
		)
		if err != nil {
			return nil, err
		}
		out = append(out, r)
	}

	return out, rows.Err()
}
