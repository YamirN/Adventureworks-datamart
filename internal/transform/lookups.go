package transform

import (
	"context"
	"database/sql"
	"fmt"
)

// LoadMap carga un mapa genérico desde una tabla DIM.
// keyCol: columna origen (ProductID, CustomerID, etc.)
// valCol: columna destino (ProductoKey , etc.)
// table: nombre de la tabla DIM
func LoadMap[K comparable, V any](ctx context.Context, db *sql.DB, table, keyCol, valCol string) (map[K]V, error) {

	query := fmt.Sprintf(`SELECT %s, %s FROM %s`, keyCol, valCol, table)

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make(map[K]V)

	for rows.Next() {
		var key K
		var val V
		if err := rows.Scan(&key, &val); err != nil {
			return nil, err
		}

		out[key] = val
	}

	return out, nil
}
