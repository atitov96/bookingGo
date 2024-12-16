package memory

import (
	"context"
	"sync"
	"time"

	"applicationDesignTest/internal/domain/models"
)

type AvailabilityRepository struct {
	mu           sync.RWMutex
	availability map[string]*models.RoomAvailability
}

func NewAvailabilityRepository(initial []*models.RoomAvailability) *AvailabilityRepository {
	repo := &AvailabilityRepository{
		availability: make(map[string]*models.RoomAvailability),
	}

	for _, a := range initial {
		key := getKey(a.HotelID, a.RoomID, a.Date)
		repo.availability[key] = a
	}

	return repo
}

func (r *AvailabilityRepository) UpdateQuota(ctx context.Context, hotelID, roomID string, date time.Time, delta int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	key := getKey(hotelID, roomID, date)
	if avail, exists := r.availability[key]; exists {
		avail.Quota += delta
		return nil
	}
	return nil
}

func (r *AvailabilityRepository) GetAvailability(ctx context.Context, hotelID, roomID string, date time.Time) (*models.RoomAvailability, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	key := getKey(hotelID, roomID, date)
	if avail, exists := r.availability[key]; exists {
		return avail, nil
	}
	return nil, nil
}

func getKey(hotelID, roomID string, date time.Time) string {
	return hotelID + "_" + roomID + "_" + date.Format("2006-01-02")
}
