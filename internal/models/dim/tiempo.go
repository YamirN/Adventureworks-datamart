package dim

import "time"

type DimTiempo struct {
	TiempoKey     int
	Fecha         time.Time
	Anio          int
	Mes           int
	NombreMes     string
	Dia           int
	DiaSemana     int
	NombreDia     string
	Trimestre     int
	SemanaISO     int
	EsFinDeSemana bool
	FechaCarga    time.Time
}
