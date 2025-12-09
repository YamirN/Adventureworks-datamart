package services

import (
	"fmt"
	"time"

	"ETL_adventure/internal/load"
	"ETL_adventure/internal/transform"
)

func RunTiempo(etl *ETLContext) error {

	fmt.Println("Generando DIM_Tiempo...")

	// ---------------- EXTRACT (no existe) ----------------
	// Tiempo no se extrae, se genera directamente.

	// ---------------- TRANSFORM ----------------
	start := time.Date(2011, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2014, 12, 31, 0, 0, 0, 0, time.UTC)

	rows := transform.GenerateDimTiempo(start, end)
	fmt.Printf("Transform Tiempo → %d filas\n", len(rows))

	// ---------------- LOAD ----------------
	if err := load.LoadDimTiempo(etl.Ctx, etl.DBDestino, rows); err != nil {
		return fmt.Errorf("load tiempo error: %w", err)
	}

	fmt.Println("✔ ETL Tiempo finalizado")
	return nil
}
