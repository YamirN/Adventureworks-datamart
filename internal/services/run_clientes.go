package services

import (
	"fmt"

	"ETL_adventure/internal/extract"
	"ETL_adventure/internal/load"
	"ETL_adventure/internal/transform"
)

func RunClientes(etx *ETLContext) error {

	// Extract
	rawRows, err := extract.ExtractClientes(etx.DBOrigen)
	if err != nil {
		return fmt.Errorf("extract error: %v", err)
	}
	fmt.Printf("Extract → %d filas\n", len(rawRows))

	// Transform
	dimRows := transform.TransformClientes(rawRows)
	fmt.Printf("Transform → %d filas\n", len(dimRows))

	// Load
	if err := load.LoadClientesBulk(etx.Ctx, etx.DBDestino, dimRows); err != nil {
		return fmt.Errorf("load error: %v", err)
	}

	fmt.Println("✔ ETL Clientes finalizado")
	return nil
}
