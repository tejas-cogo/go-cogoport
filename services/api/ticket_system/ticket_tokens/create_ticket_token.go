package ticket_system

import (
	"encoding/hex"
	"errors"
	"strconv"
	"time"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	ticketuser "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_users"
)

type TicketTokenService struct {
	TicketToken models.TicketToken
}

func CreateTicketToken(body models.TicketUser) (models.TicketToken, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_token models.TicketToken

	body.RoleID = 1

	ticket_user, _ := ticketuser.CreateTicketUser(body)

	result := strconv.FormatUint(uint64(ticket_user.ID), 10)

	result = result + time.Now().String() + "token"

	message := []byte(result)

	encoded_string := hex.EncodeToString(message)

	ticket_token.TicketToken = encoded_string

	ticket_token.ExpiryDate = time.Now().Add(time.Hour)

	ticket_token.TicketUserID = ticket_user.ID
	ticket_token.Status = "active"

	stmt := validate(ticket_token)
		if stmt != "validated" {
			return ticket_token, errors.New(stmt)
		}

	if err := tx.Create(&ticket_token).Error; err != nil {
		tx.Rollback()
		return ticket_token, errors.New("Error Occured!")
	}

	tx.Commit()
	return ticket_token, err
}

func validate(ticket_token models.TicketToken) string {
	if ticket_token.TicketUserID <= 0 {
		return ("UserID is Invalid!")
	}

	return ("validated")
}
