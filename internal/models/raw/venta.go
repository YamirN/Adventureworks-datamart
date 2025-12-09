package raw

import (
	"database/sql"

	"github.com/shopspring/decimal"
)

type VentaRAW struct {
	SalesOrderID      int
	CustomerID        int
	StoreID           int
	ProductID         int
	TerritoryID       int
	OrderQty          int
	UnitPrice         decimal.Decimal
	UnitPriceDiscount decimal.Decimal
	LineTotal         decimal.Decimal
	OrderDate         sql.NullTime
}
