package api

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	helpers "github.com/tejas-cogo/go-cogoport/services/helpers"
)

func SendCommunications(ID uint) {
	type Response struct {
		ID uuid.UUID
	}

	var rubyclient models.RubyClientInput
	var body models.Communication
	var communication_id Response

	db := config.GetDB()

	var ticket models.Ticket
	db.Where("id = ?", ID).First(&ticket)

	var ticket_user models.TicketUser
	db.Where("id = ?", ticket.TicketUserID).First(&ticket_user)
	
	rubyclient.Endpoint = "communication/create_communication"

	for _, preference := range ticket.NotificationPreferences {
		if preference == "email"{
			body.Recipient = ticket_user.Email
			body.Type = "email"
			body.Service = "user"
			body.ServiceID = "f06b29c0-b443-4f71-bf64-b61106dcaaf8"
			body.TemplateName = "Ticket System"
		} else if preference == "whatsapp"{
			body.Recipient = ticket_user.MobileNumber
			body.Type = "whatsapp"
			body.Service = "user"
			body.ServiceID = "f06b29c0-b443-4f71-bf64-b61106dcaaf8"
			body.TemplateName = "Ticket System"
		}
	
		obj, _ := helpers.PostRubyClient(body, rubyclient)
	
		bodyString := string(obj)
		err := json.Unmarshal([]byte(bodyString), &communication_id)
		if err != nil {
			log.Println(err)
		}
	
		if communication_id.ID != uuid.Nil {
			log.Printf("Successfully Created!")
		}
	}
}
