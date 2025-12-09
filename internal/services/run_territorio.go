package services

import (
	"fmt"

	"ETL_adventure/internal/extract"
	"ETL_adventure/internal/load"
	"ETL_adventure/internal/transform"
)

func RunTerritorio(etx *ETLContext) error {

	// ---------------- EXTRACT ----------------
	rawRows, err := extract.ExtractTerritorio(etx.DBOrigen)
	if err != nil {
		return fmt.Errorf("extract territorio error: %v", err)
	}
	fmt.Printf("Extract Territorio → %d filas\n", len(rawRows))

	// ---------------- TRANSFORM ----------------
	dimRows := transform.TransformTerritorio(rawRows)
	fmt.Printf("Transform Territorio → %d filas\n", len(dimRows))

	// ---------------- LOAD ----------------
	if err := load.LoadTerritorioBulk(etx.Ctx, etx.DBDestino, dimRows); err != nil {
		return fmt.Errorf("load territorio error: %v", err)
	}

	fmt.Println("✔ ETL Territorio finalizado")
	return nil
}
