package service

import (
	"errors"
	"time"
)

func TimeFormatter(input string) (time.Time, error) {
	layout := "2006-01-02T15:04:05"

	parsedTime, err := time.Parse(layout, input)
	if err != nil {
		return time.Time{}, errors.New("wrong time format")
	}

	return parsedTime, nil
}

func CheckTime(times time.Time) (time.Time, error) {
	if times.Before(time.Now()) || times.Equal(time.Now()) {
		return time.Time{}, errors.New("time must be in the future")
	}

	return times, nil
}
