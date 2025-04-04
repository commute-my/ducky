package handler

import (
	"encoding/json"
	"net/http"

	"github.com/commute-my/ducky/motis"
)

type DuckyGeocodeArea struct {
	Name string `json:"name"`
}

type DuckyGeocode struct {
	Type  motis.GeocodeLocationType `json:"type"`
	ID    string                    `json:"id"`
	Name  string                    `json:"name"`
	Areas []DuckyGeocodeArea        `json:"areas"`
	Lat   float64                   `json:"lat"`
	Lon   float64                   `json:"lon"`
}

func DuckyGeocodeFromMotis(r motis.GeocodeResult) DuckyGeocode {
	areas := make([]DuckyGeocodeArea, len(r.Areas))
	for i, area := range r.Areas {
		areas[i] = DuckyGeocodeArea{
			Name: area.Name,
		}
	}

	return DuckyGeocode{
		Type:  r.Type,
		ID:    r.ID,
		Name:  r.Name,
		Areas: areas,
		Lat:   r.Lat,
		Lon:   r.Lon,
	}
}

type Geocoder struct {
	motisCli *motis.Client
}

func NewGeocoder(motisCli *motis.Client) *Geocoder {
	return &Geocoder{
		motisCli: motisCli,
	}
}

func (h *Geocoder) Search(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Query string `json:"query"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if data.Query == "" {
		http.Error(w, "query is required", http.StatusBadRequest)
		return
	}

	motisGeocodes, err := h.motisCli.Geocode(data.Query, "en")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	geocodes := make([]DuckyGeocode, len(motisGeocodes))
	for i, motisGeocode := range motisGeocodes {
		geocodes[i] = DuckyGeocodeFromMotis(motisGeocode)
	}

	json.NewEncoder(w).Encode(geocodes)
}
