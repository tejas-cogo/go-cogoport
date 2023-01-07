package ticket_system

import (
	"fmt"
	"github.com/ChandelShikha/go-cogoport/config"
	"github.com/ChandelShikha/go-cogoport/models"
)

type TicketDefaultTypeService struct {
	TicketDefaultType models.TicketDefaultType
}

func CreateTicketDefaultType(ticket_default_type models.TicketDefaultType) models.TicketDefaultType {
	db := config.GetDB()
	fmt.Println(ticket_default_type)
	// result := map[string]interface{}{}
	db.Create(&ticket_default_type)
	return ticket_default_type
}