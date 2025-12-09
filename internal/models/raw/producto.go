package raw

import (
	"database/sql"

	"github.com/shopspring/decimal"
)

// ProductoRAW representa las columnas tal como vienen del OLTP
type ProductoRAW struct {
	ProductID    int
	Name         sql.NullString
	StandardCost decimal.Decimal
	CategoryName sql.NullString
	SubcatName   sql.NullString
}
