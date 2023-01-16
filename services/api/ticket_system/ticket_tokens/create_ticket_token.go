package ticket_system

import (
	"encoding/hex"
	"strconv"
	"time"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	ticketuser "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_users"
)

type TicketTokenService struct {
	TicketToken models.TicketToken
}

func CreateTicketToken(body models.Filter) models.TicketToken {
	db := config.GetDB()

	var ticket_token models.TicketToken

	ticket_user := ticketuser.CreateTicketUser(body.TicketUser)

	result := strconv.FormatUint(uint64(ticket_user.ID), 10)

	result = result + time.Now().String() + "dunningtoken"

	message := []byte(result)

	encoded_string := hex.EncodeToString(message)

	ticket_token.TicketToken = encoded_string

	ticket_token.ExpiryDate = time.Now().Add(time.Hour)

	ticket_token.TicketUserID = ticket_user.ID

	db.Create(&ticket_token)

	return ticket_token
}
