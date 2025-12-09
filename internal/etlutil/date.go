package etlutil

import "time"

func FechaToKey(t time.Time) int {
	return t.Year()*10000 + int(t.Month())*100 + t.Day()
}

func KeyOrDefault(m map[int]int, key int, def int) int {
	if v, ok := m[key]; ok {
		return v
	}
	return def
}
