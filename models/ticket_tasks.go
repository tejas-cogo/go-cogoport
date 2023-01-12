package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketTask struct {
	gorm.Model
	TicketID        uint
	Ticket Ticket
	Title           string
	CreatedByUserId uuid.UUID
	Status          string
}
