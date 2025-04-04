package ducky

import "strings"

var MotisStationMappings = make(map[string]string, 0)

func init() {
	for _, line := range Lines {
		for _, station := range line.Stations {
			motisID := strings.Join([]string{line.MotisPrefix, station.ID}, "_")

			MotisStationMappings[motisID] = station.ID
		}
	}
}

func ConvertToMotisID(id string) string {
	for _, line := range Lines {
		for _, station := range line.Stations {
			if station.ID == id {
				return strings.Join([]string{line.MotisPrefix, id}, "_")
			}
		}
	}
	return id
}

func ConvertToID(motisID string) string {
	if id, exists := MotisStationMappings[motisID]; exists {
		return id
	}

	return motisID
}
