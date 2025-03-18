package service

import (
	"errors"
	"time"
)

func FormatTime(times string) (time.Time, error) {
	layout := "04:05"

	formatedTime, err := time.Parse(layout, times)
	if err != nil {
		return time.Time{}, errors.New("invalid time format, expected HH:mm")
	}

	return formatedTime, nil
}

func FormatDate(date string) (time.Time, error) {
	layout := "2006-01-02"

	formatedDay, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}, errors.New("invalid date format, expected YYYY-MM-DD")
	}

	return formatedDay, nil
}
