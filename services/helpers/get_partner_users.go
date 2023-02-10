package apihelper

import (
	"encoding/json"
	"log"

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
	obj, _ := GetRubyClient(body, rubyclient)

	bodyString := string(obj)

	err := json.Unmarshal([]byte(bodyString), &partner_users)
	if err != nil {
		log.Println(err)
	}

	var users []models.User
	for _, user_details := range partner_users.List {
		users = append(users, user_details)
	}
	return users
}

// func GetUnifiedPartnerUserData(UserIDs pq.StringArray) []models.User {

// 	db2 := config.GetCDB()

// 	var users []models.User

// 	var user_ids []uuid.UUID

// 	for _, u := range UserIDs {
// 		data, err := uuid.Parse(u)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		user_ids = append(user_ids, data)
// 	}

// 	db2.Model(&models.AuthRole{}).Where("id in (?) and status = ?", user_ids, "active").Scan(&users)

// 	return users
// }
