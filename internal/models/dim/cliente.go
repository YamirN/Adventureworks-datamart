package dim

import "time"

type DimCliente struct {
	ClienteID   int
	PersonaID   int
	TiendaID    *int
	Nombre      string
	TipoCliente string
	TipoPersona string
	Ciudad      string
	Estado      string
	Pais        string
	FechaCarga  time.Time
}
