package motis

import (
	"encoding/json"
	"strings"
	"time"
)

type Plan struct {
	RequestParameters struct {
		Property1 string `json:"property1"`
		Property2 string `json:"property2"`
	} `json:"requestParameters"`
	DebugOutput struct {
		Property1 int `json:"property1"`
		Property2 int `json:"property2"`
	} `json:"debugOutput"`
	From struct {
		Name               string    `json:"name"`
		StopID             string    `json:"stopId"`
		Lat                float64   `json:"lat"`
		Lon                float64   `json:"lon"`
		Level              float64   `json:"level"`
		Arrival            time.Time `json:"arrival"`
		Departure          time.Time `json:"departure"`
		ScheduledArrival   time.Time `json:"scheduledArrival"`
		ScheduledDeparture time.Time `json:"scheduledDeparture"`
		ScheduledTrack     string    `json:"scheduledTrack"`
		Track              string    `json:"track"`
		VertexType         string    `json:"vertexType"`
	} `json:"from"`
	To struct {
		Name               string    `json:"name"`
		StopID             string    `json:"stopId"`
		Lat                float64   `json:"lat"`
		Lon                float64   `json:"lon"`
		Level              float64   `json:"level"`
		Arrival            time.Time `json:"arrival"`
		Departure          time.Time `json:"departure"`
		ScheduledArrival   time.Time `json:"scheduledArrival"`
		ScheduledDeparture time.Time `json:"scheduledDeparture"`
		ScheduledTrack     string    `json:"scheduledTrack"`
		Track              string    `json:"track"`
		VertexType         string    `json:"vertexType"`
	} `json:"to"`
	Direct []struct {
		Duration  int       `json:"duration"`
		StartTime time.Time `json:"startTime"`
		EndTime   time.Time `json:"endTime"`
		Transfers int       `json:"transfers"`
		Legs      []struct {
			Mode string `json:"mode"`
			From struct {
				Name               string    `json:"name"`
				StopID             string    `json:"stopId"`
				Lat                int       `json:"lat"`
				Lon                int       `json:"lon"`
				Level              int       `json:"level"`
				Arrival            time.Time `json:"arrival"`
				Departure          time.Time `json:"departure"`
				ScheduledArrival   time.Time `json:"scheduledArrival"`
				ScheduledDeparture time.Time `json:"scheduledDeparture"`
				ScheduledTrack     string    `json:"scheduledTrack"`
				Track              string    `json:"track"`
				VertexType         string    `json:"vertexType"`
			} `json:"from"`
			To struct {
				Name               string    `json:"name"`
				StopID             string    `json:"stopId"`
				Lat                int       `json:"lat"`
				Lon                int       `json:"lon"`
				Level              int       `json:"level"`
				Arrival            time.Time `json:"arrival"`
				Departure          time.Time `json:"departure"`
				ScheduledArrival   time.Time `json:"scheduledArrival"`
				ScheduledDeparture time.Time `json:"scheduledDeparture"`
				ScheduledTrack     string    `json:"scheduledTrack"`
				Track              string    `json:"track"`
				VertexType         string    `json:"vertexType"`
			} `json:"to"`
			Duration                 int       `json:"duration"`
			StartTime                time.Time `json:"startTime"`
			EndTime                  time.Time `json:"endTime"`
			ScheduledStartTime       time.Time `json:"scheduledStartTime"`
			ScheduledEndTime         time.Time `json:"scheduledEndTime"`
			RealTime                 bool      `json:"realTime"`
			Distance                 int       `json:"distance"`
			InterlineWithPreviousLeg bool      `json:"interlineWithPreviousLeg"`
			Headsign                 string    `json:"headsign"`
			RouteColor               string    `json:"routeColor"`
			RouteTextColor           string    `json:"routeTextColor"`
			RouteType                string    `json:"routeType"`
			AgencyName               string    `json:"agencyName"`
			AgencyURL                string    `json:"agencyUrl"`
			AgencyID                 string    `json:"agencyId"`
			TripID                   string    `json:"tripId"`
			RouteShortName           string    `json:"routeShortName"`
			Source                   string    `json:"source"`
			IntermediateStops        []struct {
				Name               string    `json:"name"`
				StopID             string    `json:"stopId"`
				Lat                int       `json:"lat"`
				Lon                int       `json:"lon"`
				Level              int       `json:"level"`
				Arrival            time.Time `json:"arrival"`
				Departure          time.Time `json:"departure"`
				ScheduledArrival   time.Time `json:"scheduledArrival"`
				ScheduledDeparture time.Time `json:"scheduledDeparture"`
				ScheduledTrack     string    `json:"scheduledTrack"`
				Track              string    `json:"track"`
				VertexType         string    `json:"vertexType"`
			} `json:"intermediateStops"`
			LegGeometry struct {
				Points string `json:"points"`
				Length int    `json:"length"`
			} `json:"legGeometry"`
			Steps []struct {
				RelativeDirection string  `json:"relativeDirection"`
				Distance          int     `json:"distance"`
				FromLevel         float64 `json:"fromLevel"`
				ToLevel           float64 `json:"toLevel"`
				OsmWay            int     `json:"osmWay"`
				Polyline          struct {
					Points string `json:"points"`
					Length int    `json:"length"`
				} `json:"polyline"`
				StreetName string `json:"streetName"`
				Exit       string `json:"exit"`
				StayOn     bool   `json:"stayOn"`
				Area       bool   `json:"area"`
			} `json:"steps"`
			Rental struct {
				SystemID         string `json:"systemId"`
				SystemName       string `json:"systemName"`
				URL              string `json:"url"`
				StationName      string `json:"stationName"`
				FromStationName  string `json:"fromStationName"`
				ToStationName    string `json:"toStationName"`
				RentalURIAndroid string `json:"rentalUriAndroid"`
				RentalURIIOS     string `json:"rentalUriIOS"`
				RentalURIWeb     string `json:"rentalUriWeb"`
				FormFactor       string `json:"formFactor"`
				PropulsionType   string `json:"propulsionType"`
				ReturnConstraint string `json:"returnConstraint"`
			} `json:"rental"`
			FareTransferIndex     int `json:"fareTransferIndex"`
			EffectiveFareLegIndex int `json:"effectiveFareLegIndex"`
		} `json:"legs"`
		FareTransfers []struct {
			Rule            string `json:"rule"`
			TransferProduct struct {
				Name          string `json:"name"`
				Amount        int    `json:"amount"`
				Currency      string `json:"currency"`
				RiderCategory struct {
					RiderCategoryName     string `json:"riderCategoryName"`
					IsDefaultFareCategory bool   `json:"isDefaultFareCategory"`
					EligibilityURL        string `json:"eligibilityUrl"`
				} `json:"riderCategory"`
				Media struct {
					FareMediaName string `json:"fareMediaName"`
					FareMediaType string `json:"fareMediaType"`
				} `json:"media"`
			} `json:"transferProduct"`
			EffectiveFareLegProducts [][]struct {
				Name          string `json:"name"`
				Amount        int    `json:"amount"`
				Currency      string `json:"currency"`
				RiderCategory struct {
					RiderCategoryName     string `json:"riderCategoryName"`
					IsDefaultFareCategory bool   `json:"isDefaultFareCategory"`
					EligibilityURL        string `json:"eligibilityUrl"`
				} `json:"riderCategory"`
				Media struct {
					FareMediaName string `json:"fareMediaName"`
					FareMediaType string `json:"fareMediaType"`
				} `json:"media"`
			} `json:"effectiveFareLegProducts"`
		} `json:"fareTransfers"`
	} `json:"direct"`
	Itineraries []struct {
		Duration  int       `json:"duration"`
		StartTime time.Time `json:"startTime"`
		EndTime   time.Time `json:"endTime"`
		Transfers int       `json:"transfers"`
		Legs      []struct {
			Mode TransitMode `json:"mode"`
			From struct {
				Name               string    `json:"name"`
				StopID             string    `json:"stopId"`
				Lat                float64   `json:"lat"`
				Lon                float64   `json:"lon"`
				Level              float64   `json:"level"`
				Arrival            time.Time `json:"arrival"`
				Departure          time.Time `json:"departure"`
				ScheduledArrival   time.Time `json:"scheduledArrival"`
				ScheduledDeparture time.Time `json:"scheduledDeparture"`
				ScheduledTrack     string    `json:"scheduledTrack"`
				Track              string    `json:"track"`
				VertexType         string    `json:"vertexType"`
			} `json:"from"`
			To struct {
				Name               string    `json:"name"`
				StopID             string    `json:"stopId"`
				Lat                float64   `json:"lat"`
				Lon                float64   `json:"lon"`
				Level              float64   `json:"level"`
				Arrival            time.Time `json:"arrival"`
				Departure          time.Time `json:"departure"`
				ScheduledArrival   time.Time `json:"scheduledArrival"`
				ScheduledDeparture time.Time `json:"scheduledDeparture"`
				ScheduledTrack     string    `json:"scheduledTrack"`
				Track              string    `json:"track"`
				VertexType         string    `json:"vertexType"`
			} `json:"to"`
			Duration                 int       `json:"duration"`
			StartTime                time.Time `json:"startTime"`
			EndTime                  time.Time `json:"endTime"`
			ScheduledStartTime       time.Time `json:"scheduledStartTime"`
			ScheduledEndTime         time.Time `json:"scheduledEndTime"`
			RealTime                 bool      `json:"realTime"`
			Distance                 float64   `json:"distance"`
			InterlineWithPreviousLeg bool      `json:"interlineWithPreviousLeg"`
			Headsign                 string    `json:"headsign"`
			RouteColor               string    `json:"routeColor"`
			RouteTextColor           string    `json:"routeTextColor"`
			RouteType                string    `json:"routeType"`
			AgencyName               string    `json:"agencyName"`
			AgencyURL                string    `json:"agencyUrl"`
			AgencyID                 string    `json:"agencyId"`
			TripID                   string    `json:"tripId"`
			RouteShortName           string    `json:"routeShortName"`
			Source                   string    `json:"source"`
			IntermediateStops        []struct {
				Name               string    `json:"name"`
				StopID             string    `json:"stopId"`
				Lat                float64   `json:"lat"`
				Lon                float64   `json:"lon"`
				Level              float64   `json:"level"`
				Arrival            time.Time `json:"arrival"`
				Departure          time.Time `json:"departure"`
				ScheduledArrival   time.Time `json:"scheduledArrival"`
				ScheduledDeparture time.Time `json:"scheduledDeparture"`
				ScheduledTrack     string    `json:"scheduledTrack"`
				Track              string    `json:"track"`
				VertexType         string    `json:"vertexType"`
			} `json:"intermediateStops"`
			LegGeometry struct {
				Points string `json:"points"`
				Length int    `json:"length"`
			} `json:"legGeometry"`
			Steps []struct {
				RelativeDirection string  `json:"relativeDirection"`
				Distance          float64 `json:"distance"`
				FromLevel         float64 `json:"fromLevel"`
				ToLevel           float64 `json:"toLevel"`
				OsmWay            int     `json:"osmWay"`
				Polyline          struct {
					Points string `json:"points"`
					Length int    `json:"length"`
				} `json:"polyline"`
				StreetName string `json:"streetName"`
				Exit       string `json:"exit"`
				StayOn     bool   `json:"stayOn"`
				Area       bool   `json:"area"`
			} `json:"steps"`
			Rental struct {
				SystemID         string `json:"systemId"`
				SystemName       string `json:"systemName"`
				URL              string `json:"url"`
				StationName      string `json:"stationName"`
				FromStationName  string `json:"fromStationName"`
				ToStationName    string `json:"toStationName"`
				RentalURIAndroid string `json:"rentalUriAndroid"`
				RentalURIIOS     string `json:"rentalUriIOS"`
				RentalURIWeb     string `json:"rentalUriWeb"`
				FormFactor       string `json:"formFactor"`
				PropulsionType   string `json:"propulsionType"`
				ReturnConstraint string `json:"returnConstraint"`
			} `json:"rental"`
			FareTransferIndex     int `json:"fareTransferIndex"`
			EffectiveFareLegIndex int `json:"effectiveFareLegIndex"`
		} `json:"legs"`
		FareTransfers []struct {
			Rule            string `json:"rule"`
			TransferProduct struct {
				Name          string `json:"name"`
				Amount        int    `json:"amount"`
				Currency      string `json:"currency"`
				RiderCategory struct {
					RiderCategoryName     string `json:"riderCategoryName"`
					IsDefaultFareCategory bool   `json:"isDefaultFareCategory"`
					EligibilityURL        string `json:"eligibilityUrl"`
				} `json:"riderCategory"`
				Media struct {
					FareMediaName string `json:"fareMediaName"`
					FareMediaType string `json:"fareMediaType"`
				} `json:"media"`
			} `json:"transferProduct"`
			EffectiveFareLegProducts [][]struct {
				Name          string `json:"name"`
				Amount        int    `json:"amount"`
				Currency      string `json:"currency"`
				RiderCategory struct {
					RiderCategoryName     string `json:"riderCategoryName"`
					IsDefaultFareCategory bool   `json:"isDefaultFareCategory"`
					EligibilityURL        string `json:"eligibilityUrl"`
				} `json:"riderCategory"`
				Media struct {
					FareMediaName string `json:"fareMediaName"`
					FareMediaType string `json:"fareMediaType"`
				} `json:"media"`
			} `json:"effectiveFareLegProducts"`
		} `json:"fareTransfers"`
	} `json:"itineraries"`
	PreviousPageCursor string `json:"previousPageCursor"`
	NextPageCursor     string `json:"nextPageCursor"`
}

func (c *Client) Plan(time time.Time, fromPlace, toPlace string, transitModes []TransitMode) (Plan, error) {
	var r Plan

	modes := make([]string, len(transitModes))
	for i, mode := range transitModes {
		modes[i] = string(mode)
	}

	url := c.ResolveURL("/plan")
	query := url.Query()
	query.Set("fromPlace", fromPlace)
	query.Set("toPlace", toPlace)
	query.Set("mode", strings.Join(modes, ","))
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
