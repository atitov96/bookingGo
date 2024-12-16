package memory

import (
	"context"
	"sync"
	"time"

	"applicationDesignTest/internal/domain/models"
)

type OrderRepository struct {
	mu     sync.RWMutex
	orders []*models.Order
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		orders: make([]*models.Order, 0),
	}
}

func (r *OrderRepository) Save(ctx context.Context, order *models.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.orders = append(r.orders, order)
	return nil
}

func (r *OrderRepository) FindByDateRange(ctx context.Context, hotelID, roomID string, from, to time.Time) ([]*models.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*models.Order
	for _, order := range r.orders {
		if order.HotelID == hotelID && order.RoomID == roomID &&
			!order.To.Before(from) && !order.From.After(to) {
			result = append(result, order)
		}
	}
	return result, nil
}
