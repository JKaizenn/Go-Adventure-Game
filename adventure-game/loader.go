package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadLocations(filename string) (map[string]Location, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	var locations []Location
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&locations); err != nil {
		return nil, fmt.Errorf("could not decode JSON: %v", err)
	}

	locationMap := make(map[string]Location)
	for _, loc := range locations {
		locationMap[loc.Name] = loc
	}
	return locationMap, nil
}
