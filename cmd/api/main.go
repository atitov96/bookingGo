package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"applicationDesignTest/internal/application/booking"
	"applicationDesignTest/internal/domain/models"
	"applicationDesignTest/internal/infrastructure/memory"
	"applicationDesignTest/internal/interfaces/http/handlers"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)

	now := time.Now().UTC()
	tomorrow := now.AddDate(0, 0, 1)

	initialAvailability := []*models.RoomAvailability{
		{"reddison", "lux", tomorrow, 1},
		{"reddison", "lux", tomorrow.AddDate(0, 0, 1), 1},
		{"reddison", "lux", tomorrow.AddDate(0, 0, 2), 1},
		{"reddison", "lux", tomorrow.AddDate(0, 0, 3), 1},
		{"reddison", "lux", tomorrow.AddDate(0, 0, 4), 0},
	}

	orderRepo := memory.NewOrderRepository()
	availabilityRepo := memory.NewAvailabilityRepository(initialAvailability)

	bookingService := booking.NewBookingService(orderRepo, availabilityRepo)

	bookingHandler := handlers.NewBookingHandler(bookingService, logger)

	mux := http.NewServeMux()
	mux.HandleFunc("/orders", bookingHandler.CreateOrder)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	logger.Printf("Server starting on :8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Server failed to start: %v", err)
	}
}
