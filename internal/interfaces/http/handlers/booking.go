package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"applicationDesignTest/internal/application/booking"
	"applicationDesignTest/internal/domain/models"
)

type BookingHandler struct {
	bookingService *booking.BookingService
	logger         *log.Logger
}

func NewBookingHandler(bookingService *booking.BookingService, logger *log.Logger) *BookingHandler {
	return &BookingHandler{
		bookingService: bookingService,
		logger:         logger,
	}
}

func (h *BookingHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		h.logger.Printf("[ERROR] Failed to decode request: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.bookingService.CreateOrder(r.Context(), &order); err != nil {
		h.logger.Printf("[ERROR] Failed to create order: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}
