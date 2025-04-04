package motis

import "encoding/json"

type GeocodeLocationType string

const (
	GeocodeLocationTypeCountry GeocodeLocationType = "ADDRESS"
	GeocodeLocationTypeRegion  GeocodeLocationType = "PLACE"
	GeocodeLocationTypeStop    GeocodeLocationType = "STOP"
)

type GeocodeArea struct {
	Name       string  `json:"name"`
	AdminLevel float64 `json:"adminLevel"`
	Matched    bool    `json:"matched"`
	Default    bool    `json:"default"`
}

type GeocodeResult struct {
	Type   GeocodeLocationType `json:"type"`
	Tokens [][]float64         `json:"tokens"`
	Name   string              `json:"name"`
	ID     string              `json:"id"`
	Lat    float64             `json:"lat"`
	Lon    float64             `json:"lon"`
	Level  float64             `json:"level,omitempty"`
	Areas  []GeocodeArea       `json:"areas"`
	Score  float64             `json:"score"`
}

func (c *Client) Geocode(text, language string) ([]GeocodeResult, error) {
	var r []GeocodeResult

	url := c.ResolveURL("/geocode")
	query := url.Query()
	query.Set("text", text)
	query.Set("language", language)
	url.RawQuery = query.Encode()

	resp, err := c.httpCli.Get(url.String())
	if err != nil {
		return r, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return r, err
	}

	return r, nil
}
