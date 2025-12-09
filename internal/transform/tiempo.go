package transform

import (
	"ETL_adventure/internal/models/dim"
	"time"
)

var monthNames = [...]string{
	"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio",
	"Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre",
}

var dayNames = [...]string{
	"Domingo", "Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado",
}

func GenerateDimTiempo(start, end time.Time) []dim.DimTiempo {
	totalDays := int(end.Sub(start).Hours()/24) + 1
	out := make([]dim.DimTiempo, 0, totalDays+1)
	now := time.Now().UTC()

	// Fila "desconocida"
	out = append(out, dim.DimTiempo{
		TiempoKey: 19000101,
		Fecha:     time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC),

		Anio:      1900,
		Mes:       1,
		NombreMes: "Desconocido",
		Dia:       1,
		DiaSemana: 0,
		NombreDia: "Desconocido",
		Trimestre: 0,
		SemanaISO: 0,

		EsFinDeSemana: false,
		FechaCarga:    now,
	})

	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		year, month, day := d.Year(), int(d.Month()), d.Day()
		weekday := d.Weekday()

		out = append(out, dim.DimTiempo{
			TiempoKey:     year*10000 + month*100 + day,
			Fecha:         d,
			Anio:          year,
			Mes:           month,
			NombreMes:     monthNames[month-1],
			Dia:           day,
			DiaSemana:     int(weekday),
			NombreDia:     dayNames[int(weekday)],
			Trimestre:     (month-1)/3 + 1,
			SemanaISO:     isoWeek(d),
			EsFinDeSemana: weekday == time.Saturday || weekday == time.Sunday,
			FechaCarga:    now,
		})
	}

	return out
}

func isoWeek(t time.Time) int {
	_, w := t.ISOWeek()
	return w
}
