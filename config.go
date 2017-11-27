package main

import "encoding/json"
import "io/ioutil"

func Parse(path string) (Config, error) {
	var config Config

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(bytes, &config)

	if err != nil {
		return config, err
	}

	return config, nil
}

type Config struct {
	Fetch FetchConfig `json:"fetch"`
	WebUI WebUIConfig `json:"webui"`
	Sort  SortConfig  `json:"sort"`
}

type WebUIConfig struct {
	Enable bool   `json:"enable"`
	Host   string `json:"host"`
}

type FetchConfig struct {
	CheckFilesystem bool              `json:"check_filesystem"`
	SavePathDefault string            `json:"save_path_default"`
	SavePaths       map[string]string `json:"save_paths"`
	SearchPaths     []string          `json:"search_paths"`
}

type SortConfig struct {
	ProvidersToMove []string          `json:"providers_to_move"`
	SavePathDefault string            `json:"save_path_default"`
	SavePaths       map[string]string `json:"save_paths"`
	SearchPaths     []string          `json:"search_paths"`
}