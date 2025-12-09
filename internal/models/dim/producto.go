package dim

import (
	"time"

	"github.com/shopspring/decimal"
)

// DimProducto representa la fila limpia para la DIM
type DimProducto struct {
	ProductoID    int
	Nombre        string
	CostoEstandar decimal.Decimal
	Categoria     string
	Subcategoria  string
	FechaCarga    time.Time
}
