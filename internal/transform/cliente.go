package transform

import (
	"time"

	"ETL_adventure/internal/etlutil"
	"ETL_adventure/internal/models/dim"
	"ETL_adventure/internal/models/raw"
)

func TransformClientes(rawRows []raw.ClienteRAW) []dim.DimCliente {
	now := time.Now().UTC()
	out := make([]dim.DimCliente, 0, len(rawRows))

	// 1. Deduplicación: Usar un mapa para asegurar que CustomerID es único (Buena Práctica)
	seen := make(map[int]bool)

	for _, r := range rawRows {

		if seen[r.CustomerID] {
			continue
		}
		seen[r.CustomerID] = true

		// 2. Transformaciones de Limpieza y Estandarización
		p := dim.DimCliente{
			ClienteID: r.CustomerID,

			Nombre:      etlutil.NormalizeString(r.NameCustomer, "Cliente Desconocido"),
			TipoCliente: etlutil.NormalizeString(r.TipoCliente, "Desconocido"),
			TipoPersona: etlutil.NormalizeCountryCode(r.PersonType, "NA"),
			Ciudad:      etlutil.NormalizeString(r.City, "Sin Ciudad"),
			Estado:      etlutil.NormalizeString(r.StateProvinceName, "Sin Estado"),
			Pais:        etlutil.NormalizeString(r.CountryRegionName, "Sin País"),

			PersonaID: r.PersonID,
			TiendaID:  r.StoreID,

			FechaCarga: now,
		}

		out = append(out, p)
	}
	return out
}
