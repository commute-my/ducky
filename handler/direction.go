package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/commute-my/ducky/ors"
)

type RouteResponse struct {
	TotalDistance   float64       `json:"total_distance_meters"`
	TotalDuration   float64       `json:"total_duration_seconds"`
	HumanDistance   string        `json:"total_distance_human"`
	HumanDuration   string        `json:"total_duration_human"`
	Transfers       int           `json:"transfers"`
	RouteSegments   []RouteSegment `json:"segments"`
	BoundingBox     []float64     `json:"bounding_box,omitempty"`
	Attribution     string        `json:"attribution,omitempty"`
}

// RouteSegment represents a segment of the journey.
type RouteSegment struct {
	Type               string      `json:"type"`
	TransportMode      string      `json:"transport_mode"`
	Line               string      `json:"line,omitempty"`
	From               string      `json:"from"`
	To                 string      `json:"to,omitempty"`
	DepartureTime      time.Time  `json:"departure_time"`
	ArrivalTime        time.Time  `json:"arrival_time"`
	Distance           float64     `json:"distance_meters"`
	Duration           float64     `json:"duration_seconds"`
	HumanDistance      string      `json:"distance_human"`
	HumanDuration      string      `json:"duration_human"`
	HeadSign           string      `json:"headsign,omitempty"`
	Instructions       []string    `json:"instructions,omitempty"`
	Stops              []Stop      `json:"stops,omitempty"`
	RouteID            string      `json:"route_id,omitempty"`
	TripID             string      `json:"trip_id,omitempty"`
}

// Stop represents a transit stop
type Stop struct {
	ID            string     `json:"id"`
	Name          string     `json:"name"`
	Location      []float64  `json:"location"`
	ArrivalTime   *time.Time `json:"arrival_time"`
	DepartureTime *time.Time `json:"departure_time"`
	IsCancelled   bool       `json:"is_cancelled"`
}

type Direction struct {
	orsCli *ors.Client
}

func NewDirection(orsCli *ors.Client) *Direction {
	return &Direction{
		orsCli: orsCli,
	}
}

func (h *Direction) Direction(w http.ResponseWriter, r *http.Request) {
	var data struct {
		StartLongitude float64 `json:"start_longitude"`
		StartLangitude float64 `json:"start_latitude"`
		EndLongitude   float64 `json:"end_longitude"`
		EndLatitude    float64 `json:"end_latitude"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dir, err := h.orsCli.GetDirection(data.StartLongitude, data.StartLangitude, data.EndLongitude, data.EndLatitude)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(dir.Features) == 0 {
		http.Error(w, "No route found", http.StatusNotFound)
		return
	}

	// TODO: Format the response to a user-friendly JSON format.
	json.NewEncoder(w).Encode(dir)
}