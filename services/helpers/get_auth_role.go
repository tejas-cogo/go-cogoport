package apihelper

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/tejas-cogo/go-cogoport/config"
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
	obj, _ := GetRubyClient(body, rubyclient)

	bodyString := string(obj)

	err := json.Unmarshal([]byte(bodyString), &auth_user)
	if err != nil {
		log.Println(err)
	}

	var auth_users []models.AuthRoleData
	for _, user_details := range auth_user.List {
		auth_users = append(auth_users, user_details)
	}
	return auth_users
}

func GetUnifiedAuthRoleData(RoleIDs pq.StringArray) []models.AuthRoleData {

	db2 := config.GetCDB()

	var auth_role_data []models.AuthRoleData

	var auth_role_ids []uuid.UUID

	for _, u := range RoleIDs {
		data, err := uuid.Parse(u)
		if err != nil {
			log.Println(err)
		}
		auth_role_ids = append(auth_role_ids, data)
	}

	db2.Model(&models.AuthRole{}).Where("id in (?) and status = ?", auth_role_ids, "active").Scan(&auth_role_data)

	return auth_role_data
}
