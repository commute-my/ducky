package ors

import (
	"encoding/json"
	"fmt"
	"time"
)

type CustomTime struct {
    time.Time
}

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
    s := string(b)
    s = s[1 : len(s)-1] // Remove quotes
    
    if s == "0000-00-00T00:00:00Z" || s == "" {
        // Set to a default value or zero time
        *ct = CustomTime{time.Time{}}
        return nil
    }
    
    t, err := time.Parse(time.RFC3339, s)
    if err != nil {
        return err
    }
    *ct = CustomTime{t}
    return nil
}

// Direction represents the response from OpenRouteService.
type Direction struct {
	Type     string    `json:"type"`
	Bbox     []float64 `json:"bbox"`
	Features []struct {
		Bbox       []float64 `json:"bbox"`
		Type       string    `json:"type"`
		Properties struct {
			Transfers int `json:"transfers"`
			Fare      int `json:"fare"`
			Segments  []struct {
				Distance float64 `json:"distance"`
				Duration float64 `json:"duration"`
				Steps    []struct {
					Distance    float64 `json:"distance"`
					Duration    float64 `json:"duration"`
					Type        int     `json:"type"`
					Instruction string  `json:"instruction"`
					Name        string  `json:"name"`
					WayPoints   []int   `json:"way_points"`
				} `json:"steps"`
			} `json:"segments"`
			Legs []struct {
				Type              string    `json:"type"`
				DepartureLocation string    `json:"departure_location"`
				RouteType         int       `json:"route_type"`
				Distance          float64   `json:"distance"`
				Duration          float64   `json:"duration"`
				Departure         CustomTime `json:"departure"`
				Arrival           CustomTime `json:"arrival"`
				Geometry          string    `json:"geometry"`
				Instructions      []struct {
					Distance    float64 `json:"distance"`
					Duration    float64 `json:"duration"`
					Type        int     `json:"type"`
					Instruction string  `json:"instruction"`
					Name        string  `json:"name"`
					WayPoints   []int   `json:"way_points"`
				} `json:"instructions,omitempty"`
				TripHeadsign              string `json:"trip_headsign,omitempty"`
				RouteLongName             string `json:"route_long_name,omitempty"`
				RouteShortName            string `json:"route_short_name,omitempty"`
				RouteDesc                 string `json:"route_desc,omitempty"`
				FeedID                    string `json:"feed_id,omitempty"`
				TripID                    string `json:"trip_id,omitempty"`
				RouteID                   string `json:"route_id,omitempty"`
				IsInSameVehicleAsPrevious bool   `json:"is_in_same_vehicle_as_previous,omitempty"`
				Stops                     []struct {
					StopID               string    `json:"stop_id"`
					Name                 string    `json:"name"`
					Location             []float64 `json:"location"`
					ArrivalCancelled     bool      `json:"arrival_cancelled"`
					DepartureTime        CustomTime `json:"departure_time,omitempty"`
					PlannedDepartureTime CustomTime `json:"planned_departure_time,omitempty"`
					DepartureCancelled   bool      `json:"departure_cancelled"`
					ArrivalTime          CustomTime `json:"arrival_time,omitempty"`
					PlannedArrivalTime   CustomTime `json:"planned_arrival_time,omitempty"`
				} `json:"stops,omitempty"`
			} `json:"legs"`
			WayPoints []int `json:"way_points"`
			Summary   struct {
				Distance float64 `json:"distance"`
				Duration float64 `json:"duration"`
			} `json:"summary"`
		} `json:"properties"`
		Geometry struct {
			Coordinates [][]float64 `json:"coordinates"`
			Type        string      `json:"type"`
		} `json:"geometry"`
	} `json:"features"`
	Metadata struct {
		Attribution string `json:"attribution"`
		Service     string `json:"service"`
		Timestamp   int64  `json:"timestamp"`
		Query       struct {
			Coordinates [][]float64 `json:"coordinates"`
			Profile     string      `json:"profile"`
			ProfileName string      `json:"profileName"`
			Format      string      `json:"format"`
		} `json:"query"`
		Engine struct {
			Version   string    `json:"version"`
			BuildDate CustomTime `json:"build_date"`
			GraphDate CustomTime `json:"graph_date"`
		} `json:"engine"`
	} `json:"metadata"`
}

// GetDirection fetches directions from OpenRouteService API.
func (c *Client) GetDirection(startLng, startLat, endLng, endLat float64) (Direction, error) {	
	url := c.ResolveURL("/directions/rapid-rail-kl")
	query := url.Query()
	query.Set("start", fmt.Sprintf("%f,%f", startLng, startLat))
	query.Set("end", fmt.Sprintf("%f,%f", endLng, endLat))
	url.RawQuery = query.Encode()

	var dir Direction

	resp, err := c.httpCli.Get(url.String())
	if err != nil {
		return dir, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&dir); err != nil {
		return dir, err
	}

	return dir, nil
}