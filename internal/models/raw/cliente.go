package raw

import "database/sql"

type ClienteRAW struct {
	CustomerID        int
	PersonID          int
	StoreID           *int
	NameCustomer      sql.NullString
	TipoCliente       sql.NullString
	PersonType        sql.NullString
	City              sql.NullString
	StateProvinceName sql.NullString
	CountryRegionName sql.NullString
}
