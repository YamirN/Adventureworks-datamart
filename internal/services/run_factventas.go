package services

import (
	"fmt"

	"ETL_adventure/internal/extract"
	"ETL_adventure/internal/load"
	"ETL_adventure/internal/transform"
)

// RunFactVentas ejecuta todo el ETL del FactVentas
func RunFactVentas(etl *ETLContext) error {

	fmt.Println("Iniciando ETL FactVentas...")

	// =====================================
	// 1. EXTRACT
	// =====================================
	rawRows, err := extract.ExtractVentas(etl.DBOrigen)
	if err != nil {
		return fmt.Errorf("extract ventas error: %w", err)
	}
	fmt.Printf("Extract: %d filas de ventas\n", len(rawRows))

	// =====================================
	// 2. LOAD MAPS (llaves foráneas)
	// =====================================
	prodMap, err := transform.LoadMap[int, int](etl.Ctx, etl.DBDestino,
		"DIM_Producto", "ProductoID", "ProductoKey")
	if err != nil {
		return fmt.Errorf("producto map error: %w", err)
	}

	cliMap, err := transform.LoadMap[int, int](etl.Ctx, etl.DBDestino,
		"DIM_Cliente", "ClienteID", "ClienteKey")
	if err != nil {
		return fmt.Errorf("cliente map error: %w", err)
	}

	terrMap, err := transform.LoadMap[int, int](etl.Ctx, etl.DBDestino,
		"DIM_Territorio", "TerritorioID", "TerritorioKey")
	if err != nil {
		return fmt.Errorf("territorio map error: %w", err)
	}

	dimTiempoRows, err := load.GetAllDimTiempo(etl.DBDestino)
	if err != nil {
		return fmt.Errorf("tiempo map error: %w", err)
	}

	timeMap := transform.BuildTimeMap(dimTiempoRows)

	fmt.Println("Maps cargados correctamente.")

	// =====================================
	// 3. TRANSFORM
	// =====================================
	factRows := transform.TransformVentas(rawRows, prodMap, cliMap, terrMap, timeMap)
	fmt.Printf("Transform: %d filas facturación\n", len(factRows))

	// =====================================
	// 4. LOAD
	// =====================================
	if err := load.LoadFactVentasBulk(etl.DBDestino, factRows); err != nil {
		return fmt.Errorf("load fact ventas error: %w", err)
	}

	fmt.Println("ETL FactVentas finalizado con éxito.")
	return nil
}
