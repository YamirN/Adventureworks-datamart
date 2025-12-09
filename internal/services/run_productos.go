package services

import (
	"fmt"

	"ETL_adventure/internal/extract"
	"ETL_adventure/internal/load"
	"ETL_adventure/internal/transform"
)

func RunProductos(etx *ETLContext) error {

	// ---------------- EXTRACT ----------------
	rawRows, err := extract.ExtractProductos(etx.DBOrigen)
	if err != nil {
		return fmt.Errorf("extract productos error: %v", err)
	}
	fmt.Printf("Extract Productos → %d filas\n", len(rawRows))

	// ---------------- TRANSFORM ----------------
	dimRows := transform.TransformProductos(rawRows)
	fmt.Printf("Transform Productos → %d filas\n", len(dimRows))

	// ---------------- LOAD ----------------
	if err := load.LoadProductosBulk(etx.Ctx, etx.DBDestino, dimRows); err != nil {
		return fmt.Errorf("load productos error: %v", err)
	}

	fmt.Println("✔ ETL Productos finalizado")
	return nil
}
