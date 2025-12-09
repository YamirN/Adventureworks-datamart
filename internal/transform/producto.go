package transform

import (
	"time"

	"ETL_adventure/internal/etlutil"
	"ETL_adventure/internal/models/dim"
	"ETL_adventure/internal/models/raw"
)

// TransformProductos limpia y transforma raw -> dim
func TransformProductos(rawRows []raw.ProductoRAW) []dim.DimProducto {
	now := time.Now().UTC()
	out := make([]dim.DimProducto, 0, len(rawRows))
	seen := make(map[int]bool)

	for _, r := range rawRows {
		if seen[r.ProductID] {
			continue
		}
		seen[r.ProductID] = true
		p := dim.DimProducto{
			ProductoID:    r.ProductID,
			Nombre:        etlutil.NormalizeString(r.Name, "Producto desconocido"),
			CostoEstandar: r.StandardCost.Round(2),
			Categoria:     etlutil.NormalizeString(r.CategoryName, "Sin categoría"),
			Subcategoria:  etlutil.NormalizeString(r.SubcatName, "Sin subcategoría"),
			FechaCarga:    now,
		}
		out = append(out, p)
	}
	return out
}
