package db

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlserver"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"database/sql"
)

// RunMigrations ejecuta todas las migraciones pendientes
func RunMigrations(db *sql.DB, migrationsFolder string) error {
	driver, err := sqlserver.WithInstance(db, &sqlserver.Config{})
	if err != nil {
		return fmt.Errorf("error creando driver SQL Server: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationsFolder,
		"sqlserver", driver,
	)
	if err != nil {
		return fmt.Errorf("error inicializando migrate: %v", err)
	}

	// Aplica todas las migraciones pendientes
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error aplicando migraciones: %v", err)
	}

	log.Println("🎉 Migraciones aplicadas correctamente")
	return nil
}
