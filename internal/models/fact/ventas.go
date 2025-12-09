package fact

import (
	"time"

	"github.com/shopspring/decimal"
)

type FactVenta struct {
	ProductoKey   int
	ClienteKey    int
	TerritorioKey int
	TiempoKey     int

	OrdenVentaID            int
	Cantidad                int
	PrecioUnitario          decimal.Decimal
	DescuentoPrecioUnitario decimal.Decimal
	Total                   decimal.Decimal
	FechaCarga              time.Time
}
