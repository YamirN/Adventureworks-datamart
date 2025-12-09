package raw

import "database/sql"

type TerritorioRAW struct {
	TerritoryID int
	Name        sql.NullString
	CodigoPais  sql.NullString
	Pais        sql.NullString
	Continente  sql.NullString
}
