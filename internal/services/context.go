package services

import (
	"context"
	"database/sql"

	"ETL_adventure/config"
)

type ETLContext struct {
	Ctx       context.Context
	Config    *config.Config
	DBOrigen  *sql.DB
	DBDestino *sql.DB
}
