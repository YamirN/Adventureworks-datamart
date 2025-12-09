package etlutil

import (
	"database/sql"
	"strings"
)

// NormalizeCountryCode normaliza códigos ISO a mayúsculas.
// Ej: "us" -> "US", "Ca" -> "CA".
func NormalizeCountryCode(ns sql.NullString, def string) string {
	if !ns.Valid {
		return def
	}

	s := strings.TrimSpace(ns.String)
	if s == "" {
		return def
	}

	return strings.ToUpper(s)
}
