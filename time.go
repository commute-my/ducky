package ducky

import "time"

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
