package service

import (
	"time"

	"github.com/AliRamdhan/trainticket/config"
	"github.com/AliRamdhan/trainticket/internal/model"
)

type TicketService struct{}

func NewTicketService() *TicketService {
	return &TicketService{}
}

func (ps *TicketService) CreateTicket(ticket *model.Ticket) error {
	ticket.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return config.DB.Create(ticket).Error
}

func (ps *TicketService) GetAllTicket() ([]model.Ticket, error) {
	var tickets []model.Ticket
	if err := config.DB.Find(&tickets).Error; err != nil {
		return nil, err
	}
	return tickets, nil
}

func (ps *TicketService) GetTicketById(ticketId uint) (*model.Ticket, error) {
	var tickets model.Ticket
	if err := config.DB.First(&tickets, "ticket_id = ?", ticketId).Error; err != nil {
		return nil, err // Ticket not found
	}
	return &tickets, nil
}

func (ps *TicketService) UpdateTicket(ticketID uint, updatedTicket *model.Ticket) error {
	// Find the Ticket with the given ID
	var existingTicket model.Ticket
	if err := config.DB.First(&existingTicket, "ticket_id = ?", ticketID).Error; err != nil {
		return err // Ticket not found
	}

	// Update fields of existing Ticket with the new values
	existingTicket.TicketTrain = updatedTicket.TicketTrain
	existingTicket.TicketFromCity = updatedTicket.TicketFromCity
	existingTicket.TicketToCity = updatedTicket.TicketToCity
	existingTicket.TicketClass = updatedTicket.TicketClass
	existingTicket.TicketPrice = updatedTicket.TicketPrice
	existingTicket.TicketDate = updatedTicket.TicketDate
	existingTicket.TicketDeparture = updatedTicket.TicketDeparture
	existingTicket.TicketArrived = updatedTicket.TicketArrived
	existingTicket.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	// Save the updated Ticket
	return config.DB.Save(&existingTicket).Error
}

func (ps *TicketService) DeleteTicket(ticketId uint) error {
	// Find the Ticket with the given ID
	var ticket model.Ticket
	if err := config.DB.First(&ticket, "ticket_id = ?", ticketId).Error; err != nil {
		return err // ticket not found
	}
	// Delete the ticket
	return config.DB.Delete(&ticket).Error
}
