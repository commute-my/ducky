package motis

type TransitMode string

const (
	TransitModeBus    TransitMode = "BUS"
	TransitModeSubway TransitMode = "SUBWAY"
	TransitModeWalk   TransitMode = "WALK"
	TransitModeOther  TransitMode = "OTHER"
)
