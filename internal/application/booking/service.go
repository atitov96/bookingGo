package booking

import (
	"context"
	"errors"
	"time"

	"applicationDesignTest/internal/domain/models"
	"applicationDesignTest/internal/domain/repository"
)

type BookingService struct {
	orderRepo        repository.OrderRepository
	availabilityRepo repository.AvailabilityRepository
}

func NewBookingService(orderRepo repository.OrderRepository, availabilityRepo repository.AvailabilityRepository) *BookingService {
	return &BookingService{
		orderRepo:        orderRepo,
		availabilityRepo: availabilityRepo,
	}
}

func (s *BookingService) CreateOrder(ctx context.Context, order *models.Order) error {
	if err := order.Validate(); err != nil {
		return err
	}

	days := getDaysBetween(order.From, order.To)

	for _, day := range days {
		avail, err := s.availabilityRepo.GetAvailability(ctx, order.HotelID, order.RoomID, day)
		if err != nil {
			return err
		}
		if avail == nil || avail.Quota < 1 {
			return errors.New("нет доступных номеров на выбранные даты")
		}
	}

	for _, day := range days {
		err := s.availabilityRepo.UpdateQuota(ctx, order.HotelID, order.RoomID, day, -1)
		if err != nil {
			return err
		}
	}

	return s.orderRepo.Save(ctx, order)
}

func getDaysBetween(from, to time.Time) []time.Time {
	var days []time.Time
	for d := from; !d.After(to); d = d.AddDate(0, 0, 1) {
		days = append(days, time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.UTC))
	}
	return days
}
