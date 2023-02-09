package apihelper

import (
	"encoding/json"
	"fmt"

	"github.com/lib/pq"
	"github.com/tejas-cogo/go-cogoport/models"
)

func GetAuthRoleData(RoleIDs pq.StringArray) []models.AuthRoleData {
	var rubyclient models.RubyClientInput

	type Filters struct {
		ID pq.StringArray `json:"id"`
	}

	type Body struct {
		Filters Filters `json:"filters"`
	}

	type Response struct {
		List       []models.AuthRoleData
		Page       uint `json:"page"`
		Total      uint `json:"total"`
		TotalCount uint `json:"total_count"`
		PageLimit  uint `json:"page_limit"`
	}

	var auth_user Response
	var body Body

	rubyclient.Endpoint = "auth/list_auth_roles"
	body.Filters.ID = RoleIDs
	obj, _ := RubyClient(body, rubyclient)

	bodyString := string(obj)

	err := json.Unmarshal([]byte(bodyString), &auth_user)
	if err != nil {
		fmt.Println(err, "Error occured")
	}

	var auth_users []models.AuthRoleData
	for _, user_details := range auth_user.List {
		auth_users = append(auth_users, user_details)
	}
	return auth_users
}
