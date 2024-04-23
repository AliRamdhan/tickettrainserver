package model

type Seat struct {
	SeatId     uint   `gorm:"primaryKey"`
	SeatNumber string `gorm:"not null"`
	SeatTicket uint   `gorm:"not null"`
	Ticket     Ticket `gorm:"foreignKey:SeatTicket"`
	CreatedAt  string `gorm:"not null"`
	UpdatedAt  string
}
