package handler

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/commute-my/ducky"
	"github.com/commute-my/ducky/motis"
)

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
				stop := DuckyStop{
					Name: is.Name,
					Lat:  is.Lat,
					Lon:  is.Lon,
				}

				for _, l := range ducky.Lines {
					for _, s := range l.Stations {
						if strings.Join([]string{l.MotisPrefix, s.ID}, "_") == is.StopID {
							stop.StopID = s.ID
						}
					}
				}

				intermediateStops[k] = stop
			}

			leg := DuckyLeg{
				Mode: l.Mode,
				From: DuckyStop{
					Name: l.From.Name,
					Lat:  l.From.Lat,
					Lon:  l.From.Lon,
				},
				To: DuckyStop{
					Name: l.To.Name,
					Lat:  l.To.Lat,
					Lon:  l.To.Lon,
				},
				Duration:          l.Duration,
				StartTime:         l.StartTime,
				EndTime:           l.EndTime,
				Headsign:          l.Headsign,
				RouteShortName:    l.RouteShortName,
				IntermediateStops: intermediateStops,
			}

			for _, l := range ducky.Lines {
				for _, s := range l.Stations {
					if strings.Join([]string{l.MotisPrefix, s.ID}, "_") == leg.From.StopID {
						leg.From.StopID = s.ID

						break
					}

					if strings.Join([]string{l.MotisPrefix, s.ID}, "_") == leg.To.StopID {
						leg.To.StopID = s.ID

						break
					}
				}
			}

			legs[j] = leg
		}

		itineraries[i] = DuckyItinerary{
			Duration:  it.Duration,
			StartTime: it.StartTime,
			EndTime:   it.EndTime,
			Transfers: it.Transfers,
			Legs:      legs,
		}
	}

	plan := DuckyPlan{
		From: DuckyStop{
			Name:   p.From.Name,
			StopID: p.From.StopID,
			Lat:    p.From.Lat,
			Lon:    p.From.Lon,
		},
		To: DuckyStop{
			Name:   p.To.Name,
			StopID: p.To.StopID,
			Lat:    p.To.Lat,
			Lon:    p.To.Lon,
		},
		Itineraries: itineraries,
	}

	for _, l := range ducky.Lines {
		for _, s := range l.Stations {
			if strings.Join([]string{l.MotisPrefix, s.ID}, "_") == plan.From.StopID {
				plan.From.StopID = s.ID

				break
			}

			if strings.Join([]string{l.MotisPrefix, s.ID}, "_") == plan.To.StopID {
				plan.To.StopID = s.ID

				break
			}
		}
	}

	return plan
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if data.From == "" || data.To == "" {
		http.Error(w, "from and to are required", http.StatusBadRequest)
		return
	}

	for _, l := range ducky.Lines {
		for _, s := range l.Stations {
			if s.ID == data.From {
				data.From = strings.Join([]string{l.MotisPrefix, s.ID}, "_")
			}

			if s.ID == data.To {
				data.To = strings.Join([]string{l.MotisPrefix, s.ID}, "_")
			}
		}
	}

	motisPlan, err := h.motisCli.Plan(time.Now(), data.From, data.To, []motis.TransitMode{motis.TransitModeSubway, motis.TransitModeBus, motis.TransitModeOther})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(DuckyPlanFromMotis(motisPlan))
}
