package models

import (
	"errors"
	"time"
)

type Order struct {
	HotelID   string    `json:"hotel_id"`
	RoomID    string    `json:"room_id"`
	UserEmail string    `json:"email"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
}

func (o *Order) Validate() error {
	if o.HotelID == "" || o.RoomID == "" || o.UserEmail == "" {
		return errors.New("все поля заказа обязательны для заполнения")
	}

	now := time.Now().UTC()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	fromDay := time.Date(o.From.Year(), o.From.Month(), o.From.Day(), 0, 0, 0, 0, time.UTC)
	toDay := time.Date(o.To.Year(), o.To.Month(), o.To.Day(), 0, 0, 0, 0, time.UTC)

	if fromDay.Before(today) {
		return errors.New("дата начала бронирования не может быть в прошлом")
	}

	if toDay.Before(fromDay) {
		return errors.New("дата окончания не может быть раньше даты начала")
	}

	return nil
}

type RoomAvailability struct {
	HotelID string    `json:"hotel_id"`
	RoomID  string    `json:"room_id"`
	Date    time.Time `json:"date"`
	Quota   int       `json:"quota"`
}
