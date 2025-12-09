package dim

import "time"

type DimTerritorio struct {
	TerritorioID     int
	NombreTerritorio string
	CodigoPais       string
	Pais             string
	Continente       string
	FechaCarga       time.Time
}
