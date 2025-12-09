package services

import "fmt"

func RunAll(etx *ETLContext) error {

	fmt.Println("→ Ejecutando ETL Clientes…")
	if err := RunClientes(etx); err != nil {
		return err
	}

	fmt.Println("→ Ejecutando ETL Productos…")
	if err := RunProductos(etx); err != nil {
		return err
	}

	fmt.Println("→ Ejecutando ETL Territorio…")
	if err := RunTerritorio(etx); err != nil {
		return err
	}

	fmt.Println("→ Ejecutando ETL Tiempo…")
	if err := RunTiempo(etx); err != nil {
		return err
	}

	fmt.Println("→ Ejecutando ETL FactVentas…")
	if err := RunFactVentas(etx); err != nil {
		return err
	}

	fmt.Println("✔ Todos los ETL finalizados correctamente")
	return nil
}
