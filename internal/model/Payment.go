package model

import "github.com/google/uuid"

type Payment struct {
	PaymentId      uint      `gorm:"primaryKey"`
	PaymentStatus  string    `gorm:"not null"` // Not Complete , Complete
	PaymentTotal   uint      `gorm:"not null"`
	PaymentMethod  string    `gorm:"not null"`
	PaymentCode    uuid.UUID `gorm:"not null"`
	PaymentOrderId uint      `gorm:"not null"`
	Order          Order     `gorm:"foreignKey:PaymentOrderId"`
	CreatedAt      string    `gorm:"not null"`
	UpdatedAt      string
}
