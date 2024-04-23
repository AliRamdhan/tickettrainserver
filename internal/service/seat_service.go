package service

import (
	"time"

	"github.com/AliRamdhan/trainticket/config"
	"github.com/AliRamdhan/trainticket/internal/model"
)

type SeatService struct{}

func NewSeatService() *SeatService {
	return &SeatService{}
}

func (ps *SeatService) CreateSeat(ticket *model.Seat, trainTicketId uint) error {
	ticket.SeatTicket = trainTicketId
	ticket.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return config.DB.Create(ticket).Error
}

//	func (ps *SeatService) GetAllSeatsFromTicketId(ticketId uint) (*model.Seat, error) {
//		var trainSeat model.Seat
//		if err := config.DB.Where("train_ticket_id_refer = ?", ticketId).First(&trainSeat).Error; err != nil {
//			return nil, err
//		}
//		return &trainSeat, nil
//	}
func (ps *SeatService) GetAllSeatsFromTicketId(ticketID uint) ([]model.Seat, error) {
	var trainSeats []model.Seat
	if err := config.DB.Where("seat_ticket = ?", ticketID).Find(&trainSeats).Error; err != nil {
		return nil, err
	}
	return trainSeats, nil
}

func (ps *SeatService) UpdateSeats(seatId uint, updateSeat *model.Seat, trainTicketId uint) error {
	// Find the product with the given ID
	var existingSeat model.Seat
	if err := config.DB.First(&existingSeat, "seat_id = ?", seatId).Error; err != nil {
		return err // Product not found
	}

	// Update fields of existing product with the new values
	existingSeat.SeatTicket = trainTicketId
	existingSeat.SeatNumber = updateSeat.SeatNumber
	existingSeat.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	// Save the updated product
	return config.DB.Save(&existingSeat).Error
}

func (ps *SeatService) DeleteSeat(seatId uint) error {
	// Find the product with the given ID
	var seat model.Seat
	if err := config.DB.First(&seat, "seat_id = ?", seatId).Error; err != nil {
		return err // seat not found
	}
	// Delete the seat
	return config.DB.Delete(&seat).Error
}