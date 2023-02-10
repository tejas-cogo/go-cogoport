package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type TicketActivity struct {
	gorm.Model
	TicketID    uint      `gorm:"not null"`
	UserID      uuid.UUID `gorm:"type:uuid"`
	UserType    string    `gorm:"not null"`
	Description string
	Type        string `gorm:"not null"`
	Data        Data
	IsRead      bool
	Status      string
	Ticket      Ticket `gorm:"foreignKey:TicketID"`
}
type TicketActivityData struct {
	TicketID    uint      `gorm:"not null"`
	UserID      uuid.UUID `gorm:"type:uuid"`
	UserType    string    `gorm:"not null"`
	Description string
	Type        string `gorm:"not null"`
	Data        Data
	IsRead      bool
	Status      string
	TicketUser  TicketUser
	Ticket      Ticket `gorm:"foreignKey:TicketID"`
	CreatedAt   time.Time
}

type Activity struct {
	PerformedByID uuid.UUID `gorm:"type:uuid"`
	TicketID      []uint
	UserID        uuid.UUID `gorm:"type:uuid"`
	UserType      string
	Description   string
	Data          Data
	Type          string
	Status        string
}

type Data struct {
	Url    pq.StringArray `gorm:"type:text[]"`
	User   []UserData
	UserID uint `json:"user_id"`
}
