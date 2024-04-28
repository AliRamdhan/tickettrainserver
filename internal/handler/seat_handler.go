package handler

import (
	"fmt"
	"net/http"

	"github.com/AliRamdhan/trainticket/internal/model"
	"github.com/AliRamdhan/trainticket/internal/service"
	"github.com/gin-gonic/gin"
)

type SeatTrainHandler struct {
	seatTrain *service.SeatService
}

func NewSeatTrainHandler(cs *service.SeatService) *SeatTrainHandler {
	return &SeatTrainHandler{seatTrain: cs}
}

func (sh *SeatTrainHandler) GetAllSeats(c *gin.Context) {
	seats, err := sh.seatTrain.GetAllSeats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "All Seats", "Seats": seats})
}

func (sh *SeatTrainHandler) GetSeatBySeatName(c *gin.Context) {
	seatName := c.Param("seatName")
	seats, err := sh.seatTrain.GetSeatBySeatName(seatName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Details Seats", "Seats": seats})
}

func (sh *SeatTrainHandler) CreateSeat(c *gin.Context) {
	var seat model.Seat
	if err := c.ShouldBindJSON(&seat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Extract serviceCategoryID and customerId from the request body
	ticketTrainId := seat.SeatTicket

	// Check if customerId is missing or invalid
	if ticketTrainId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing customer ID"})
		return
	}

	if err := sh.seatTrain.CreateSeat(&seat, ticketTrainId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Seat created successfully", "order": seat})
}

func (sh *SeatTrainHandler) GetAllSeatsFromTicketId(c *gin.Context) {
	seatIdStr := c.Param("ticketId")

	var seatId uint
	_, err := fmt.Sscanf(seatIdStr, "%d", &seatId)
	// trackID, err := uuid.Parse(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid track number format"})
		return
	}

	seats, err := sh.seatTrain.GetAllSeatsFromTicketId(seatId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Seats not found"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Seats All", "seats": seats})
}

func (sh *SeatTrainHandler) UpdateSeat(c *gin.Context) {
	// Extract seat ID from the URL parameter
	seatIDStr := c.Param("seatId")

	var seatID uint
	_, err := fmt.Sscanf(seatIDStr, "%d", &seatID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid seat ID format"})
		return
	}

	// Bind the updated seat data from JSON
	var updateSeat model.Seat
	if err := c.ShouldBindJSON(&updateSeat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract train ticket ID from the request body
	trainTicketID := updateSeat.SeatTicket

	// Check if train ticket ID is missing or invalid
	if trainTicketID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing train ticket ID"})
		return
	}

	// Call the service to update the seat
	if err := sh.seatTrain.UpdateSeats(seatID, &updateSeat, trainTicketID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Seat updated successfully"})
}

func (sh *SeatTrainHandler) DeleteSeat(c *gin.Context) {
	// Extract seat ID from the URL parameter
	seatIDStr := c.Param("seatId")

	var seatID uint
	_, err := fmt.Sscanf(seatIDStr, "%d", &seatID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid seat ID format"})
		return
	}

	// Call the service to delete the seat
	if err := sh.seatTrain.DeleteSeat(seatID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Seat deleted successfully"})
}
