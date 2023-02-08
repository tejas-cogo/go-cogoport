package apihelper

import (
	"encoding/json"
	"fmt"

	"github.com/lib/pq"
	"github.com/tejas-cogo/go-cogoport/models"
)

func GetUserData(IDs pq.StringArray) []models.UserData {
	var rubyclient models.RubyClientInput

	type Filters struct {
		ID pq.StringArray `json:"id"`
	}

	type Body struct {
		Filters Filters `json:"filters"`
	}

	type Response struct {
		List       []models.UserData
		Page       uint `json:"page"`
		Total      uint `json:"total"`
		TotalCount uint `json:"total_count"`
		PageLimit  uint `json:"page_limit"`
	}

	var user Response
	var body Body

	rubyclient.Endpoint = "user/list_users"
	body.Filters.ID = IDs
	obj, _ := GetRubyClient(body, rubyclient)

	bodyString := string(obj)

	err := json.Unmarshal([]byte(bodyString), &user)
	if err != nil {
		fmt.Println(err, "Error occured")
	}

	var users []models.UserData
	for _, user_details := range user.List {
		users = append(users, user_details)
	}
	return users
}
