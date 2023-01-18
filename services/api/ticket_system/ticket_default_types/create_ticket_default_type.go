package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

type TicketDefaultTypeService struct {
	TicketDefaultType models.TicketDefaultType
}

func CreateTicketDefaultType(ticket_default_type models.TicketDefaultType) models.TicketDefaultType {
	db := config.GetDB()
	ticket_default_type.Status = "active"
	// result := map[string]interface{}{}
	db.Create(&ticket_default_type)
	return ticket_default_type
}
