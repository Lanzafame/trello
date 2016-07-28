package main

import (
	"encoding/json"
	"os"
)

func parseConfig(path string) (config, error) {
	f, err := os.Open(path)
	if err != nil {
		return config{}, err
	}
	defer f.Close()

	cfg := config{}
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		return config{}, err
	}

	return cfg, nil
}
