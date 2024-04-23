package model

type Ticket struct {
	TicketId        uint   `gorm:"primaryKey"`
	TicketTrain     string `gorm:"not null"`
	TicketFromCity  string `gorm:"not null"`
	TicketToCity    string `gorm:"not null"`
	TicketClass     string `gorm:"not null"`
	TicketPrice     string `gorm:"not null"`
	TicketDate      string `gorm:"not null"`
	TicketDeparture string `gorm:"not null"`
	TicketArrived   string `gorm:"not null"`
	CreatedAt       string `gorm:"not null"`
	UpdatedAt       string
}
