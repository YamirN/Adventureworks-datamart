package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"ETL_adventure/config"
	"ETL_adventure/internal/db"
	"ETL_adventure/internal/services"
)

func main() {

	// Load config
	cfg, err := config.Load("config/config.json")
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	// Global timeout
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(cfg.ETL.TimeoutMinutes)*time.Minute,
	)
	defer cancel()

	// Open connections once
	dbOrigen, err := db.OpenSQLServer(cfg.Origen.Server, cfg.Origen.Database, cfg.Origen.TrustedConnection, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer dbOrigen.Close()

	dbDestino, err := db.OpenSQLServer(cfg.Destino.Server, cfg.Destino.Database, cfg.Destino.TrustedConnection, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer dbDestino.Close()

	// Migrations
	fmt.Println("Ejecutando migraciones…")
	if err := db.RunMigrations(dbDestino, "migrations"); err != nil {
		log.Fatalf("migration error: %v", err)
	}

	// Build ETLContext
	etlCtx := &services.ETLContext{
		Ctx:       ctx,
		Config:    cfg,
		DBOrigen:  dbOrigen,
		DBDestino: dbDestino,
	}

	// Run all ETLs
	fmt.Println("Ejecutando todo el ETL…")
	if err := services.RunAll(etlCtx); err != nil {
		log.Fatalf("etl error: %v", err)
	}
}
