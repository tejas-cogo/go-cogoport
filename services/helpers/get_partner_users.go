package apihelper

import (
	"encoding/json"
	"fmt"

	"github.com/lib/pq"
	"github.com/tejas-cogo/go-cogoport/models"
)

func GetPartnerUserData(UserIDs pq.StringArray) []models.User {
	var rubyclient models.RubyClientInput

	type Filters struct {
		UserID pq.StringArray `json:"user_id"`
	}

	type Body struct {
		Filters Filters `json:"filters"`
	}

	type Response struct {
		List       []models.User
		Page       uint `json:"page"`
		Total      uint `json:"total"`
		TotalCount uint `json:"total_count"`
		PageLimit  uint `json:"page_limit"`
	}
	var body Body

	rubyclient.Endpoint = "partner/list_partner_users"
	body.Filters.UserID = UserIDs

	var partner_users Response
	obj, _ := RubyClient(body, rubyclient)

	bodyString := string(obj)

	err := json.Unmarshal([]byte(bodyString), &partner_users)
	if err != nil {
		fmt.Println(err, "Error occured")
	}

	var users []models.User
	for _, user_details := range partner_users.List {
		users = append(users, user_details)
	}
	return users
}
