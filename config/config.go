package config

import (
	"encoding/json"
	"os"
)

type DBConf struct {
	Server            string `json:"server"`
	Database          string `json:"database"`
	TrustedConnection bool   `json:"trusted_connection"`
}

type ETLConf struct {
	TimeoutMinutes int `json:"timeout_minutes"`
}

type Config struct {
	Origen  DBConf  `json:"origen"`
	Destino DBConf  `json:"destino"`
	ETL     ETLConf `json:"etl"`
}

func Load(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var c Config
	if err := json.NewDecoder(f).Decode(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
