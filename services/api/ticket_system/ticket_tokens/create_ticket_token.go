package ticket_system

import (
	"crypto/sha256"
	"strconv"
	"time"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

type TicketTokenService struct {
	TicketToken models.TicketToken
}

func CreateTicketToken(ticket_token models.TicketToken) models.TicketToken {
	db := config.GetDB()

	result := strconv.FormatUint(uint64(ticket_token.TicketUserID), 10)

	result += time.Now().String()

	message := []byte(result)
	sha256.Sum256(message)

	ticket_token.TicketToken = "dunning_token" + string(message)

	ticket_token.ExpiryDate = time.Now().Add(time.Hour)

	db.Create(&ticket_token)

	return ticket_token
}
