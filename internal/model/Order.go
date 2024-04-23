package model

import "github.com/google/uuid"

type Order struct {
	OrderId              uint      `gorm:"primaryKey"`
	OrderNumber          uuid.UUID `gorm:"not null"`
	OrderPassengerName   string    `gorm:"not null"`
	OrderPassengerAmount string    `gorm:"not null"`
	OrderHpNumber        string    `gorm:"not null"`
	OrderResidenceNumber string    `gorm:"not null"` // NIK
	OrderTicketStatus    string    `gorm:"not null"` // pending , active (paid) , inactive (payment process)
	OrderUserId          uint      `gorm:"not null"` //check authentication
	User                 User      `gorm:"foreignKey:OrderUserId"`
	OrderSeatId          uint      `gorm:"not null"`
	Seat                 Seat      `gorm:"foreignKey:OrderSeatId"`
	OrderTicketId        uint      `gorm:"not null"`
	Ticket               Ticket    `gorm:"foreignKey:OrderTicketId"`
	CreatedAt            string    `gorm:"not null"`
	UpdatedAt            string
}
