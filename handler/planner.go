package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/commute-my/ducky"
	"github.com/commute-my/ducky/motis"
)

var SupportedTransitModes = []motis.TransitMode{motis.TransitModeSubway, motis.TransitModeBus, motis.TransitModeOther}

type DuckyPlan struct {
	From        DuckyStop        `json:"from"`
	To          DuckyStop        `json:"to"`
	Itineraries []DuckyItinerary `json:"itineraries"`
}

type DuckyStop struct {
	Name   string  `json:"name"`
	StopID string  `json:"stop_id"`
	Lat    float64 `json:"lat"`
	Lon    float64 `json:"lon"`
}

type DuckyItinerary struct {
	Duration  int        `json:"duration"`
	StartTime time.Time  `json:"start_time"`
	EndTime   time.Time  `json:"end_time"`
	Transfers int        `json:"transfers"`
	Legs      []DuckyLeg `json:"legs"`
}

type DuckyLeg struct {
	Mode              motis.TransitMode `json:"mode"`
	From              DuckyStop         `json:"from"`
	To                DuckyStop         `json:"to"`
	Duration          int               `json:"duration"`
	StartTime         time.Time         `json:"start_time"`
	EndTime           time.Time         `json:"end_time"`
	Headsign          string            `json:"headsign"`
	RouteShortName    string            `json:"route_short_name"`
	IntermediateStops []DuckyStop       `json:"intermediate_stops"`
}

func DuckyPlanFromMotis(p motis.Plan) DuckyPlan {
	itineraries := make([]DuckyItinerary, len(p.Itineraries))

	for i, it := range p.Itineraries {
		legs := make([]DuckyLeg, len(it.Legs))

		for j, l := range it.Legs {
			intermediateStops := make([]DuckyStop, len(l.IntermediateStops))

			for k, is := range l.IntermediateStops {
				intermediateStops[k] = DuckyStop{
					Name:   is.Name,
					StopID: ducky.ConvertToID(l.From.StopID),
					Lat:    is.Lat,
					Lon:    is.Lon,
				}
			}

			legs[j] = DuckyLeg{
				Mode: l.Mode,
				From: DuckyStop{
					Name:   l.From.Name,
					StopID: ducky.ConvertToID(l.From.StopID),
					Lat:    l.From.Lat,
					Lon:    l.From.Lon,
				},
				To: DuckyStop{
					Name:   l.To.Name,
					StopID: ducky.ConvertToID(l.From.StopID),
					Lat:    l.To.Lat,
					Lon:    l.To.Lon,
				},
				Duration:          l.Duration,
				StartTime:         l.StartTime,
				EndTime:           l.EndTime,
				Headsign:          l.Headsign,
				RouteShortName:    l.RouteShortName,
				IntermediateStops: intermediateStops,
			}
		}

		itineraries[i] = DuckyItinerary{
			Duration:  it.Duration,
			StartTime: it.StartTime,
			EndTime:   it.EndTime,
			Transfers: it.Transfers,
			Legs:      legs,
		}
	}

	return DuckyPlan{
		From: DuckyStop{
			Name:   p.From.Name,
			StopID: ducky.ConvertToID(p.From.StopID),
			Lat:    p.From.Lat,
			Lon:    p.From.Lon,
		},
		To: DuckyStop{
			Name:   p.To.Name,
			StopID: ducky.ConvertToID(p.From.StopID),
			Lat:    p.To.Lat,
			Lon:    p.To.Lon,
		},
		Itineraries: itineraries,
	}
}

type Planner struct {
	motisCli *motis.Client
}

func NewPlanner(motisCli *motis.Client) *Planner {
	return &Planner{
		motisCli: motisCli,
	}
}

func (h *Planner) Plan(w http.ResponseWriter, r *http.Request) {
	var data struct {
		From string `json:"from"`
		To   string `json:"to"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request format: %v", err), http.StatusBadRequest)
		return
	}

	if data.From == "" || data.To == "" {
		http.Error(w, "Both 'from' and 'to' fields are required", http.StatusBadRequest)
		return
	}

	fromID := ducky.ConvertToMotisID(data.From)
	toID := ducky.ConvertToMotisID(data.To)

	motisPlan, err := h.motisCli.Plan(time.Now(), fromID, toID, SupportedTransitModes)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get route: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(DuckyPlanFromMotis(motisPlan)); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}
}
