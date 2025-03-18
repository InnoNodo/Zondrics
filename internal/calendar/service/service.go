package service

import (
	"Zondrics/internal/database/models"
	"errors"
	"gorm.io/gorm"
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

func GetAvailableSlots(db *gorm.DB, organizationID uint, organizationUserID uint, slotDuration time.Duration, workStart time.Time, workEnd time.Time) ([]models.SlotStatus, error) {
	var slots []models.SlotStatus

	now := time.Now()

	var allPossibleSlots []models.Timeslot
	currentTime := workStart
	for currentTime.Add(slotDuration).Before(workEnd) || currentTime.Add(slotDuration).Equal(workEnd) {
		slotEnd := currentTime.Add(slotDuration)
		allPossibleSlots = append(allPossibleSlots, models.Timeslot{
			StartTime: currentTime,
			EndTime:   slotEnd,
		})
		currentTime = currentTime.Add(slotDuration)
	}

	var bookedSlots []models.Timeslot
	if err := db.Where("organization_id = ? AND organization_user_id = ? AND ((start_time >= ? AND start_time < ?) OR (end_time > ? AND end_time <= ?))",
		organizationID, organizationUserID, workStart, workEnd, workStart, workEnd).Find(&bookedSlots).Error; err != nil {
		return nil, err
	}

	for _, possibleSlot := range allPossibleSlots {
		isBooked := false

		for _, bookedSlot := range bookedSlots {
			if possibleSlot.StartTime.Equal(bookedSlot.StartTime) && possibleSlot.EndTime.Equal(bookedSlot.EndTime) {
				isBooked = true
				break
			}
		}

		if possibleSlot.EndTime.Before(now) {
			isBooked = true
		}

		slots = append(slots, models.SlotStatus{
			StartTime: possibleSlot.StartTime,
			EndTime:   possibleSlot.EndTime,
			IsBooked:  isBooked,
		})
	}

	return slots, nil
}
