package ticket_system

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func SendTicketActivity(ticket_activity models.TicketActivity) {
	db := config.GetDB()

	if ticket_activity.Type == "respond" {
		var ticket_user models.TicketUser

		db.Where("id = ?", ticket_activity.TicketUserID).First(&ticket_user)

		hc := http.Client{}

		var body models.Post

		body.Recipient = ticket_user.Email
		body.Type = "email"
		body.Service = "user"
		body.ServiceID = ticket_user.SystemUserID
		body.TemplateName = "Ticket System"

		reqBody, err := json.Marshal(body)

		req, err := http.NewRequest("POST", "https://api-apollo3.dev.cogoport.io/communication/create_communication", bytes.NewBuffer(reqBody))
		if err != nil {
			log.Printf("Request Failed: %s", err)
			return
		}

		req.Header.Add("Authorization", "Bearer: 787b8f21-ca0a-4e79-af6e-81e3ca847909")
		req.Header.Add("AuthorizationScope", "partner")

		req.Header.Add("Content-type", "application/json; charset=utf-8")

		resp, err := hc.Do(req)

		if err == nil {

			defer resp.Body.Close()
		}

		resp_body, err := ioutil.ReadAll(resp.Body)

		// Log the request body
		bodyString := string(resp_body)
		log.Print(bodyString)
		// Unmarshal result
		post := models.Post{}
		err = json.Unmarshal(resp_body, &post)
		if err != nil {
			log.Printf("Reading body failed: %s", err)
			return
		}

		log.Printf("Successfully Created")

	}
}
