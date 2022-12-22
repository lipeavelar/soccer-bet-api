package helpers

import (
	"errors"
	"strings"
	"time"
)

var mapTimezones = map[string]string{
	"brazil": "America/Sao_Paulo",
}

func ConvertToTimezone(utcTime time.Time, timezone string) (time.Time, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, err
	}
	return utcTime.In(location), nil
}

func GetTimezoneString(country string) (string, error) {
	timezone, ok := mapTimezones[strings.ToLower(country)]
	if !ok {
		return "", errors.New("timezone not found")
	}
	return timezone, nil
}

// CompareDates compare two dates
// return -1 if date1 < date2
// return 0 if date1 = date2
// return 1 if date1 > date2
func CompareDates(date1 time.Time, date2 time.Time) int {
	d1 := time.Date(date1.Year(), date1.Month(), date1.Day(), 0, 0, 0, 0, time.UTC)
	d2 := time.Date(date2.Year(), date2.Month(), date2.Day(), 0, 0, 0, 0, time.UTC)
	if d1.Before(d2) {
		return -1
	}
	if d1.After(d2) {
		return 1
	}
	return 0
}
