package etlutil

import (
	"database/sql"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var titleCase = cases.Title(language.English)

// NormalizeString limpia cadenas de texto para ETL.
// - Convierte NullString usando `def`
// - Recorta espacios
// - Convierte a lower y luego a TitleCase
// - Filtra valores inválidos ("N/A", "NULL", etc.)
func NormalizeString(ns sql.NullString, def string) string {

	if !ns.Valid {
		return def
	}

	s := strings.TrimSpace(ns.String)
	if s == "" {
		return def
	}

	// Filtrar valores que no aportan
	invalid := map[string]bool{
		"n/a":  true,
		"na":   true,
		"null": true,
		"none": true,
	}
	if invalid[strings.ToLower(s)] {
		return def
	}

	// Normalización
	s = strings.ToLower(s)

	return titleCase.String(s)
}
