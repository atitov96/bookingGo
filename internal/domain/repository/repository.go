package repository

import (
	"context"
	"time"

	"applicationDesignTest/internal/domain/models"
)

type OrderRepository interface {
	Save(ctx context.Context, order *models.Order) error
	FindByDateRange(ctx context.Context, hotelID, roomID string, from, to time.Time) ([]*models.Order, error)
}

type AvailabilityRepository interface {
	UpdateQuota(ctx context.Context, hotelID, roomID string, date time.Time, delta int) error
	GetAvailability(ctx context.Context, hotelID, roomID string, date time.Time) (*models.RoomAvailability, error)
}
