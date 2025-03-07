package ducky

import "fmt"

// FormatDistance converts meters to km with one decimal place as a string.
func FormatDistance(meters float64) string {
	if meters < 1000 {
		return fmt.Sprintf("%.0f m", meters)
	}
	return fmt.Sprintf("%.1f km", meters/1000)
}

// FormatDuration converts seconds to a human-readable duration.
func FormatDuration(seconds float64) string {
	minutes := int(seconds / 60)
	if minutes < 60 {
		return fmt.Sprintf("%d min", minutes)
	}
	hours := minutes / 60
	mins := minutes % 60
	if mins == 0 {
		return fmt.Sprintf("%d hr", hours)
	}
	return fmt.Sprintf("%d hr %d min", hours, mins)
}

// GetTransportationType returns a human-readable description of transport types.
func GetTransportationType(routeType int) string {
	switch routeType {
	case 0:
		return "tram"
	case 1:
		return "subway"
	case 2:
		return "rail"
	case 3:
		return "bus"
	case 4:
		return "ferry"
	case 5:
		return "cable_tram"
	case 6:
		return "aerial_lift"
	case 7:
		return "funicular"
	case 8:
		return "trolleybus"
	case 9:
		return "monorail"
	case 11:
		return "walking"
	case 12:
		return "cycling"
	default:
		return "transit"
	}
}
